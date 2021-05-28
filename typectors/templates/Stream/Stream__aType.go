package fung

// #importAnchor

func (companion streamCompanion) Single__aType(a aType) Stream_aType {
	return func() (aType, Stream_aType) {
		return a, nil
	}
}

// #startOfFuncsWithNoInternalDependants

func (companion streamCompanion) Empty__aType() Stream_aType {
	return nil
}

func (companion streamCompanion) Forever__aType(a aType) Stream_aType {
	return func() (aType, Stream_aType) {
		return a, companion.Forever__aType(a)
	}
}

func (companion streamCompanion) FromSlice__aType(slice []aType) Stream_aType {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (aType, Stream_aType) {
			return slice[0], companion.FromSlice__aType(slice[1:])
		}
	}
}
