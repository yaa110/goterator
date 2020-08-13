package goterator

import "github.com/yaa110/goterator/generator"

// Iterator represents an iterator object to apply map and reduce functions.
type Iterator struct {
	generator generator.Generator
	mappers   []interface{}
}

// New creates a new instance of `Iterator`.
func New(gen generator.Generator) *Iterator {
	return &Iterator{
		generator: gen,
		mappers:   make([]interface{}, 0),
	}
}
