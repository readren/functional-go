package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validate__Combine4__aType__bType__cType__dType(
	ka kType, va Validate_aType_idx_kType,
	kb kType, vb Validate_bType_idx_kType,
	kc kType, vc Validate_cType_idx_kType,
	kd kType, vd Validate_dType_idx_kType,
	f func(aType, bType, cType, dType) Validate_sType_idx_kType,
) Validate_sType_idx_kType {
	return func() (valuS sType, errsS Errors_kType) {
		valuA, errsA := va()
		valuB, errsB := vb()
		valuC, errsC := vc()
		valuD, errsD := vd()
		if errsA.IsEmpty() && errsB.IsEmpty() && errsC.IsEmpty() && errsD.IsEmpty() {
			valuS, errsS = f(valuA, valuB, valuC, valuD)()
		} else {
			errs := Errors__New__kType()
			errs.PutAll(errsA)
			errs.PutAll(errsB)
			errs.PutAll(errsC)
			errs.PutAll(errsD)
			errsS = errs
		}
		return
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validate__Map4__aType__bType__cType__dType(
	ka kType, va Validate_aType_idx_kType,
	kb kType, vb Validate_bType_idx_kType,
	kc kType, vc Validate_cType_idx_kType,
	kd kType, vd Validate_dType_idx_kType,
	f func(aType, bType, cType, dType) sType,
) Validate_sType_idx_kType {
	return Validate__Combine4__aType__bType__cType__dType(ka, va, kb, vb, kc, vc, kd, vd, func(a aType, b bType, c cType, d dType) Validate_sType_idx_kType {
		return func() (sType, Errors_kType) {
			return f(a, b, c, d), nil
		}
	})
}
