package result

// Result is a container that might contain a value.
type Result[T any, E error] struct {
	val *T
	err *E
}

func Ok[T any](val T) *Result[T, error] {
	return &Result[T, error]{
		val: &val,
	}
}

func Err[T any, E error](val E) *Result[T, E] {
	return &Result[T, E]{
		val: nil,
		err: &val,
	}
}

// IsOk returns true if the Result is ok
func (r Result[T, E]) IsOk() bool {
	return r.val != nil
}

// IsErr returns true if the Result is an error
func (r Result[T, E]) IsErr() bool {
	return r.val == nil

}

// Yank will pull a value out of Result, panicking if it is nil.
func (r Result[T, E]) Yank() T {
	if r.IsErr() {
		err := r.err
		panic("error:" + (*err).Error())
	}

	return *r.val
}

// Match will call ok or err depending on whether the Result is an `err` or not.
func (r Result[T, E]) Match(ok func(T), err func(E)) {
	if r.IsErr() {
		err(*r.err)
	} else {
		ok(*r.val)
	}
}
