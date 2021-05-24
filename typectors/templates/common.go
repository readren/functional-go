package fung

import "errors"

type any = interface{}

type Guard func(any) bool

// Builds a sentinel guard useful for the `guard` parameter of `Try` methods.
// The returned guard recovers from panics whose argument is an `Error` that has the received `sentinel` error in the its chain.
func SentinelGuard(sentinel error) Guard {
	return func(anything any) bool {
		if err, ok := anything.(error); ok && errors.Is(err, sentinel) {
			return true
		}
		return false
	}
}

func catch(errPtr *any, guard Guard) {
	if any := recover(); guard(any) {
		*errPtr = any
	}
}
