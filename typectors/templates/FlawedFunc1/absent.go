package fung

type aType struct{}
type xType struct{}
type yType struct{}

type Guard func(error) bool

type Errors_aType map[aType]error

type FlawedFuncFrom_yType_to_aType func(yType) (aType, error)
type FlawedFuncFrom_xType_to_aType func(xType) (aType, error)

func recover__catch(errPtr *error, guard Guard)                           {}
func recover__catch__aType(errsPtr *Errors_aType, guard Guard, key aType) {}

type FuncFrom_xType_to_ValiResu_yType_aType func(xType) ValiResu_yType_idx_aType

type ValiResu_yType_idx_aType struct {
	Val  yType
	Errs Errors_aType
}
