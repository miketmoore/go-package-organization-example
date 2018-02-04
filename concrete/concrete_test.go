package concrete_test

import (
	"testing"

	"github.com/miketmoore/test/concrete"
	"github.com/stretchr/testify/assert"
)

func TestNewConcrete(t *testing.T) {
	c := concrete.New(1, "Mike")
	assert.Equal(t, 1, c.Id())
	assert.Equal(t, "Mike", c.Name())
}
