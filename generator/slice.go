package generator

// slice implements `Generator` interface for slices.
type slice struct {
	elements []interface{}
	value    interface{}
}

// Newslice creates a new instance of `Generator` for slices.
func NewSlice(elements []interface{}) Generator {
	return &slice{elements, nil}
}

// Next prepares a new element, and returns `false` at the end of iteration.
// This method consumes (shifts) the elements from the slice.
func (g *slice) Next() bool {
	if len(g.elements) == 0 {
		return false
	}
	g.value, g.elements = g.elements[0], g.elements[1:]
	return true
}

// Value returns the value of prepared element.
func (g *slice) Value() interface{} {
	return g.value
}
