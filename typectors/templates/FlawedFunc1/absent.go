package fung

type recoverCompanion struct{}

type aType struct{}
type xType struct{}
type yType struct{}

type Guard func(error) bool

type Errors_aType map[aType]error

type FlawedFuncFrom_yType_to_aType func(yType) (aType, error)
type FlawedFuncFrom_xType_to_aType func(xType) (aType, error)

type FuncFrom_xType_to_ValiResu_yType_aType func(xType) ValiResu_yType_by_aType

var Recover recoverCompanion

func (companion recoverCompanion) catch(*error, Guard)
func (companion recoverCompanion) catch__aType(errsPtr *Errors_aType, guard Guard, key aType) {}

type ValiResu_yType_by_aType struct {
	Val  yType
	Errs Errors_aType
}
