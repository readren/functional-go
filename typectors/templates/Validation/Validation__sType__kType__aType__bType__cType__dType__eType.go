package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"eType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validation__Combine5__sType__kType__aType__bType__cType__dType__eType(
	ka kType, va Validation_aType_idx_kType,
	kb kType, vb Validation_bType_idx_kType,
	kc kType, vc Validation_cType_idx_kType,
	kd kType, vd Validation_dType_idx_kType,
	ke kType, ve Validation_eType_idx_kType,
	f func(aType, bType, cType, dType, eType) Validation_sType_idx_kType,
) Validation_sType_idx_kType {
	return func() (valuS sType, errsS Errors_kType) {
		valuA, errsA := va()
		valuB, errsB := vb()
		valuC, errsC := vc()
		valuD, errsD := vd()
		valuE, errsE := ve()
		if errsA.IsEmpty() && errsB.IsEmpty() && errsC.IsEmpty() && errsD.IsEmpty() && errsE.IsEmpty() {
			valuS, errsS = f(valuA, valuB, valuC, valuD, valuE)()
		} else {
			errs := Errors__New__kType()
			errs.PutAll(errsA)
			errs.PutAll(errsB)
			errs.PutAll(errsC)
			errs.PutAll(errsD)
			errs.PutAll(errsE)
			errsS = errs
		}
		return
	}
}

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"eType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validation__Map5__sType__kType__aType__bType__cType__dType__eType(
	ka kType, va Validation_aType_idx_kType,
	kb kType, vb Validation_bType_idx_kType,
	kc kType, vc Validation_cType_idx_kType,
	kd kType, vd Validation_dType_idx_kType,
	ke kType, ve Validation_eType_idx_kType,
	f func(aType, bType, cType, dType, eType) sType,
) Validation_sType_idx_kType {
	return Validation__Combine5__sType__kType__aType__bType__cType__dType__eType(ka, va, kb, vb, kc, vc, kd, vd, ke, ve, func(a aType, b bType, c cType, d dType, e eType) Validation_sType_idx_kType {
		return func() (sType, Errors_kType) {
			return f(a, b, c, d, e), nil
		}
	})
}
