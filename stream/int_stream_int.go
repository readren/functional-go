package stream

// The first type parameter of the methods
type a_int = int

func (es Int) Mapped_int(f func(e_int) a_int) Int {
	if es == nil {
		return nil
	} else {
		return func() (e_int, Int) {
			h, t := es()
			return f(h), t.Mapped_int(f)
		}
	}
}

func (es Int) Binded_int(f func(e_int) Int) Int {
	if es == nil {
		return nil
	} else {
		return func() (e_int, Int) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Binded_int(f))
		}
	}
}

func (es Int) FoldLeft_int(z a_int, f func(a_int, e_int) a_int) a_int {
	var h e_int
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es Int) FoldRight_int(f func(e_int, a_int) a_int, z a_int) a_int {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_int(f, z))
	}
}
