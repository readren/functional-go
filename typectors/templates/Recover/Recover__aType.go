package fung

// #dependsOn {"typeCtor":"Errors", "funcTArgs":[{"type":"aType"}]}
func recover__catch__aType(errsPtr *Errors_aType, guard Guard, key aType) {
	if any := recover(); any != nil {
		if err, ok := any.(error); ok && guard(err) {
			if *errsPtr == nil {
				*errsPtr = Errors.New__aType()
			}
			(*errsPtr)[key] = err
		} else {
			panic(any)
		}
	}
}
