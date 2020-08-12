package goterator

// Iterator represents an iterator object to apply map and reduce functions.
type Iterator struct {
	generator Generator
	mappers   []interface{}
}

// New creates a new instance of `Iterator`.
func New(generator Generator) *Iterator {
	return &Iterator{
		generator: generator,
		mappers:   make([]interface{}, 0),
	}
}
