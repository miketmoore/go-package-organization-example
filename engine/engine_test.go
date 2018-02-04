package engine_test

import (
	"testing"

	"github.com/miketmoore/test/engine"
	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	e := engine.Engine{1}
	assert.Equal(t, 1, e.Id)
}
