package stream

// The first type parameter of the methods
type a_string = string

func (es Int) Mapped_String(f func(e e_int) a_string) String {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return t.Mapped_String(f).PrecededBy(f(h))
	}
}

func (es Int) Binded_String(f func(e_int) String) String {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return f(h).FollowedBy(t.Binded_String(f))
	}
}
