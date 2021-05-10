package stream

// The first type parameter of the methods
type a_int = int

func (es Int) Mapped_Int(f func(e_int) a_int) Int {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return t.Mapped_Int(f).PrecededBy(f(h))
	}
}

func (es Int) Binded_Int(f func(e_int) Int) Int {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return f(h).FollowedBy(t.Binded_Int(f))
	}
}
