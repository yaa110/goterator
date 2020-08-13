package generator

// Channel implements `Generator` interface for channels.
type Channel struct {
	channel <-chan interface{}
	value   interface{}
}

// NewChannel creates a new instance of `Channel` for `channel`.
func NewChannel(channel <-chan interface{}) *Channel {
	return &Channel{channel, nil}
}

// Next prepares a new element, and returns `false` at the end of iteration.
func (g *Channel) Next() bool {
	var ok bool
	g.value, ok = <-g.channel
	return ok
}

// Value returns the value of prepared element.
func (g *Channel) Value() interface{} {
	return g.value
}
