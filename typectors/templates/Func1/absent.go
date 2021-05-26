package fung

type aType struct{}
type xType struct{}
type yType struct{}

type Guard func(error) bool

type Errors_aType map[aType]error

type FuncFrom_yType_to_aType func(y yType) aType
type FuncFrom_xType_to_aType func(x xType) aType
type FuncFrom_xType_to_ValiResu_yType_aType func(xType) ValiResu_yType_aType

func recover__catch(errPtr *error, guard Guard)                           {}
func recover__catch__aType(errsPtr *Errors_aType, guard Guard, key aType) {}

type ValiResu_yType_aType struct {
	Val  yType
	Errs Errors_aType
}
