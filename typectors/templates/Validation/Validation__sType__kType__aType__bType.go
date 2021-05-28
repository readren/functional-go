package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validation__Combine2__sType__kType__aType__bType(
	ka kType, va Validation_aType_by_kType,
	kb kType, vb Validation_bType_by_kType,
	f func(aType, bType) Validation_sType_by_kType,
) Validation_sType_by_kType {
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

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validation__Map2__sType__kType__aType__bType(
	ka kType, va Validation_aType_by_kType,
	kb kType, vb Validation_bType_by_kType,
	f func(aType, bType) sType,
) Validation_sType_by_kType {
	return Validation__Combine2__sType__kType__aType__bType(ka, va, kb, vb, func(a aType, b bType) Validation_sType_by_kType {
		return func() (sType, Errors_kType) {
			return f(a, b), nil
		}
	})
}
