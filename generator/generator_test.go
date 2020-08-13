package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestGenerator struct {
	i      int
	length int
	value  int
}

func (g *TestGenerator) Next() bool {
	if g.i == g.length {
		return false
	}
	g.value = g.i
	g.i++
	return true
}

func (g *TestGenerator) Value() interface{} {
	return g.value
}

func TestCustomGenerator(t *testing.T) {
	assert := assert.New(t)
	length := 6

	generator := &TestGenerator{0, length, 0}

	for i := 0; i < length; i++ {
		ok := generator.Next()
		assert.True(ok)
		assert.Equal(i, generator.Value())
	}

	ok := generator.Next()
	assert.False(ok)
}
