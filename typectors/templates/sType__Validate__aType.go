package fung

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"aType"}]}
func (vs Validate_sType) Mapped__aType(f FuncFrom_sType_to_aType) Validate_aType {
	return func() (a aType, err interface{}) {
		s, err := vs()
		if err == nil {
			return f.Try(s)
		} else {
			return a, err
		}
	}
}
