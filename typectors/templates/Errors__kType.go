package fung

// #importAnchor

type Errors_kType map[kType]error

func Errors__New__kType() Errors_kType {
	return make(map[kType]error)
}
