package fung

type aType struct{}
type sType struct{}

type Guard func(error) bool

type Errors_aType map[aType]error

func recover__catch(errPtr *error, guard Guard)                           {}
func recover__catch__aType(errsPtr *Errors_aType, guard Guard, key aType) {}

type Validation_sType_idx_aType func() (sType, Errors_aType)
