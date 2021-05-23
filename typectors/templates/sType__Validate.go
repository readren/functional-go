package fung

type Validate_sType func() (sType, interface{})

// #dependsOn {"typeCtor":"Try", "baseTArgs": [{"type":"sType"}]}
func Validate_sType__Try(f Try_sType) Validate_sType {
	return func() (sType, interface{}) {
		return f.Catch()
	}
}
