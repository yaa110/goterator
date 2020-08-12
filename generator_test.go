package goterator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaa110/goterator"
)

type TestGenerator struct {
	i      int
	length int
}

func (g *TestGenerator) Next() (interface{}, error) {
	if g.i == g.length {
		return nil, goterator.End()
	}
	element := g.i
	g.i++
	return element, nil
}

func TestSliceGenerator(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	length := len(elements)

	generator := goterator.NewSliceGenerator(elements)

	for i := 0; i < length; i++ {
		element, _ := generator.Next()
		assert.Equal(i, element)
	}
	_, err := generator.Next()
	assert.NotNil(err)
}

func TestChannelGenerator(t *testing.T) {
	assert := assert.New(t)
	length := 6
	channel := make(chan interface{}, length)

	generator := goterator.NewChannelGenerator(channel)

	for i := 0; i < length; i++ {
		channel <- i
		element, _ := generator.Next()
		assert.Equal(i, element)
	}

	close(channel)

	_, err := generator.Next()
	assert.NotNil(err)
}

func TestCustomGenerator(t *testing.T) {
	assert := assert.New(t)
	length := 6

	generator := &TestGenerator{0, length}

	for i := 0; i < length; i++ {
		element, _ := generator.Next()
		assert.Equal(i, element)
	}
	_, err := generator.Next()
	assert.NotNil(err)
}
