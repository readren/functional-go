package fung

// Giver is like a function that takes no parameter.
type Giver_sType func() sType

func (ss Giver_sType) Apply() sType {
	return ss()
}

func (ss Giver_sType) Try(guard func(any) bool) (s sType, err any) {
	defer catch(&err, guard)
	s = ss()
	return
}
