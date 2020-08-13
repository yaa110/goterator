package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaa110/goterator/generator"
)

func TestSlice(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	length := len(elements)

	generator := generator.NewSlice(elements)

	for i := 0; i < length; i++ {
		ok := generator.Next()
		assert.True(ok)
		assert.Equal(i, generator.Value())
	}
	ok := generator.Next()
	assert.False(ok)
}
