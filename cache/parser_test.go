package cache_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/odarbelaeze/openapier/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParserCache(t *testing.T) {
	pc := cache.NewParserCache()
	assert.NotNil(t, pc)

	// Create a temporary file to parse
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.go")
	content := []byte("package test\n\ntype T struct{}\n")
	err := os.WriteFile(tmpFile, content, 0644)
	require.NoError(t, err)

	// Test first parse (cache miss)
	file1, err := pc.Parse(tmpFile)
	require.NoError(t, err)
	assert.NotNil(t, file1)
	assert.Equal(t, "test", file1.Name.Name)

	// Test second parse (cache hit)
	file2, err := pc.Parse(tmpFile)
	require.NoError(t, err)
	assert.Same(t, file1, file2)

	// Test non-existent file
	file3, err := pc.Parse(filepath.Join(tmpDir, "nonexistent.go"))
	assert.Error(t, err)
	assert.Nil(t, file3)
}
