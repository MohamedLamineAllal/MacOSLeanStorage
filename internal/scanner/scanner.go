package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bmatcuk/doublestar/v4"
	"go.uber.org/zap"
)

// Target represents a directory or set of paths to be scanned for cleanup.
// It defines the criteria for identifying stale data.
type Target struct {
	Name        string
	Path        string
	Threshold   time.Duration
	SafetyLevel int
	Type        string // "file", "folder", or "both"
}
// Result contains the aggregated findings about a scanned target.
type Result struct {
	TargetName string
	Files      []string
	FileSizes  []int64
	TotalSize  int64
}

// Scanner handles the directory traversal and analysis of filesystem paths.
// It identifies stale files and directories based on age thresholds and configuration.
type Scanner struct {
	logger         *zap.Logger
	ignorePatterns []string
}

// New creates a new Scanner instance with the provided logger and global ignore patterns.
func New(logger *zap.Logger, ignorePatterns []string) *Scanner {
	return &Scanner{logger: logger, ignorePatterns: ignorePatterns}
}

// isIgnored checks if a file or directory name matches any of the configured ignore patterns.
func (s *Scanner) isIgnored(name string) bool {
	for _, pattern := range s.ignorePatterns {
		matched, err := filepath.Match(pattern, name)
		if err == nil && matched {
			return true
		}
	}
	return false
}

// Scan analyzes a target and returns a list of paths that match the cleanup criteria.
// It uses a single-pass traversal for efficiency.
func (s *Scanner) Scan(target Target, targetIgnorePatterns []string) (*Result, error) {
	allIgnorePatterns := append(s.ignorePatterns, targetIgnorePatterns...)
	originalPatterns := s.ignorePatterns
	s.ignorePatterns = allIgnorePatterns
	defer func() { s.ignorePatterns = originalPatterns }()

	expandedPath, err := expandPath(target.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to expand path %s: %w", target.Path, err)
	}

	paths, err := doublestar.FilepathGlob(expandedPath)
	if err != nil {
		return nil, fmt.Errorf("failed to glob path %s: %w", expandedPath, err)
	}

	result := &Result{
		TargetName: target.Name,
		Files:      []string{},
		FileSizes:  []int64{},
	}

	now := time.Now()
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			s.logger.Debug("Failed to stat path", zap.String("path", p), zap.Error(err))
			continue
		}

		if info.IsDir() {
			// Traverse the directory once
			files, sizes, totalSize, isStale := s.traverse(p, target, now)
			if target.Type == "folder" || target.Type == "both" {
				if isStale {
					result.Files = append(result.Files, p)
					result.FileSizes = append(result.FileSizes, totalSize)
					result.TotalSize += totalSize
					continue
				}
			}
			if target.Type == "file" || target.Type == "both" {
				result.Files = append(result.Files, files...)
				result.FileSizes = append(result.FileSizes, sizes...)
				result.TotalSize += totalSize
			}
		} else if target.Type == "file" || target.Type == "both" {
			if now.Sub(info.ModTime()) > target.Threshold {
				result.Files = append(result.Files, p)
				result.FileSizes = append(result.FileSizes, info.Size())
				result.TotalSize += info.Size()
			}
		}
	}
	return result, nil
}

func (s *Scanner) traverse(path string, target Target, now time.Time) (files []string, sizes []int64, totalSize int64, isStale bool) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, nil, 0, false
	}

	isStale = true
	for _, entry := range entries {
		if s.isIgnored(entry.Name()) {
			continue
		}

		fullPath := filepath.Join(path, entry.Name())
		info, err := entry.Info()
		if err != nil {
			continue
		}

		if entry.IsDir() {
			subFiles, subSizes, subTotal, subStale := s.traverse(fullPath, target, now)
			if !subStale {
				isStale = false
			}
			files = append(files, subFiles...)
			sizes = append(sizes, subSizes...)
			totalSize += subTotal
		} else {
			fileAge := now.Sub(info.ModTime())
			if fileAge <= target.Threshold {
				isStale = false
			}
			if fileAge > target.Threshold {
				files = append(files, fullPath)
				sizes = append(sizes, info.Size())
				totalSize += info.Size()
			}
		}
	}
	return files, sizes, totalSize, isStale
}


// expandPath converts a filesystem path (supporting ~ for home directory) into an absolute path.
func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[2:]), nil
	}
	return filepath.Abs(path)
}
