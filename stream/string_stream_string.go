package stream

// The first type parameter of the methods
// type a_string = string

func (es String) Mapped_any(f func(e_string) a_string) String {
	if es == nil {
		return nil
	} else {
		return func() (e_string, String) {
			h, t := es()
			return f(h), t.Mapped_any(f)
		}
	}
}

func (es String) Binded_any(f func(e_string) String) String {
	if es == nil {
		return nil
	} else {
		return func() (e_string, String) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Binded_any(f))
		}
	}
}

func (es String) FoldLeft_any(z a_string, f func(a_string, e_string) a_string) a_string {
	var h e_string
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es String) FoldRight_any(f func(e_string, a_string) a_string, z a_string) a_string {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_any(f, z))
	}
}
