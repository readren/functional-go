package fung

// #importAnchor

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func (companion valiResuCompanion) Combine2__sType__kType__aType__bType(
	ka kType, va ValiResu_aType_by_kType,
	kb kType, vb ValiResu_bType_by_kType,
	f func(aType, bType) ValiResu_sType_by_kType,
) (vs ValiResu_sType_by_kType) {
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
func (companion valiResuCompanion) Map2__sType__kbType__aType__bType(
	ka kType, va ValiResu_aType_by_kType,
	kb kType, vb ValiResu_bType_by_kType,
	f func(aType, bType) sType,
) (vs ValiResu_sType_by_kType) {
	return companion.Combine2__sType__kType__aType__bType(ka, va, kb, vb, func(a aType, b bType) ValiResu_sType_by_kType {
		return ValiResu_sType_by_kType{f(a, b), nil}
	})
}
