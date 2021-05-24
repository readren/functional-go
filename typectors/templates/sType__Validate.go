package fung

type Validate_sType func() (sType, error)

// Creates a `Validate_sType` instance from a guarded `Giver_sType`.
// #dependsOn {"typeCtor":"Giver1", "baseTArgs": [{"type":"sType"}]}
func Validate_sType__Try(g Giver_sType, guard Guard) Validate_sType {
	return func() (sType, error) {
		return g.Try(guard)
	}
}
