package templates

// DO NOT COPY contained lines - BEGIN. They exist to make the compiler happy.

type someTypeA struct{}

// The type of the stream whose elements are of type `a_type`
type AStream func() (a_type, AStream)

func AStream_Single(a a_type) AStream         { panic("This template line should have been removed") }
func (as AStream) FollowedBy(AStream) AStream { panic("This template line should have been removed") }

// DO NOT COPY contained lines - END

// The first type parameter of the methods
type a_type = someTypeA

// Note: this method is partially lazy. Applying it traverses the first elements of this stream until the first included element inclusive.
func (es EStream) Collected_a(f func(e_type) (bool, a_type)) AStream {
	var e e_type
	for es != nil {
		e, es = es()
		isIncluded, a := f(e)
		if isIncluded {
			return func() (a_type, AStream) {
				return a, es.Collected_a(f)
			}
		}
	}
	return nil
}

// KEI stands for knowing elements indexes
// Note: this method is partially lazy. Applying it traverses the first elements of this stream until the first included element inclusive.
func (es EStream) CollectedKEI_a(baseIndex int, f func(e e_type, index int) (bool, a_type)) AStream {
	var e e_type
	for es != nil {
		e, es = es()
		c, a := f(e, baseIndex)
		if c {
			return func() (a_type, AStream) {
				return a, es.CollectedKEI_a(baseIndex+1, f)
			}
		}
	}
	return nil
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es EStream) Mapped_a(f func(e_type) a_type) AStream {
	if es == nil {
		return nil
	} else {
		return func() (a_type, AStream) {
			h, t := es()
			return f(h), t.Mapped_a(f)
		}
	}
}

// KEI stands for knowing elements indexes
// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es EStream) MappedKEI_a(indexBase int, f func(e e_type, index int) a_type) AStream {
	if es == nil {
		return nil
	} else {
		return func() (a_type, AStream) {
			h, t := es()
			return f(h, indexBase), t.MappedKEI_a(indexBase+1, f)
		}
	}
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es EStream) Scanned_a(z a_type, f func(a_type, e_type) a_type) AStream {
	if es == nil {
		return AStream_Single(z)
	} else {
		return func() (a_type, AStream) {
			var e e_type
			e, es = es()
			a := f(z, e)
			return a, es.Scanned_a(a, f)
		}
	}
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es EStream) Bound_a(f func(e_type) AStream) AStream {
	if es == nil {
		return nil
	} else {
		return func() (a_type, AStream) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Bound_a(f))
		}
	}
}

// KEI stands for knowing elements indexes
// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es EStream) BoundKEI_a(indexBase int, f func(e e_type, index int) AStream) AStream {
	if es == nil {
		return nil
	} else {
		return func() (a_type, AStream) {
			he, te := es()
			ha, ta := f(he, indexBase)()
			return ha, ta.FollowedBy(te.BoundKEI_a(indexBase+1, f))
		}
	}
}

func (es EStream) FoldLeft_a(z a_type, f func(a_type, e_type) a_type) a_type {
	var h e_type
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

// CAUTION: this method is recursive and consumes stack space proportional to both, the stream size and the size of its elements. Use `FoldLeft` instead if possible.
func (es EStream) FoldRight_a(f func(e_type, a_type) a_type, z a_type) a_type {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_a(f, z))
	}
}

func (es EStream) Corresponds_a(as AStream, f func(e e_type, a a_type) bool) bool {
	var e e_type
	var a a_type
	for es != nil && as != nil {
		e, es = es()
		a, as = as()
		if !f(e, a) {
			return false
		}
	}
	return es == nil && as == nil
}
