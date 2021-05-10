package stream

// The first type parameter of the methods
type a_any = interface{}

func (es Any) Mapped_any(f func(e_any) a_any) Any {
	if es == nil {
		return nil
	} else {
		return func() (e_any, Any) {
			h, t := es()
			return f(h), t.Mapped_any(f)
		}
	}
}

func (es Any) Binded_any(f func(e_any) Any) Any {
	if es == nil {
		return nil
	} else {
		return func() (e_any, Any) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Binded_any(f))
		}
	}
}

func (es Any) FoldLeft_any(z a_any, f func(a_any, e_any) a_any) a_any {
	var h e_any
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es Any) FoldRight_any(f func(e_any, a_any) a_any, z a_any) a_any {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_any(f, z))
	}
}
