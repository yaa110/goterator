package goterator

// EOF implements `error` interface  to indicate the end of iteration.
type EOF struct{}

// Error returns the error message of `EOF`
func (e *EOF) Error() string {
	return "End of iteration"
}

// End creates a new instance of `EOF`.
func End() *EOF {
	return &EOF{}
}
