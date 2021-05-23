package fung

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"}] }
func (vs ValiResu_sType) Mapped__aType(f func(sType) aType) ValiResu_aType {
	if vs.Err == nil {
		return ValiResu_aType__Try(func() aType {
			return f(vs.Val)
		})
	} else {
		var va ValiResu_aType
		va.Err = vs.Err
		return va
	}
}
