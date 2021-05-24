package fung

// Giver is like a function that takes no parameter.
type Giver_sType func() sType

func (gs Giver_sType) Apply() sType {
	return gs()
}

func (gs Giver_sType) Try(guard Guard) (s sType, err error) {
	defer catch(&err, guard)
	s = gs()
	return
}
