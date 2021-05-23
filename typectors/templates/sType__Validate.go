package fung

type Validate_sType func() (sType, interface{})

// #dependsOn {"typeCtor":"Giver1", "baseTArgs": [{"type":"sType"}]}
func Validate_sType__Try(g Giver_sType) Validate_sType {
	return func() (sType, interface{}) {
		return g.Try()
	}
}
