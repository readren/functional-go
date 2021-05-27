package fung

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType"},{"type":"kType"}] }

// #importAnchor

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"},{"type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func (vrs ValiResu_sType_idx_kType) Map__aType(f func(sType) aType) (vra ValiResu_aType_idx_kType) {
	if vrs.Errs.IsEmpty() {
		vra.Val = f(vrs.Val)
	} else {
		vra.Errs = vrs.Errs
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"},{"type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func (vrs ValiResu_sType_idx_kType) Bind__aType(f func(sType) ValiResu_aType_idx_kType) (vra ValiResu_aType_idx_kType) {
	if vrs.Errs.IsEmpty() {
		vra = f(vrs.Val)
	} else {
		vra.Errs = vrs.Errs
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"},{"type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"aType"}], "funcTArgs":[{"type":"kType"}]}
func (vrs ValiResu_sType_idx_kType) GuardMap__aType(key kType, guard Guard, f FuncFrom_sType_to_aType) (vra ValiResu_aType_idx_kType) {
	if vrs.Errs.IsEmpty() {
		vra.Val, vra.Errs = f.Guarded__kType(key, guard)(vrs.Val)
	} else {
		vra.Errs = vrs.Errs
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"},{"type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"ValiResu_aType_idx_kType"}], "funcTArgs":[{"type":"kType"}]}
func (vrs ValiResu_sType_idx_kType) GuardBind__aType(key kType, guard Guard, f FuncFrom_sType_to_ValiResu_aType_idx_kType) (vra ValiResu_aType_idx_kType) {
	if vrs.Errs.IsEmpty() {
		var errs Errors_kType
		vra, errs = f.Guarded__kType(key, guard)(vrs.Val)
		if !errs.IsEmpty() {
			vra.Errs = errs
		}
	} else {
		vra.Errs = vrs.Errs
	}
	return
}
