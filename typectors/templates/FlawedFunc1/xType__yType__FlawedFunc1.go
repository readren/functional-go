package fung

// #importAnchor

type FlawedFuncFrom_xType_to_yType func(x xType) (yType, error)

// #startOfFuncsWithNoInternalDependants

func (fxy FlawedFuncFrom_xType_to_yType) Fixed(x xType) func() (yType, error) {
	return func() (yType, error) {
		return fxy(x)
	}
}

// #dependsOn {"typeCtor":"Recover"}
func (fxy FlawedFuncFrom_xType_to_yType) Guarded(guard Guard) func(xType) (yType, error) {
	return func(x xType) (y yType, err error) {
		defer recover__catch(&err, guard)
		y, err = fxy(x)
		return
	}
}
