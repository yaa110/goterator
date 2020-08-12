package goterator

type mapFunc MapFunc
type filterFunc PredicateFunc
type takeWhileFunc PredicateFunc
type skipWhileFunc PredicateFunc
type takeValue *int
type skipValue *int

// Map lazily applies `f` on each element and returns iterator.
func (iter *Iterator) Map(f MapFunc) *Iterator {
	iter.mappers = append(iter.mappers, mapFunc(f))
	return iter
}

// Filter lazily yields element if `f` returns true.
func (iter *Iterator) Filter(f PredicateFunc) *Iterator {
	iter.mappers = append(iter.mappers, filterFunc(f))
	return iter
}

// Take lazily yields `n` elements.
func (iter *Iterator) Take(n int) *Iterator {
	iter.mappers = append(iter.mappers, takeValue(&n))
	return iter
}

// TakeWhile lazily yields element while `f` returns true.
func (iter *Iterator) TakeWhile(f PredicateFunc) *Iterator {
	iter.mappers = append(iter.mappers, takeWhileFunc(f))
	return iter
}

// Skip lazily skips `n` elements.
func (iter *Iterator) Skip(n int) *Iterator {
	iter.mappers = append(iter.mappers, skipValue(&n))
	return iter
}

// SkipWhile lazily skips element while `f` returns true.
// After returning `false` from `f`, this mapper is over.
func (iter *Iterator) SkipWhile(f PredicateFunc) *Iterator {
	iter.mappers = append(iter.mappers, skipWhileFunc(f))
	return iter
}
