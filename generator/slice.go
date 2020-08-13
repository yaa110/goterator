package generator

// Slice implements `Generator` interface for slices.
type Slice struct {
	elements []interface{}
	value    interface{}
}

// NewSlice creates a new instance of `Slice` for `elements`.
func NewSlice(elements []interface{}) *Slice {
	return &Slice{elements, nil}
}

// Next prepares a new element, and returns `false` at the end of iteration.
// This method consumes (shifts) the elements from the slice.
func (g *Slice) Next() bool {
	if len(g.elements) == 0 {
		return false
	}
	g.value, g.elements = g.elements[0], g.elements[1:]
	return true
}

// Value returns the value of prepared element.
func (g *Slice) Value() interface{} {
	return g.value
}
