package fung

// Giver is like a function that takes no parameter.
type Giver_sType func() sType

func (ss Giver_sType) Apply() sType {
	return ss()
}

func (ss Giver_sType) Try() (s sType, err interface{}) {
	defer catch(&err)
	s = ss()
	return
}
