// Package engine is the core, generic functionality that other packages should use
package engine

// Engine is a simple struct that encapsulates an ID
type Engine struct {
	id int
}

// Id is a getter for the internal ID value
func (e Engine) Id() int {
	return e.id
}

// New returns an instance of Engine, with the specified ID
func New(id int) Engine {
	return Engine{id}
}
