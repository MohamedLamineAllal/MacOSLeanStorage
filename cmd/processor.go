package cmd

import (
	"time"

	"github.com/mohamedlamineallal/MacosLeanStorage/internal/cleaner"
	"github.com/mohamedlamineallal/MacosLeanStorage/internal/config"
	"github.com/mohamedlamineallal/MacosLeanStorage/internal/scanner"
	"github.com/mohamedlamineallal/MacosLeanStorage/internal/scheduler"
	"go.uber.org/zap"
)

type TargetProcessor struct {
	scanner   *scanner.Scanner
	cleaner   *cleaner.Cleaner
	scheduler *scheduler.Scheduler
	logger    *zap.Logger
}

func NewTargetProcessor(logger *zap.Logger, ignorePatterns []string, dryRun bool) *TargetProcessor {
	return &TargetProcessor{
		scanner:   scanner.New(logger, ignorePatterns),
		cleaner:   cleaner.New(logger, dryRun),
		scheduler: scheduler.New(logger),
		logger:    logger,
	}
}

func (tp *TargetProcessor) ProcessTargets(targets []config.TargetConfig) ([]string, []string, []string, int64, error) {
	var allPaths []string
	var allCommands []string
	var commandNames []string
	var totalSize int64

	for _, t := range targets {
		if t.Command != "" {
			if tp.scheduler.ShouldRunCommand(t.Name, t.IntervalDays) {
				allCommands = append(allCommands, t.Command)
				commandNames = append(commandNames, t.Name)
			} else {
				tp.logger.Info("Skipping command target (interval not met)", zap.String("name", t.Name))
			}
			continue
		}

		target := scanner.Target{
			Name:        t.Name,
			Path:        t.Path,
			Threshold:   time.Duration(t.Threshold) * 24 * time.Hour,
			SafetyLevel: t.SafetyLevel,
			Type:        t.Type,
		}

		result, err := tp.scanner.Scan(target, t.IgnorePatterns)
		if err != nil {
			tp.logger.Error("Scan failed for target", zap.String("name", t.Name), zap.Error(err))
			continue
		}

		allPaths = append(allPaths, result.Files...)
		totalSize += result.TotalSize
	}

	return allPaths, allCommands, commandNames, totalSize, nil
}
