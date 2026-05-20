package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDefaultConfig(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "mls-test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, ".MrLeanStorage.yaml")
	err = CreateDefaultConfig(configPath)
	assert.NoError(t, err)

	_, err = os.Stat(configPath)
	assert.NoError(t, err)

	// Test it doesn't overwrite
	err = os.WriteFile(configPath, []byte("test"), 0644)
	assert.NoError(t, err)
	err = CreateDefaultConfig(configPath)
	assert.NoError(t, err)
	content, err := os.ReadFile(configPath)
	assert.NoError(t, err)
	assert.Equal(t, "test", string(content))
}

func TestGetDefaultConfigPath(t *testing.T) {
	path, err := GetDefaultConfigPath()
	assert.NoError(t, err)
	assert.Contains(t, path, ".MrLeanStorage.yaml")
}
