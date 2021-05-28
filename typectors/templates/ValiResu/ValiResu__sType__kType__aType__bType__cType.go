package fung

// #importAnchor

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func (companion valiResuCompanion) Combine3__sType__kType__aType__bType__cType(
	ka kType, va ValiResu_aType_by_kType,
	kb kType, vb ValiResu_bType_by_kType,
	kc kType, vc ValiResu_cType_by_kType,
	f func(aType, bType, cType) ValiResu_sType_by_kType,
) (vs ValiResu_sType_by_kType) {
	if va.Errs.IsEmpty() && vb.Errs.IsEmpty() && vc.Errs.IsEmpty() {
		vs = f(va.Val, vb.Val, vc.Val)
	} else {
		errs := Errors__New__kType()
		errs.PutAll(va.Errs)
		errs.PutAll(vb.Errs)
		errs.PutAll(vc.Errs)
		vs.Errs = errs
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func (companion valiResuCompanion) Map3__sType__kType__aType__bType__cType(
	ka kType, va ValiResu_aType_by_kType,
	kb kType, vb ValiResu_bType_by_kType,
	kc kType, vc ValiResu_cType_by_kType,
	f func(aType, bType, cType) sType,
) (vs ValiResu_sType_by_kType) {
	return companion.Combine3__sType__kType__aType__bType__cType(ka, va, kb, vb, kc, vc, func(a aType, b bType, c cType) ValiResu_sType_by_kType {
		return ValiResu_sType_by_kType{f(a, b, c), nil}
	})
}
