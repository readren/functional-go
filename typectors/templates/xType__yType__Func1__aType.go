package fung

// #dependsOn {"typeCtor":"Func1", "baseTArgs":[{"type":"xType"},{"type":"yType"}]}

// #importAnchor

// #dependsOn {"typeCtor":"Func1", "baseTArgs":[{"type":"yType"},{"type":"aType"}]}
// #dependsOn {"typeCtor":"Func1", "baseTArgs":[{"type":"xType"},{"type":"aType"}]}
func (f FuncFrom_xType_to_yType) Composed__aType(g FuncFrom_yType_to_aType) FuncFrom_xType_to_aType {
	return func(x xType) aType {
		return g(f(x))
	}
}

// #dependsOn {"typeCtor":"Recover", "funcTArgs":[{"type":"aType"}]}
func (fxy FuncFrom_xType_to_yType) Guarded__aType(key aType, guard Guard) func(xType) (yType, Errors_aType) {
	return func(x xType) (y yType, errs Errors_aType) {
		defer recover__catch__aType(&errs, guard, key)
		y = fxy(x)
		return
	}
}

// TODO change the name, move or delete, because it is not cohesive.
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs":[{"type":"yType"},{"type":"aType"}]}
// #dependsOn {"typeCtor":"Func1", "baseTArgs":[{"type":"xType"},{"type":"ValiResu_yType_aType"}] }
func (f FuncFrom_xType_to_yType) Validated__aType(key aType) FuncFrom_xType_to_ValiResu_yType_aType {
	return func(x xType) (vr ValiResu_yType_aType) {
		vr.Val = f(x)
		return
	}
}
