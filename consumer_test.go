package goterator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaa110/goterator"
	"github.com/yaa110/goterator/generator"
)

func TestForEach(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	var values []interface{}
	iterator.ForEach(func(e interface{}) {
		values = append(values, e)
	})
	assert.Equal(elements, values)
}

func TestCollect(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	values := iterator.Collect()
	assert.Equal(elements, values)
}

func TestReduce(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	sum := iterator.Reduce(0, func(state, e interface{}) interface{} {
		return state.(int) + e.(int)
	})
	assert.Equal(15, sum)
}

func TestFind(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	three, err := iterator.Find(func(e interface{}) bool {
		return e.(int) == 3
	})
	assert.Equal(3, three)
	assert.Nil(err)

	gen = generator.NewSlice(elements)
	iterator = goterator.New(gen)

	_, err = iterator.Find(func(e interface{}) bool {
		return e.(int) == 6
	})
	assert.NotNil(err)
}

func TestMin(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, -1, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	min := iterator.Min(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	assert.Equal(-1, min)
}

func TestMax(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 7, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	max := iterator.Max(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	assert.Equal(7, max)
}

func TestAll(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, -1, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	all := iterator.All(func(e interface{}) bool {
		return e.(int) < 6
	})
	assert.True(all)

	gen = generator.NewSlice(make([]interface{}, 0))
	iterator = goterator.New(gen)

	all = iterator.All(func(e interface{}) bool {
		return e.(int) < 6
	})
	assert.True(all)

	gen = generator.NewSlice(elements)
	iterator = goterator.New(gen)

	all = iterator.All(func(e interface{}) bool {
		return e.(int) > 3
	})
	assert.False(all)
}

func TestAny(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, -1, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	any := iterator.Any(func(e interface{}) bool {
		return e.(int) < 0
	})
	assert.True(any)

	gen = generator.NewSlice(make([]interface{}, 0))
	iterator = goterator.New(gen)

	any = iterator.Any(func(e interface{}) bool {
		return e.(int) < 6
	})
	assert.False(any)

	gen = generator.NewSlice(elements)
	iterator = goterator.New(gen)

	any = iterator.Any(func(e interface{}) bool {
		return e.(int) > 6
	})
	assert.False(any)
}

func TestLast(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 7, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	last := iterator.Last()
	assert.Equal(5, last)
}

func TestNth(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 7, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	nth, err := iterator.Nth(5)
	assert.Equal(7, nth)
	assert.Nil(err)

	gen = generator.NewSlice(elements)
	iterator = goterator.New(gen)

	_, err = iterator.Nth(len(elements))
	assert.NotNil(err)
}

func TestCount(t *testing.T) {
	assert := assert.New(t)
	elements := []interface{}{0, 1, 2, 3, 4, 7, 5}
	gen := generator.NewSlice(elements)
	iterator := goterator.New(gen)

	length := iterator.Count()
	assert.Equal(len(elements), length)

	gen = generator.NewSlice(make([]interface{}, 0))
	iterator = goterator.New(gen)

	length = iterator.Count()
	assert.Zero(length)
}
