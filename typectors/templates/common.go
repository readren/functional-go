package fung

import "errors"

type Guard func(error) bool

// Builds a sentinel guard useful for the `guard` parameter of `Try` methods.
// The returned guard recovers from panics whose argument is an `Error` that has the received `sentinel` error in the its chain.
func SentinelGuard(sentinel error) Guard {
	return func(err error) bool {
		if err, ok := err.(error); ok && errors.Is(err, sentinel) {
			return true
		}
		return false
	}
}

func catch(errPtr *error, guard Guard) {
	check := func(any interface{}) bool {
		err, ok := any.(error)
		if ok && guard(err) {
			*errPtr = err
			return true
		}
		return false
	}
	if any := recover(); check(any) {
		// do nothing
	}
}
