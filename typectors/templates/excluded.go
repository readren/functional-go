package fung

// The type of the elements contained by collection types
type eType = struct{}

// The type of fiel where the successful value is stored
type sType = struct{}

// The first type parameter of the polymorphic methods
type aType = struct{}

// The second type parameter of the polymorphic methods
type bType = struct{}

type xType = struct{}
type yType = struct{}

type ValiResu_aType struct {
	Val aType
	Err interface{}
}

type Validate_aType func() (aType, interface{})

type FuncFrom_sType_to_aType func(s sType) aType

func (fsa FuncFrom_sType_to_aType) Try(s sType) (aType, interface{}) {
	panic("This template line should have been removed")
}
