package fung

type aType struct{}
type bType struct{}
type cType struct{}
type dType struct{}
type eType struct{}
type sType struct{}
type kType struct{}

type Errors_aType map[aType]error
type Errors_bType map[bType]error
type Errors_kType map[kType]error

func Errors__New__kType() Errors_kType {
	panic(1)
}
func (a Errors_kType) PutAll(b map[kType]error) {
	panic(1)
}

type Guard func(error) bool

type ValiResu_sType_aType struct {
	Val  sType
	Errs Errors_aType
}
type ValiResu_aType_bType struct {
	Val  aType
	Errs Errors_bType
}
type ValiResu_aType_kType struct {
	Val  aType
	Errs Errors_kType
}
type ValiResu_bType_kType struct {
	Val  bType
	Errs Errors_kType
}
type ValiResu_cType_kType struct {
	Val  cType
	Errs Errors_kType
}
type ValiResu_dType_kType struct {
	Val  dType
	Errs Errors_kType
}
type ValiResu_eType_kType struct {
	Val  eType
	Errs Errors_kType
}

func (errs Errors_kType) IsEmpty() bool {
	panic(1)
}

type FuncFrom_sType_to_aType func(sType) aType

type FuncFrom_sType_to_ValiResu_aType_kType func(sType, ValiResu_aType_kType)

func (f FuncFrom_sType_to_aType) Guarded__kType(kType, Guard) func(sType) (aType, Errors_kType)

func (fsa FuncFrom_sType_to_ValiResu_aType_kType) Guarded__kType(key kType, guard Guard) func(sType) (ValiResu_aType_kType, Errors_kType) {
	panic(1)
}
