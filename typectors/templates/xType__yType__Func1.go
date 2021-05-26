package fung

// #importAnchor

type FuncFrom_xType_to_yType func(x xType) yType

// #startOfFuncsWithNoInternalDependants

func (fxy FuncFrom_xType_to_yType) Fixed(x xType) func() yType {
	return func() yType {
		return fxy(x)
	}
}

// #dependsOn {"typeCtor":"Recover"}
func (fxy FuncFrom_xType_to_yType) Guarded(guard Guard) func(xType) (yType, error) {
	return func(x xType) (y yType, err error) {
		defer recover__catch(&err, guard)
		y = fxy(x)
		return
	}
}
