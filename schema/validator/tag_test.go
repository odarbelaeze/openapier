package validator_test

import (
	"testing"

	"github.com/odarbelaeze/openapier/schema/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidatorTagInterface(t *testing.T) {
	var _ validator.ValidatorTag
	assert.True(t, true)
}
