package fung

// #importAnchor

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"eType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func ValiResu__Combine5__aType__bType__cType__dType__eType(
	ka kType, va ValiResu_aType_idx_kType,
	kb kType, vb ValiResu_bType_idx_kType,
	kc kType, vc ValiResu_cType_idx_kType,
	kd kType, vd ValiResu_dType_idx_kType,
	ke kType, ve ValiResu_eType_idx_kType,
	f func(aType, bType, cType, dType, eType) ValiResu_sType_idx_kType,
) (vs ValiResu_sType_idx_kType) {
	if va.Errs.IsEmpty() && vb.Errs.IsEmpty() && vc.Errs.IsEmpty() {
		vs = f(va.Val, vb.Val, vc.Val, vd.Val, ve.Val)
	} else {
		errs := Errors__New__kType()
		errs.PutAll(va.Errs)
		errs.PutAll(vb.Errs)
		errs.PutAll(vc.Errs)
		errs.PutAll(vd.Errs)
		errs.PutAll(ve.Errs)
		vs.Errs = errs
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"eType", "type":"kType"}] }
// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func ValiResu__Map5__aType__bType__cType__dType__eType(
	ka kType, va ValiResu_aType_idx_kType,
	kb kType, vb ValiResu_bType_idx_kType,
	kc kType, vc ValiResu_cType_idx_kType,
	kd kType, vd ValiResu_dType_idx_kType,
	ke kType, ve ValiResu_eType_idx_kType,
	f func(aType, bType, cType, dType, eType) sType,
) (vs ValiResu_sType_idx_kType) {
	return ValiResu__Combine5__aType__bType__cType__dType__eType(ka, va, kb, vb, kc, vc, kd, vd, ke, ve, func(a aType, b bType, c cType, d dType, e eType) ValiResu_sType_idx_kType {
		return ValiResu_sType_idx_kType{f(a, b, c, d, e), nil}
	})
}
