package fung

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"}] }
func (vs Validate_sType) Mapped__aType(f func(sType) aType) Validate_aType {
	return func() (a aType, err error) {
		var s sType
		s, err = vs()
		if err == nil {
			a = f(s)
		}
		return
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"}] }
func (vs Validate_sType) Bound__aType(f func(s sType) Validate_aType) Validate_aType {
	return func() (a aType, err error) {
		var s sType
		s, err = vs()
		if err == nil {
			va := f(s)
			a, err = va()
		}
		return
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"aType"}]}
func (vs Validate_sType) GuardMapped__aType(guard Guard, f FuncFrom_sType_to_aType) Validate_aType {
	return func() (a aType, err error) {
		s, err := vs()
		if err == nil {
			return f.Try(s, guard)
		} else {
			return a, err
		}
	}
}

// #dependsOn {"typeCtor":"Validate", "baseTArgs": [{"type":"aType"}] }
// #dependsOn {"typeCtor":"Func1", "baseTArgs": [{"type":"sType"},{"type":"Validate_aType"}]}
func (vs Validate_sType) GuardBound__aType(guard Guard, f FuncFrom_sType_to_Validate_aType) Validate_aType {
	return func() (a aType, err error) {
		var s sType
		s, err = vs()
		if err == nil {
			var va Validate_aType
			va, err = f.Try(s, guard)
			if err == nil {
				a, err = va()
			}
		}
		return
	}
}
