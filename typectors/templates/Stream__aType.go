package fung

// #importAnchor

// The type of the stream whose elements are of type `aType`
type Stream_aType func() (aType, Stream_aType)

func Stream__Empty__aType() Stream_aType {
	return nil
}

func Stream__Single__aType(a aType) Stream_aType {
	return func() (aType, Stream_aType) {
		return a, nil
	}
}

func Stream__Forever__aType(a aType) Stream_aType {
	return func() (aType, Stream_aType) {
		return a, Stream__Forever__aType(a)
	}
}

func Stream__FromSlice__aType(slice []aType) Stream_aType {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (aType, Stream_aType) {
			return slice[0], Stream__FromSlice__aType(slice[1:])
		}
	}
}
