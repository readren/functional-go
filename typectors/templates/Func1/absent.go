package fung

type aType struct{}
type xType struct{}
type yType struct{}

type Guard func(error) bool

type Errors_aType map[aType]error

type FuncFrom_yType_to_aType func(y yType) aType
type FuncFrom_xType_to_aType func(x xType) aType

func recover__catch(errPtr *error, guard Guard)                           {}
func recover__catch__aType(errsPtr *Errors_aType, guard Guard, key aType) {}
