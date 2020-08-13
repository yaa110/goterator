package goterator

func (iter *Iterator) iter(f PredicateFunc) {
	skipWhileIsOver := false
ElementLoop:
	for {
		ok := iter.generator.Next()
		if !ok {
			return
		}
		element := iter.generator.Value()
		for _, m := range iter.mappers {
			switch mapper := m.(type) {
			case mapFunc:
				element = mapper(element)
			case filterFunc:
				if !mapper(element) {
					continue ElementLoop
				}
			case takeValue:
				if *mapper == 0 {
					return
				}
				*mapper--
			case takeWhileFunc:
				if !mapper(element) {
					return
				}
			case skipValue:
				if *mapper > 0 {
					*mapper--
					continue ElementLoop
				}
			case skipWhileFunc:
				if !skipWhileIsOver && mapper(element) {
					continue ElementLoop
				}
				skipWhileIsOver = true
			}
		}
		if !f(element) {
			return
		}
	}
}

// ForEach consumes elements and runs `f` for each element.
// The iteration is broken if `f` returns false.
func (iter *Iterator) ForEach(f IterFunc) {
	iter.iter(func(element interface{}) bool {
		f(element)
		return true
	})
}

// Collect consumes elements and returns a slice of converted elements.
func (iter *Iterator) Collect() []interface{} {
	var elements []interface{}
	iter.iter(func(element interface{}) bool {
		elements = append(elements, element)
		return true
	})
	return elements
}

// Reduce consumes elements and runs `f` for each element.
// Returns the final state after iteration over all elements.
func (iter *Iterator) Reduce(initialState interface{}, f ReduceFunc) interface{} {
	iter.iter(func(element interface{}) bool {
		initialState = f(initialState, element)
		return true
	})
	return initialState
}

// Find consumes elements and returns the first element that satisfies `f` (returning `true`).
// Returns an `EOF` error if no element is found.
func (iter *Iterator) Find(f PredicateFunc) (interface{}, error) {
	var elem interface{}
	err := End()
	iter.iter(func(element interface{}) bool {
		if f(element) {
			elem = element
			err = nil
			return false
		}
		return true
	})
	return elem, err
}

// Min consumes elements and returns the minimum element.
func (iter *Iterator) Min(f LessFunc) interface{} {
	first := true
	var min interface{}
	iter.iter(func(element interface{}) bool {
		if first || f(element, min) {
			min = element
			first = false
		}
		return true
	})
	return min
}

// Max consumes elements and returns the maximum element.
func (iter *Iterator) Max(f LessFunc) interface{} {
	first := true
	var max interface{}
	iter.iter(func(element interface{}) bool {
		if first || f(max, element) {
			max = element
			first = false
		}
		return true
	})
	return max
}

// All consumes elements and returns true if `f` returns true for all elements.
// Returns `true` for empty iterators.
func (iter *Iterator) All(f PredicateFunc) bool {
	all := true
	iter.iter(func(element interface{}) bool {
		if !f(element) {
			all = false
			return false
		}
		return true
	})
	return all
}

// Any consumes elements and returns true if `f` returns true for at least one element.
// Returns `false` for empty iterators.
func (iter *Iterator) Any(f PredicateFunc) bool {
	any := false
	iter.iter(func(element interface{}) bool {
		if f(element) {
			any = true
			return false
		}
		return true
	})
	return any
}

// Last consumes elements and returns the last element.
func (iter *Iterator) Last() interface{} {
	var last interface{}
	iter.iter(func(element interface{}) bool {
		last = element
		return true
	})
	return last
}

// Nth consumes elements and returns the `n`th element. Indexing starts from `0`.
// Returns an `EOF` error if the length of iterator is less than `n`.
func (iter *Iterator) Nth(n int) (interface{}, error) {
	var nth interface{}
	err := End()
	index := 0
	iter.iter(func(element interface{}) bool {
		if index == n {
			nth = element
			err = nil
			return false
		}
		index++
		return true
	})
	return nth, err
}

// Count consumes elements and returns the length of elements.
func (iter *Iterator) Count() int {
	count := 0
	iter.iter(func(element interface{}) bool {
		count++
		return true
	})
	return count
}
