package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validate__Combine2__sType__kType__aType__bType(
	ka kType, va Validate_aType_idx_kType,
	kb kType, vb Validate_bType_idx_kType,
	f func(aType, bType) Validate_sType_idx_kType,
) Validate_sType_idx_kType {
	return func() (valuS sType, errsS Errors_kType) {
		valuA, errsA := va()
		valuB, errsB := vb()
		if errsA.IsEmpty() && errsB.IsEmpty() {
			valuS, errsS = f(valuA, valuB)()
		} else {
			errs := Errors__New__kType()
			errs.PutAll(errsA)
			errs.PutAll(errsB)
			errsS = errs
		}
		return
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validate__Map2__sType__kType__aType__bType(
	ka kType, va Validate_aType_idx_kType,
	kb kType, vb Validate_bType_idx_kType,
	f func(aType, bType) sType,
) Validate_sType_idx_kType {
	return Validate__Combine2__sType__kType__aType__bType(ka, va, kb, vb, func(a aType, b bType) Validate_sType_idx_kType {
		return func() (sType, Errors_kType) {
			return f(a, b), nil
		}
	})
}
