package fung

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType"},{"type":"kType"}] }

// #importAnchor

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"},{"type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func (vs Validate_sType_idx_kType) Mapped__aType(f func(sType) aType) Validate_aType_idx_kType {
	return func() (a aType, errs Errors_kType) {
		var s sType
		s, errs = vs()
		if errs.IsEmpty() {
			a = f(s)
		}
		return
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"},{"type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func (vs Validate_sType_idx_kType) Bound__aType(f func(s sType) Validate_aType_idx_kType) Validate_aType_idx_kType {
	return func() (a aType, errs Errors_kType) {
		var s sType
		s, errs = vs()
		if errs.IsEmpty() {
			va := f(s)
			a, errs = va()
		}
		return
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"},{"type":"kType"}] }
// #dependsOn {"typeCtor":"Recover"}
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"aType"}], "funcTArgs":[{"type":"kType"}]}
func (vs Validate_sType_idx_kType) GuardMapped__aType(key kType, guard Guard, f FuncFrom_sType_to_aType) Validate_aType_idx_kType {
	return func() (a aType, errs Errors_kType) {
		var s sType
		s, errs = vs()
		if errs.IsEmpty() {
			a, errs = f.Guarded__kType(key, guard)(s)
		}
		return
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"},{"type":"kType"}] }
// #dependsOn {"typeCtor":"Recover"}
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"Validate_aType_idx_kType"}], "funcTArgs":[{"type":"kType"}]}
func (vs Validate_sType_idx_kType) GuardBound__aType(key kType, guard Guard, f FuncFrom_sType_to_Validate_aType_idx_kType) Validate_aType_idx_kType {
	return func() (a aType, errs Errors_kType) {
		var s sType
		s, errs = vs()
		if errs.IsEmpty() {
			var va Validate_aType_idx_kType
			va, errs = f.Guarded__kType(key, guard)(s)
			if errs.IsEmpty() {
				a, errs = va()
			}
		}
		return
	}
}
