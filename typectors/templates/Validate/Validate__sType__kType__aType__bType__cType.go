package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validate__Combine3__sType__kType__aType__bType__cType(
	ka kType, va Validate_aType_idx_kType,
	kb kType, vb Validate_bType_idx_kType,
	kc kType, vc Validate_cType_idx_kType,
	f func(aType, bType, cType) Validate_sType_idx_kType,
) Validate_sType_idx_kType {
	return func() (valuS sType, errsS Errors_kType) {
		valuA, errsA := va()
		valuB, errsB := vb()
		valuC, errsC := vc()
		if errsA.IsEmpty() && errsB.IsEmpty() && errsC.IsEmpty() {
			valuS, errsS = f(valuA, valuB, valuC)()
		} else {
			errs := Errors__New__kType()
			errs.PutAll(errsA)
			errs.PutAll(errsB)
			errs.PutAll(errsC)
			errsS = errs
		}
		return
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validate__Map3__sType__kType__aType__bType__cType(
	ka kType, va Validate_aType_idx_kType,
	kb kType, vb Validate_bType_idx_kType,
	kc kType, vc Validate_cType_idx_kType,
	f func(aType, bType, cType) sType,
) Validate_sType_idx_kType {
	return Validate__Combine3__sType__kType__aType__bType__cType(ka, va, kb, vb, kc, vc, func(a aType, b bType, c cType) Validate_sType_idx_kType {
		return func() (sType, Errors_kType) {
			return f(a, b, c), nil
		}
	})
}
