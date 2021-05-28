package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validation__Combine4__sType__kType__aType__bType__cType__dType(
	ka kType, va Validation_aType_by_kType,
	kb kType, vb Validation_bType_by_kType,
	kc kType, vc Validation_cType_by_kType,
	kd kType, vd Validation_dType_by_kType,
	f func(aType, bType, cType, dType) Validation_sType_by_kType,
) Validation_sType_by_kType {
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

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"dType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validation__Map4__sType__kType__aType__bType__cType__dType(
	ka kType, va Validation_aType_by_kType,
	kb kType, vb Validation_bType_by_kType,
	kc kType, vc Validation_cType_by_kType,
	kd kType, vd Validation_dType_by_kType,
	f func(aType, bType, cType, dType) sType,
) Validation_sType_by_kType {
	return Validation__Combine4__sType__kType__aType__bType__cType__dType(ka, va, kb, vb, kc, vc, kd, vd, func(a aType, b bType, c cType, d dType) Validation_sType_by_kType {
		return func() (sType, Errors_kType) {
			return f(a, b, c, d), nil
		}
	})
}
