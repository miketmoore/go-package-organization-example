package concrete

import "github.com/miketmoore/test/engine"

type Concrete struct {
	name   string
	engine engine.Engine
}

func New(id int, name string) Concrete {
	e := engine.New(id)
	return Concrete{name, e}
}

func (c Concrete) Id() int {
	return c.engine.Id()
}

func (c Concrete) Name() string {
	return c.name
}
