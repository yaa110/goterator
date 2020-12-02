package generator

// channel implements `Generator` interface for channels.
type channel struct {
	chn   <-chan interface{}
	value interface{}
}

// Newchannel creates a new instance of `Generator` for channels.
func NewChannel(chn <-chan interface{}) Generator {
	return &channel{chn, nil}
}

// Next prepares a new element, and returns `false` at the end of iteration.
func (g *channel) Next() bool {
	var ok bool
	g.value, ok = <-g.chn
	return ok
}

// Value returns the value of prepared element.
func (g *channel) Value() interface{} {
	return g.value
}
