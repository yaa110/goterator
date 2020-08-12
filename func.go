package goterator

// IterFunc gets an element.
type IterFunc = func(element interface{})

// MapFunc get an element and returns a converted element.
type MapFunc = func(element interface{}) interface{}

// ReduceFunc gets a state and an element, then returns the state.
type ReduceFunc = func(state, element interface{}) interface{}

// PredicateFunc gets an element and returns true if the element should be yielded.
type PredicateFunc = func(element interface{}) bool

// LessFunc returns true if first < second
type LessFunc = func(first, second interface{}) bool
