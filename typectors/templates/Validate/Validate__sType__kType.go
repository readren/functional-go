package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
func Validate__Successful__sType__kType(s sType) Validate_sType_idx_kType {
	return func() (sType, Errors_kType) {
		return s, nil
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"sType", "type":"kType"}] }
func Validate__Failed__sType__kType(key kType, err error) Validate_sType_idx_kType {
	return func() (s sType, errs Errors_kType) {
		errs = map[kType]error{key: err}
		return
	}
}

// #dependsOn {"typeCtor":"Stream", "baseTArgs": [{"type":"Validate_sType_idx_kType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"[]sType"},"type":"kType"] }
// #dependsOn {"typeCtor":"Validate", "funcTArgs":[{"type":"Validate_sType_idx_kType"},{"type":"kType"},{"type":"sType"}]}
// #dependsOn {"typeCtor":"Func1", "funcTArgs":[{"type":"Validate_sType_idx_kType"}]}
func Validate__Sequenced__sType__kType(stream Stream_Validate_sType_idx_kType) Validate_slice_sType_idx_kType {
	return Validate__Traverse__Validate_sType_idx_kType__kType__sType(stream, Func1__Identity__Validate_sType_idx_kType)
}
