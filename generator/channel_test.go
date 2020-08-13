package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaa110/goterator/generator"
)

func TestChannelGenerator(t *testing.T) {
	assert := assert.New(t)
	length := 6
	channel := make(chan interface{}, length)

	generator := generator.NewChannel(channel)

	for i := 0; i < length; i++ {
		channel <- i
		ok := generator.Next()
		assert.True(ok)
		assert.Equal(i, generator.Value())
	}

	close(channel)

	ok := generator.Next()
	assert.False(ok)
}
