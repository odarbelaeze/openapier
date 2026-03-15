package schema_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/pkg/schema"
	"github.com/stretchr/testify/assert"
)

func TestLocator_String(t *testing.T) {
	tests := []struct {
		name    string
		locator schema.Locator
		want    string
	}{
		{
			name: "simple locator",
			locator: schema.Locator{
				Path:    "github.com/odarbelaeze/openapier/pkg/schema",
				Package: "models",
				Name:    "User",
			},
			want: "github_com_odarbelaeze_openapier_pkg_schema:models.User",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.locator.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
