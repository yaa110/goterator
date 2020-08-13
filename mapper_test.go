package goterator_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaa110/goterator"
	"github.com/yaa110/goterator/generator"
)

func TestMap(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	generator := generator.NewSlice(elements)

	doubled := goterator.New(generator).Map(func(e interface{}) interface{} {
		return e.(int) * 2
	}).Collect()

	assert.Equal([]interface{}{0, 2, 4, 6, 8, 10}, doubled)
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	generator := generator.NewSlice(elements)

	odds := goterator.New(generator).Filter(func(e interface{}) bool {
		return e.(int)%2 != 0
	}).Collect()

	assert.Equal([]interface{}{1, 3, 5}, odds)
}

func TestTake(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	generator := generator.NewSlice(elements)

	odds := goterator.New(generator).Take(3).Collect()

	assert.Equal([]interface{}{0, 1, 2}, odds)
}

func TestSkip(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	generator := generator.NewSlice(elements)

	odds := goterator.New(generator).Skip(3).Collect()

	assert.Equal([]interface{}{3, 4, 5}, odds)
}

func TestSkipWhile(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 2, 5}
	generator := generator.NewSlice(elements)

	values := goterator.New(generator).SkipWhile(func(e interface{}) bool {
		return e.(int) < 3
	}).Collect()

	assert.Equal([]interface{}{3, 2, 5}, values)
}

func TestTakeWhile(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 2, 5}
	generator := generator.NewSlice(elements)

	values := goterator.New(generator).TakeWhile(func(e interface{}) bool {
		return e.(int) < 3
	}).Collect()

	assert.Equal([]interface{}{0, 1, 2}, values)
}

func TestChain(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{"test1", "word", "TeSt2", "WorD", "TEST3", "WORD", "Test4", "Word"}
	generator := generator.NewSlice(elements)

	values := goterator.New(generator).
		Map(func(e interface{}) interface{} {
			return strings.ToLower(e.(string))
		}).
		Filter(func(s interface{}) bool {
			return s.(string) != "word"
		}).
		Take(3).
		Map(func(e interface{}) interface{} {
			return strings.ToUpper(e.(string))
		}).
		Skip(1).
		Take(1).
		Collect()

	assert.Equal([]interface{}{"TEST2"}, values)
}
