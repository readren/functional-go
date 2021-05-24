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
	Err any
}

type Validate_aType func() (aType, any)

type FuncFrom_sType_to_aType func(s sType) aType

func (fsa FuncFrom_sType_to_aType) Try(s sType, guard Guard) (aType, any) {
	panic("This template line should have been removed")
}

type FuncFrom_sType_to_ValiResu_aType func(s sType) ValiResu_aType

func (fs2vra FuncFrom_sType_to_ValiResu_aType) Try(s sType, guard Guard) (ValiResu_aType, any) {
	panic("This template line should have been removed")
}

type FuncFrom_sType_to_Validate_aType func(s sType) Validate_aType

func (fs2va FuncFrom_sType_to_Validate_aType) Try(s sType, guard Guard) (Validate_aType, any) {
	panic("This template line should have been removed")
}
