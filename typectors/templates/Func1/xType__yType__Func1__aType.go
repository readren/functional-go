package fung

// #importAnchor

// #dependsOn {"typeCtor":"Recover", "funcTArgs":[{"type":"aType"}]}
func (fxy FuncFrom_xType_to_yType) Guarded__aType(key aType, guard Guard) func(xType) (yType, Errors_aType) {
	return func(x xType) (y yType, errs Errors_aType) {
		defer recover__catch__aType(&errs, guard, key)
		y = fxy(x)
		return
	}
}

// #startOfFuncsWithNoInternalDependants

// #dependsOn {"typeCtor":"Func1", "baseTArgs":[{"type":"yType"},{"type":"aType"}]}
// #dependsOn {"typeCtor":"Func1", "baseTArgs":[{"type":"xType"},{"type":"aType"}]}
func (f FuncFrom_xType_to_yType) Composed__aType(g FuncFrom_yType_to_aType) FuncFrom_xType_to_aType {
	return func(x xType) aType {
		return g(f(x))
	}
}
