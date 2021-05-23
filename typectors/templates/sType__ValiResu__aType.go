package fung

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"aType"}]}
func (vs ValiResu_sType) Mapped__aType(f FuncFrom_sType_to_aType) (va ValiResu_aType) {
	if vs.Err == nil {
		va = ValiResu_aType__Try(f.Fix(vs.Val))
	} else {
		va.Err = vs.Err
	}
	return
}
