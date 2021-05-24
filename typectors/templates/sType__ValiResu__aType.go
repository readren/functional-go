package fung

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"}] }
func (vrs ValiResu_sType) Mapped__aType(f func(sType) aType) (vra ValiResu_aType) {
	if vrs.Err == nil {
		vra.Val = f(vrs.Val)
	} else {
		vra.Err = vrs.Err
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"}] }
func (vrs ValiResu_sType) Bound__aType(f func(sType) ValiResu_aType) (vra ValiResu_aType) {
	if vrs.Err == nil {
		vra = f(vrs.Val)
	} else {
		vra.Err = vrs.Err
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"aType"}]}
func (vrs ValiResu_sType) GuardMapped__aType(guard Guard, f FuncFrom_sType_to_aType) (vra ValiResu_aType) {
	if vrs.Err == nil {
		vra.Val, vra.Err = f.Try(vrs.Val, guard)
	} else {
		vra.Err = vrs.Err
	}
	return
}

// #dependsOn {"typeCtor":"ValiResu", "baseTArgs": [{"type":"aType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"ValiResu_aType"}]}
func (vrs ValiResu_sType) GuardBound__aType(guard Guard, f FuncFrom_sType_to_ValiResu_aType) (vra ValiResu_aType) {
	if vrs.Err == nil {
		var err error
		vra, err = f.Try(vrs.Val, guard)
		if err != nil {
			vra.Err = err
		}
	} else {
		vra.Err = vrs.Err
	}
	return
}
