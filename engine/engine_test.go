package engine_test

import (
	"testing"

	"github.com/miketmoore/go-package-organization-example/engine"
	"github.com/stretchr/testify/assert"
)

func TestNewEngine(t *testing.T) {
	e := engine.New(1)
	assert.Equal(t, 1, e.Id())
}
