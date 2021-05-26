package fung

type aType struct{}
type bType struct{}
type cType struct{}
type dType struct{}
type eType struct{}
type kType struct{}
type sType struct{}
type xType struct{}
type yType struct{}

type Errors_aType map[aType]error
type Errors_bType map[bType]error

func Errors__New__aType() Errors_aType {
	panic(1)
}

type FlawedFuncFrom_yType_to_aType func(yType) (aType, error)
type FlawedFuncFrom_xType_to_aType func(xType) (aType, error)

type FuncFrom_sType_to_aType func(s sType) aType
type FuncFrom_yType_to_aType func(y yType) aType
type FuncFrom_xType_to_aType func(x xType) aType

func (gs Giver_sType) Guarded__kType(key kType, guard Guard) func() (sType, Errors_kType) {
	panic(1)
}

func (fsa FuncFrom_sType_to_aType) Guarded__kType(key kType, guard Guard) func(sType) (aType, Errors_kType) {
	panic(1)
}

type FuncFrom_sType_to_Validate_aType_kType func(s sType) Validate_aType_kType

func (fsa FuncFrom_sType_to_Validate_aType_kType) Guarded__kType(key kType, guard Guard) func(sType) (Validate_aType_kType, Errors_kType) {
	panic(1)
}

type FuncFrom_sType_to_ValiResu_aType_kType func(s sType) ValiResu_aType_kType

func (fsvrak FuncFrom_sType_to_ValiResu_aType_kType) Guarded__kType(key kType, guard Guard) func(sType) (ValiResu_aType_kType, Errors_kType) {
	panic(1)
}

type FuncFrom_xType_to_ValiResu_yType_aType func(xType) ValiResu_yType_aType

type Stream_eType func() (eType, Stream_eType)

func Stream__Single__eType(e eType) Stream_eType {
	panic(1)
}

type Validate_sType_aType func() (sType, Errors_aType)
type Validate_aType_bType func() (aType, Errors_bType)
type Validate_aType_kType func() (aType, Errors_kType)
type Validate_bType_kType func() (bType, Errors_kType)
type Validate_cType_kType func() (cType, Errors_kType)
type Validate_dType_kType func() (dType, Errors_kType)
type Validate_eType_kType func() (eType, Errors_kType)
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
type ValiResu_yType_aType struct {
	Val  yType
	Errs Errors_aType
}
