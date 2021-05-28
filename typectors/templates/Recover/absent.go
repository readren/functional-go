package fung

type aType struct{}

type Errors_aType map[aType]error

type errorsCompanion struct{}

var Errors errorsCompanion

func (companion errorsCompanion) New__aType() Errors_aType {
	panic(1)
}
