package validator

import (
	"fmt"
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockTag struct {
	tag   string
	err   error
	opts  []options.SchemaOption
	usage string
}

func (m *mockTag) Tag() string   { return m.tag }
func (m *mockTag) Usage() string { return m.usage }
func (m *mockTag) Parse(string, string) ([]options.SchemaOption, error) {
	return m.opts, m.err
}

func TestRegistry_Parse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := NewRegistry()
		r.Register(&mockTag{tag: "required", opts: []options.SchemaOption{options.WithRequired()}})

		opts, err := r.Parse("required", "object")
		require.NoError(t, err)
		assert.Len(t, opts, 1)
	})

	t.Run("invalid tag format", func(t *testing.T) {
		r := NewRegistry()
		_, err := r.Parse("foo=bar=baz", "object")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid validator tag")
	})

	t.Run("tag parse error", func(t *testing.T) {
		r := NewRegistry()
		r.Register(&mockTag{tag: "fail", err: fmt.Errorf("boom")})

		_, err := r.Parse("fail", "object")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse tag fail")
	})

	t.Run("unknown tag ignored", func(t *testing.T) {
		r := NewRegistry()
		opts, err := r.Parse("unknown", "object")
		require.NoError(t, err)
		assert.Empty(t, opts)
	})

	t.Run("multiple tags", func(t *testing.T) {
		r := NewRegistry()
		r.Register(&mockTag{tag: "t1", opts: []options.SchemaOption{options.WithRequired()}})
		r.Register(&mockTag{tag: "t2", opts: []options.SchemaOption{options.WithRequired()}})

		opts, err := r.Parse("t1,t2", "object")
		require.NoError(t, err)
		assert.Len(t, opts, 2)
	})
}
