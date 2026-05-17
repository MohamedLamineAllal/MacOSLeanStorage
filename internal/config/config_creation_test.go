package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDefaultConfig_WithMissingDir(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "mls-test-dir")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, "subdir", ".MacosLeanStorage.yaml")
	
	// This should now succeed
	err = CreateDefaultConfig(configPath)
	assert.NoError(t, err)

	_, err = os.Stat(configPath)
	assert.NoError(t, err, "Config file should exist")

	// Verify content
	content, err := os.ReadFile(configPath)
	assert.NoError(t, err)
	assert.Contains(t, string(content), "targets:")
}
