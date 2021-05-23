package fung

type FuncFrom_xType_to_yType func(x xType) yType

func (fxy FuncFrom_xType_to_yType) Fix(x xType) func() yType {
	return func() yType {
		return fxy(x)
	}
}
