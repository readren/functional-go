package stream

// The first type parameter of the methods
type a_string = string

func (es Int) Mapped_string(f func(e_int) a_string) String {
	if es == nil {
		return nil
	} else {
		return func() (a_string, String) {
			h, t := es()
			return f(h), t.Mapped_string(f)
		}
	}
}

func (es Int) Binded_string(f func(e_int) String) String {
	if es == nil {
		return nil
	} else {
		return func() (a_string, String) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Binded_string(f))
		}
	}
}

func (es Int) FoldLeft_string(z a_string, f func(a_string, e_int) a_string) a_string {
	var h e_int
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es Int) FoldRight_string(f func(e_int, a_string) a_string, z a_string) a_string {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_string(f, z))
	}
}
