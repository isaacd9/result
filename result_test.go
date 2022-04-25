package result

import (
	"errors"
	"testing"
)

func TestOk(t *testing.T) {
	r := Ok("hi there")
	val := r.Yank()
	if val != "hi there" {
		t.Fatalf("[unexpected] wanted no value out of Result[T], got: %v", val)
	}

	r.Match(func(s string) {
		if s != "hi there" {
			t.Fatalf("[unexpected] wanted no value out of Result[T], got: %v", s)
		}
	}, func(err error) {
		if err != nil {
			t.Fatalf("[unexpected] wanted no err, got: %+v", err)
		}
	})
}

func TestErr(t *testing.T) {
	r := Err[string](errors.New("oh no"))
	r.Match(func(s string) {
		if s != "" {
			t.Fatalf("[unexpected] wanted no value out of Result[T], got: %v", s)
		}
	}, func(err error) {
		if err == nil {
			t.Fatalf("[unexpected] wanted an err")
		}
	})
}
