package fung

// #importAnchor

// #dependsOn {"typeCtor":"FlawedFunc1", "baseTArgs":[{"type":"yType"},{"type":"aType"}]}
// #dependsOn {"typeCtor":"FlawedFunc1", "baseTArgs":[{"type":"xType"},{"type":"aType"}]}
func (f FlawedFuncFrom_xType_to_yType) Composed__aType(g FlawedFuncFrom_yType_to_aType) FlawedFuncFrom_xType_to_aType {
	return func(x xType) (a aType, err error) {
		var y yType
		y, err = f(x)
		if err == nil {
			a, err = g(y)
		}
		return
	}
}

// #dependsOn {"typeCtor":"Recover", "funcTArgs":[{"type":"aType"}]}
// #dependsOn {"typeCtor":"Errors", "funcTArgs":[{"type":"aType"}]}
func (f FlawedFuncFrom_xType_to_yType) Guarded__aType(key aType, guard Guard) func(xType) (yType, Errors_aType) {
	return func(x xType) (y yType, errs Errors_aType) {
		defer Recover.catch__aType(&errs, guard, key)
		var err error
		y, err = f(x)
		if err != nil {
			errs = Errors_aType{key: err}
		}
		return
	}
}

// TODO change the name, move or delete, because it is not cohesive.
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs":[{"type":"yType"},{"type":"aType"}]}
// #dependsOn {"typeCtor":"Func1", "baseTArgs":[{"type":"xType"},{"type":"ValiResu_yType_by_aType"}] }
// #dependsOn {"typeCtor":"Errors", "funcTArgs":[{"type":"aType"}]}
func (f FlawedFuncFrom_xType_to_yType) Validationd__aType(key aType) FuncFrom_xType_to_ValiResu_yType_aType {
	return func(x xType) (vr ValiResu_yType_by_aType) {
		var err error
		vr.Val, err = f(x)
		if err != nil {
			vr.Errs = Errors_aType{key: err}
		}
		return
	}
}
