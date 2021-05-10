package stream

// The first type parameter of the methods
type a_any = interface{}

func (es Any) Mapped_Any(f func(e_any) a_any) Any {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return t.Mapped_Any(f).PrecededBy(f(h))
	}
}

func (es Any) Binded_Any(f func(e_any) Any) Any {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return f(h).FollowedBy(t.Binded_Any(f))
	}
}
