package stream

// The first type parameter of the methods
type a_any = interface{}

func (es Any_Stream) Mapped_Any(f func(e_any) a_any) Any_Stream {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return t.Mapped_Any(f).PrecededBy(f(h))
	}
}

func (es Any_Stream) Binded_Any(f func(e_any) Any_Stream) Any_Stream {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return f(h).FollowedBy(t.Binded_Any(f))
	}
}
