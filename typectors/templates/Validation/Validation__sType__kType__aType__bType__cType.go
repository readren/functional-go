package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func (companion validationCompanion) Combine3__sType__kType__aType__bType__cType(
	ka kType, va Validation_aType_by_kType,
	kb kType, vb Validation_bType_by_kType,
	kc kType, vc Validation_cType_by_kType,
	f func(aType, bType, cType) Validation_sType_by_kType,
) Validation_sType_by_kType {
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

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"bType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"cType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func (companion validationCompanion) Map3__sType__kType__aType__bType__cType(
	ka kType, va Validation_aType_by_kType,
	kb kType, vb Validation_bType_by_kType,
	kc kType, vc Validation_cType_by_kType,
	f func(aType, bType, cType) sType,
) Validation_sType_by_kType {
	return companion.Combine3__sType__kType__aType__bType__cType(ka, va, kb, vb, kc, vc, func(a aType, b bType, c cType) Validation_sType_by_kType {
		return func() (sType, Errors_kType) {
			return f(a, b, c), nil
		}
	})
}