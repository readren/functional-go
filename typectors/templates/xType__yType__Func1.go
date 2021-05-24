package fung

type FuncFrom_xType_to_yType func(x xType) yType

func (fxy FuncFrom_xType_to_yType) Apply(x xType) yType {
	return fxy(x)
}

func (fxy FuncFrom_xType_to_yType) Try(x xType, guard Guard) (y yType, err any) {
	defer catch(&err, guard)
	y = fxy(x)
	return
}

func (fxy FuncFrom_xType_to_yType) Fix(x xType) func() yType {
	return func() yType {
		return fxy(x)
	}
}
