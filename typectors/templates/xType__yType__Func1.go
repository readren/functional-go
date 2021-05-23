package fung

type FuncFrom_xType_to_yType func(x xType) yType

func (fxy FuncFrom_xType_to_yType) Apply(x xType) yType {
	return fxy(x)
}

func (fxy FuncFrom_xType_to_yType) Try(x xType) (y yType, err interface{}) {
	defer catch(&err)
	y = fxy(x)
	return
}
