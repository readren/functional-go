package fung

// #dependsOn {"typeCtor":"Stream", "baseTArgs": [{"type":"sType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Validation", "baseTArgs": [{"type":"[]aType", "type":"kType"}] }
// #dependsOn {"typeCtor":"Errors", "baseTArgs": [{"type":"kType"}] }
func Validation__Traverse__sType__kType__aType(stream Stream_sType, f func(sType) Validation_aType_idx_kType) Validation_Slice_aType_idx_kType {
	return func() (slice []aType, errsAccum Errors_kType) {
		slice = make([]aType, 0)
		errsAccum = Errors__New__kType()
		for stream != nil {
			var e sType
			e, stream = stream()
			a, errs := f(e)()
			if errs.IsEmpty() && errsAccum.IsEmpty() {
				slice = append(slice, a)
			} else {
				errsAccum.PutAll(errs)
				slice = nil
			}
		}
		return
	}
}
