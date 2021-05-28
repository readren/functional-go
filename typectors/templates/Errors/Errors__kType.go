package fung

// #importAnchor

type Errors_kType map[kType]error

func (companion errorsCompanion) New__kType() Errors_kType {
	return make(map[kType]error)
}
