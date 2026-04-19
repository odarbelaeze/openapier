package cache_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTypeDefCache(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a simple go project structure
	err := os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte("module testproject\n\ngo 1.21\n"), 0644)
	require.NoError(t, err)

	pkgDir := filepath.Join(tmpDir, "pkg1")
	err = os.Mkdir(pkgDir, 0755)
	require.NoError(t, err)

	goFile := filepath.Join(pkgDir, "file.go")
	content := []byte("package pkg1\n\ntype User struct {\n\tID int\n\tName string\n}\n")
	err = os.WriteFile(goFile, content, 0644)
	require.NoError(t, err)

	pc := cache.NewParserCache()
	tc := cache.NewTypeDefCache(tmpDir, pc)

	ctx := context.Background()

	// Test Load
	err = tc.Load(ctx, "./pkg1")
	require.NoError(t, err)

	// Test Get
	def, ok := tc.Get("./pkg1", "User")
	assert.True(t, ok)
	require.NotNil(t, def)
	assert.Equal(t, "User", def.TypeSpec.Name.Name)
	assert.Equal(t, "pkg1", def.Locator.Package)
	assert.Equal(t, "User", def.Locator.Name)

	// Test Get non-existent type
	def, ok = tc.Get("./pkg1", "NonExistent")
	assert.False(t, ok)
	assert.Nil(t, def)

	// Test Get non-existent package
	def, ok = tc.Get("./nonexistent", "User")
	assert.False(t, ok)
	assert.Nil(t, def)

	// Test Load already loaded package (should return nil early)
	err = tc.Load(ctx, "./pkg1")
	require.NoError(t, err)

	t.Run("Parser error", func(t *testing.T) {
		mpc := cache.NewMockParserCache(t)
		tc2 := cache.NewTypeDefCache(tmpDir, mpc)

		// We need to trigger packages.Load to return some files
		// but make the mock parser return an error for them.
		mpc.EXPECT().Parse(mock.Anything).Return(nil, assert.AnError)

		err := tc2.Load(ctx, "./pkg1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse file")
	})

	t.Run("Load error", func(t *testing.T) {
		tc3 := cache.NewTypeDefCache("/non/existent/dir", pc)
		err := tc3.Load(ctx, "./pkg1")
		// packages.Load doesn't always return an error if the dir is missing,
		// it might just return empty packages or errors within packages.
		// But let's see what happens.
		assert.Error(t, err)
	})
}
