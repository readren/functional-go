package fung

type ValiResu_sType struct {
	Val sType
	Err interface{}
}

func ValiResu_sType__Try(f func() sType) (vrs ValiResu_sType) {
	defer catch(&vrs.Err)
	vrs.Val = f()
	return
}
