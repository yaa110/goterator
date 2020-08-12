package goterator

// Generator is an interface to be implemented by element providers.
type Generator interface {
	// Next generates a new element, or returns a non-nil error to end the iteration.
	Next() (interface{}, error)
}

// SliceGenerator implements `Generator` interface for slices.
type SliceGenerator struct {
	elements []interface{}
}

// ChannelGenerator implements `Generator` interface for channels.
type ChannelGenerator struct {
	channel <-chan interface{}
}

// NewSliceGenerator creates a new instance of `SliceGenerator` for `elements`.
func NewSliceGenerator(elements []interface{}) *SliceGenerator {
	return &SliceGenerator{elements}
}

// Next returns next element of slice or a non-nil error to end the iteration.
// This method consumes (shifts) the elements from the slice.
func (g *SliceGenerator) Next() (interface{}, error) {
	if len(g.elements) == 0 {
		return nil, End()
	}
	var element interface{}
	element, g.elements = g.elements[0], g.elements[1:]
	return element, nil
}

// NewChannelGenerator creates a new instance of `ChannelGenerator` for `channel`.
func NewChannelGenerator(channel <-chan interface{}) *ChannelGenerator {
	return &ChannelGenerator{channel}
}

// Next returns next element read from channel or a non-nil error to end the iteration.
func (g *ChannelGenerator) Next() (interface{}, error) {
	element, ok := <-g.channel
	if !ok {
		return nil, End()
	}
	return element, nil
}
