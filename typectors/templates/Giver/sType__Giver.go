package fung

// #importAnchor

// Giver is like a function that takes no parameter.
type Giver_sType func() sType

// #startOfFuncsWithNoInternalDependants

func (gs Giver_sType) Apply() sType {
	return gs()
}

// #dependsOn {"typeCtor":"Recover"}
func (gs Giver_sType) Guarded(guard Guard) func() (sType, error) {
	return func() (s sType, err error) {
		defer recover__catch(&err, guard)
		s = gs()
		return
	}
}
