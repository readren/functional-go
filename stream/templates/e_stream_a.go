package templates

// DO NOT COPY contained lines - BEGIN. They exist to make the compiler happy.

type someTypeA struct{}

// The type of the stream whose elements are of type `a_type`
type AStream func() (a_type, AStream)

func (as AStream) FollowedBy(AStream) AStream { panic("This template line should have been removed") }

// DO NOT COPY contained lines - END

// The first type parameter of the methods
type a_type = someTypeA

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

func (es EStream) Binded_a(f func(e_type) AStream) AStream {
	if es == nil {
		return nil
	} else {
		return func() (a_type, AStream) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Binded_a(f))
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

func (es EStream) FoldRight_a(f func(e_type, a_type) a_type, z a_type) a_type {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_a(f, z))
	}
}
