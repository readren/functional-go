package fung

// #importAnchor

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
func Validation__Successful__sType__kType(s sType) Validation_sType_idx_kType {
	return func() (sType, Errors_kType) {
		return s, nil
	}
}

// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"sType", "type":"kType"}] }
func Validation__Failed__sType__kType(key kType, err error) Validation_sType_idx_kType {
	return func() (s sType, errs Errors_kType) {
		errs = map[kType]error{key: err}
		return
	}
}

// #dependsOn {"typeCtor":"Stream", "baseTArgs": [{"type":"Validation_sType_idx_kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"[]sType"},"type":"kType"] }
// #dependsOn {"typeCtor":"Validation", "funcTArgs":[{"type":"Validation_sType_idx_kType"},{"type":"kType"},{"type":"sType"}]}
// #dependsOn {"typeCtor":"Func1", "funcTArgs":[{"type":"Validation_sType_idx_kType"}]}
func Validation__Sequenced__sType__kType(stream Stream_Validation_sType_idx_kType) Validation_slice_sType_idx_kType {
	return Validation__Traverse__Validation_sType_idx_kType__kType__sType(stream, Func1__Identity__Validation_sType_idx_kType)
}
