package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"eType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validate__Combine5__aType__bType__cType__dType__eType(
	ka kType, va Validate_aType_kType,
	kb kType, vb Validate_bType_kType,
	kc kType, vc Validate_cType_kType,
	kd kType, vd Validate_dType_kType,
	ke kType, ve Validate_eType_kType,
	f func(aType, bType, cType, dType, eType) Validate_sType_kType,
) Validate_sType_kType {
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

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"eType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validate__Map5__aType__bType__cType__dType__eType(
	ka kType, va Validate_aType_kType,
	kb kType, vb Validate_bType_kType,
	kc kType, vc Validate_cType_kType,
	kd kType, vd Validate_dType_kType,
	ke kType, ve Validate_eType_kType,
	f func(aType, bType, cType, dType, eType) sType,
) Validate_sType_kType {
	return Validate__Combine5__aType__bType__cType__dType__eType(ka, va, kb, vb, kc, vc, kd, vd, ke, ve, func(a aType, b bType, c cType, d dType, e eType) Validate_sType_kType {
		return func() (sType, Errors_kType) {
			return f(a, b, c, d, e), nil
		}
	})
}
