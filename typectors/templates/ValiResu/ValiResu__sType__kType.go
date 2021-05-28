package fung

// #importAnchor

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
func (companion valiResuCompanion) Successful__sType__kType(a sType) (vr ValiResu_sType_by_kType) {
	vr.Val = a
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"sType", "type":"kType"}] }
func (companion valiResuCompanion) Failed__sType__kType(key kType, err error) (vr ValiResu_sType_by_kType) {
	vr.Errs = map[kType]error{key: err}
	return
}
