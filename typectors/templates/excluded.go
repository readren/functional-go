package fung

// The type of the elements contained by collection types
type eType = struct{}

// The type of fiel where the successful value is stored
type sType = struct{}

// The first type parameter of the polymorphic methods
type aType = struct{}

// The second type parameter of the polymorphic methods
type bType = struct{}

type Try_aType func() aType

func (ta Try_aType) Catch() (a aType, err interface{}) {
	panic("This template line should have been removed")
}

type ValiResu_aType struct {
	Val aType
	Err interface{}
}

func ValiResu_aType__Try(f func() aType) (vrs ValiResu_aType) {
	panic("This template line should have been removed")
}

type Validate_aType func() (aType, interface{})

func Validate_aType__Try(f Try_aType) Validate_aType {
	panic("This template line should have been removed")
}
