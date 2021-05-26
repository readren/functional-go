package fung

// #importAnchor

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func ValiResu__Combine4__aType__bType__cType__dType(
	ka kType, va ValiResu_aType_kType,
	kb kType, vb ValiResu_bType_kType,
	kc kType, vc ValiResu_cType_kType,
	kd kType, vd ValiResu_dType_kType,
	f func(aType, bType, cType, dType) ValiResu_sType_kType,
) (vs ValiResu_sType_kType) {
	if va.Errs.IsEmpty() && vb.Errs.IsEmpty() && vc.Errs.IsEmpty() {
		vs = f(va.Val, vb.Val, vc.Val, vd.Val)
	} else {
		errs := Errors__New__kType()
		errs.PutAll(va.Errs)
		errs.PutAll(vb.Errs)
		errs.PutAll(vc.Errs)
		errs.PutAll(vd.Errs)
		vs.Errs = errs
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func ValiResu__Map4__aType__bType__cType__dType(
	ka kType, va ValiResu_aType_kType,
	kb kType, vb ValiResu_bType_kType,
	kc kType, vc ValiResu_cType_kType,
	kd kType, vd ValiResu_dType_kType,
	f func(aType, bType, cType, dType) sType,
) (vs ValiResu_sType_kType) {
	return ValiResu__Combine4__aType__bType__cType__dType(ka, va, kb, vb, kc, vc, kd, vd, func(a aType, b bType, c cType, d dType) ValiResu_sType_kType {
		return ValiResu_sType_kType{f(a, b, c, d), nil}
	})
}
