package functional

// #exclude-section-begin These lines are not included in the generated source files. They exist to make the template file compiler friendly.

// The first type parameter of the methods
type aType = struct{}

// The type of the stream whose elements are of type `aType`
// #requires {"typeCtor":"stream", "baseTArgs": [{"type":"aType"}]}
type Stream_aType func() (aType, Stream_aType)

func Stream_aType__Single(a aType) Stream_aType {
	panic("This template line should have been removed")
}
func (as Stream_aType) FollowedBy(Stream_aType) Stream_aType {
	panic("This template line should have been removed")
}

// #exclude-section-end

// Note: this method is partially lazy. Applying it traverses the first elements of this stream until the first included element inclusive.
func (es Stream_eType) Collected_aType(f func(eType) (bool, aType)) Stream_aType {
	var e eType
	for es != nil {
		e, es = es()
		isIncluded, a := f(e)
		if isIncluded {
			return func() (aType, Stream_aType) {
				return a, es.Collected_aType(f)
			}
		}
	}
	return nil
}

// KEI stands for knowing elements indexes
// Note: this method is partially lazy. Applying it traverses the first elements of this stream until the first included element inclusive.
// #requires {"typeCtor":"stream", "baseTArgs": [{"type":"aType"}]}
func (es Stream_eType) CollectedKEI_aType(baseIndex int, f func(e eType, index int) (bool, aType)) Stream_aType {
	var e eType
	for es != nil {
		e, es = es()
		c, a := f(e, baseIndex)
		if c {
			return func() (aType, Stream_aType) {
				return a, es.CollectedKEI_aType(baseIndex+1, f)
			}
		}
	}
	return nil
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es Stream_eType) Mapped_aType(f func(eType) aType) Stream_aType {
	if es == nil {
		return nil
	} else {
		return func() (aType, Stream_aType) {
			h, t := es()
			return f(h), t.Mapped_aType(f)
		}
	}
}

// KEI stands for knowing elements indexes
// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es Stream_eType) MappedKEI_aType(indexBase int, f func(e eType, index int) aType) Stream_aType {
	if es == nil {
		return nil
	} else {
		return func() (aType, Stream_aType) {
			h, t := es()
			return f(h, indexBase), t.MappedKEI_aType(indexBase+1, f)
		}
	}
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es Stream_eType) Scanned_aType(z aType, f func(aType, eType) aType) Stream_aType {
	if es == nil {
		return Stream_aType__Single(z)
	} else {
		return func() (aType, Stream_aType) {
			var e eType
			e, es = es()
			a := f(z, e)
			return a, es.Scanned_aType(a, f)
		}
	}
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es Stream_eType) Bound_aType(f func(eType) Stream_aType) Stream_aType {
	if es == nil {
		return nil
	} else {
		return func() (aType, Stream_aType) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Bound_aType(f))
		}
	}
}

// KEI stands for knowing elements indexes
// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es Stream_eType) BoundKEI_aType(indexBase int, f func(e eType, index int) Stream_aType) Stream_aType {
	if es == nil {
		return nil
	} else {
		return func() (aType, Stream_aType) {
			he, te := es()
			ha, ta := f(he, indexBase)()
			return ha, ta.FollowedBy(te.BoundKEI_aType(indexBase+1, f))
		}
	}
}

func (es Stream_eType) FoldLeft_aType(z aType, f func(aType, eType) aType) aType {
	var h eType
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

// CAUTION: this method is recursive and consumes stack space proportional to both, the stream size and the size of its elements. Use `FoldLeft` instead if possible.
func (es Stream_eType) FoldRight_aType(f func(eType, aType) aType, z aType) aType {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_aType(f, z))
	}
}

func (es Stream_eType) Corresponds_aType(as Stream_aType, f func(e eType, a aType) bool) bool {
	var e eType
	var a aType
	for es != nil && as != nil {
		e, es = es()
		a, as = as()
		if !f(e, a) {
			return false
		}
	}
	return es == nil && as == nil
}
