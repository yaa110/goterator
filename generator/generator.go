package generator

// Generator is an interface to be implemented by element providers.
type Generator interface {
	// Next prepares a new element, and returns `false` at the end of iteration.
	Next() bool

	// Value returns the value of prepared element.
	Value() interface{}
}
