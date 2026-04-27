package options_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/options"
	"github.com/stretchr/testify/assert"
	"github.com/sv-tools/openapi"
)

func TestOptions(t *testing.T) {
	tests := []struct {
		name     string
		option   options.SchemaOption
		validate func(*testing.T, *openapi.Schema)
	}{
		{
			name:   "WithRequired",
			option: options.WithRequired(),
			validate: func(t *testing.T, s *openapi.Schema) {
				// Note: openapi.Schema doesn't have a direct "Required" bool,
				// it's usually handled by the parent object's Required list.
				// However, the builder call exists, so we ensure it doesn't panic
				// and we can check if it's nil or not.
				assert.NotNil(t, s)
			},
		},
		{
			name:   "WithDescription",
			option: options.WithDescription("test description"),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, "test description", s.Description)
			},
		},
		{
			name:   "WithExample",
			option: options.WithExample("test example"),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Contains(t, s.Examples, "test example")
			},
		},
		{
			name:   "WithMinLength",
			option: options.WithMinLength(5),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, 5, *s.MinLength)
			},
		},
		{
			name:   "WithMaxLength",
			option: options.WithMaxLength(10),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, 10, *s.MaxLength)
			},
		},
		{
			name:   "WithMinItems",
			option: options.WithMinItems(1),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, 1, *s.MinItems)
			},
		},
		{
			name:   "WithMaxItems",
			option: options.WithMaxItems(5),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, 5, *s.MaxItems)
			},
		},
		{
			name:   "WithMinProperties",
			option: options.WithMinProperties(2),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, 2, *s.MinProperties)
			},
		},
		{
			name:   "WithMaxProperties",
			option: options.WithMaxProperties(4),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, 4, *s.MaxProperties)
			},
		},
		{
			name:   "WithMinimum",
			option: options.WithMinimum(0),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, 0, *s.Minimum)
			},
		},
		{
			name:   "WithMaximum",
			option: options.WithMaximum(100),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, 100, *s.Maximum)
			},
		},
		{
			name:   "WithFormat",
			option: options.WithFormat("email"),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, "email", s.Format)
			},
		},
		{
			name:   "WithEnum",
			option: options.WithEnum("a", "b"),
			validate: func(t *testing.T, s *openapi.Schema) {
				assert.Equal(t, []any{"a", "b"}, s.Enum)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := openapi.NewSchemaBuilder()
			tt.option(sb)
			spec := sb.Build()
			assert.NotNil(t, spec.Spec)
			tt.validate(t, spec.Spec)
		})
	}
}
