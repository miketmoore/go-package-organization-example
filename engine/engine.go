package engine

type Engine struct {
	id int
}

func (e Engine) Id() int {
	return e.id
}

func New(id int) Engine {
	return Engine{id}
}
