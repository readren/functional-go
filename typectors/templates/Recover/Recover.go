package fung

import "errors"

type recoverCompanion struct{}

var Recover recoverCompanion

type Guard func(error) bool

// Builds a sentinel guard useful for the `guard` parameter of `Try` methods.
// The returned guard recovers from panics whose argument is an `Error` that has the received `sentinel` error in the its chain.
func (companion recoverCompanion) SentinelGuard(sentinel error) Guard {
	return func(err error) bool {
		return errors.Is(err, sentinel)
	}
}

func (companion recoverCompanion) catch(errPtr *error, guard Guard) {
	if any := recover(); any != nil {
		if err, ok := any.(error); ok && guard(err) {
			*errPtr = err
		} else {
			panic(any)
		}
	}
}
