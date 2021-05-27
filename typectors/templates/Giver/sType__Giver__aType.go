package fung

// #dependsOn {"typeCtor": "Giver1", "baseTypeArgs": [{"type": "sType"}] }

// #importAnchor

// #dependsOn {"typeCtor": "Errors", "baseTypeArgs": [{"type": "aType"}] }
func (gs Giver_sType) Guarded__aType(key aType, guard Guard) Validate_sType_idx_aType {
	return func() (s sType, errs Errors_aType) {
		defer recover__catch__aType(&errs, guard, key)
		s = gs()
		return
	}
}
