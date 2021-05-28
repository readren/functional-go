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
type Errors_cType map[cType]error
type Errors_kType map[kType]error

func Errors__New__kType() Errors_kType {
	panic(1)
}
func (a Errors_kType) PutAll(b map[kType]error) {
	panic(1)
}

type Guard func(error) bool

type Validation_sType_idx_aType func() (sType, Errors_aType)
type Validation_aType_idx_bType func() (aType, Errors_bType)
type Validation_aType_idx_kType func() (aType, Errors_kType)
type Validation_bType_idx_kType func() (bType, Errors_kType)
type Validation_cType_idx_kType func() (cType, Errors_kType)
type Validation_dType_idx_kType func() (dType, Errors_kType)
type Validation_eType_idx_kType func() (eType, Errors_kType)

func (errs Errors_kType) IsEmpty() bool {
	panic(1)
}

type FuncFrom_sType_to_aType func(sType) aType

type FuncFrom_sType_to_Validation_aType_idx_kType func(sType, Validation_aType_idx_kType)

func (f FuncFrom_sType_to_aType) Guarded__kType(kType, Guard) func(sType) (aType, Errors_kType)

func (fsa FuncFrom_sType_to_Validation_aType_idx_kType) Guarded__kType(key kType, guard Guard) func(sType) (Validation_aType_idx_kType, Errors_kType) {
	panic(1)
}

type Stream_sType func() (sType, Stream_sType)

type Validation_Slice_aType_idx_kType func() ([]aType, Errors_kType)

type Stream_Validation_sType_idx_kType func() (Validation_sType_idx_kType, Stream_Validation_sType_idx_kType)

type Validation_slice_sType_idx_kType func() ([]sType, Errors_kType)

func Validation__Traverse__Validation_sType_idx_kType__kType__sType(stream Stream_Validation_sType_idx_kType, f func(Validation_sType_idx_kType) Validation_sType_idx_kType) Validation_slice_sType_idx_kType {
	panic(1)
}

func Func1__Identity__Validation_sType_idx_kType(v Validation_sType_idx_kType) Validation_sType_idx_kType {
	return v
}
