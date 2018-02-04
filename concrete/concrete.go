// Package concrete uses the engine package
package concrete

import "github.com/miketmoore/test/engine"

// Concrete is a struct that encapsulates an engine.Engine instance and an
// additional name value
type Concrete struct {
	name   string
	engine engine.Engine
}

// New returns an instance of Concrete with the specified values
func New(id int, name string) Concrete {
	e := engine.New(id)
	return Concrete{name, e}
}

// Id is a getter for the internal engine.Engine ID value
func (c Concrete) Id() int {
	return c.engine.Id()
}

// Name is a getter for the name value
func (c Concrete) Name() string {
	return c.name
}
