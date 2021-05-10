package stream

// The first type parameter of the methods
// type a_int = int

func (es String) Mapped_int(f func(e_string) a_int) Int {
	if es == nil {
		return nil
	} else {
		return func() (a_int, Int) {
			h, t := es()
			return f(h), t.Mapped_int(f)
		}
	}
}

func (es String) Binded_int(f func(e_string) Int) Int {
	if es == nil {
		return nil
	} else {
		return func() (a_int, Int) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Binded_int(f))
		}
	}
}

func (es String) FoldLeft_int(z a_int, f func(a_int, e_string) a_int) a_int {
	var h e_string
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es String) FoldRight_int(f func(e_string, a_int) a_int, z a_int) a_int {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_int(f, z))
	}
}
