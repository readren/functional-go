package fung

// #importAnchor

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"bType"}] }
func ValiResu__Successful__aType__bType(a aType) (vr ValiResu_aType_bType) {
	vr.Val = a
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"bType"}] }
func ValiResu__Failed__aType__bType(key bType, err error) (vr ValiResu_aType_bType) {
	vr.Errs = map[bType]error{key: err}
	return
}

// #startOfFuncsWithNoInternalDependants

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func ValiResu__Combine2__aType__bType(
	ka kType, va ValiResu_aType_kType,
	kb kType, vb ValiResu_bType_kType,
	f func(aType, bType) ValiResu_sType_kType,
) (vs ValiResu_sType_kType) {
	if va.Errs.IsEmpty() && vb.Errs.IsEmpty() {
		vs = f(va.Val, vb.Val)
	} else {
		errs := Errors__New__kType()
		errs.PutAll(va.Errs)
		errs.PutAll(vb.Errs)
		vs.Errs = errs
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func ValiResu__Map2__aType__bType(
	ka kType, va ValiResu_aType_kType,
	kb kType, vb ValiResu_bType_kType,
	f func(aType, bType) sType,
) (vs ValiResu_sType_kType) {
	return ValiResu__Combine2__aType__bType(ka, va, kb, vb, func(a aType, b bType) ValiResu_sType_kType {
		return ValiResu_sType_kType{f(a, b), nil}
	})
}
