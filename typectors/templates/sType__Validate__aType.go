package fung

// #dependsOn {"typeCtor":"Try", "baseTArgs": [{"type":"aType"}] }
// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"}] }
func (vs Validate_sType) Mapped__aType(f func(sType) aType) Validate_aType {
	return func() (aType, interface{}) {
		s, err := vs()
		if err == nil {
			return Try_aType(func() aType {
				return f(s)
			}).Catch()
		} else {
			var a aType
			return a, err
		}
	}
}
