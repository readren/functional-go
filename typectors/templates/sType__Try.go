package fung

type Try_sType func() sType

func (ts Try_sType) Catch() (s sType, err interface{}) {
	defer catch(&err)
	s = ts()
	return
}
