package stream

type B_Any = interface{}

func (as Any_Stream) Mapped_Any(f func(A_Any) B_Any) Any_Stream {
	if as == nil {
		return nil
	} else {
		h, t := as()
		return t.Mapped_Any(f).PrecededBy(f(h))
	}
}

func (as Any_Stream) Binded_Any(f func(A_Any) Any_Stream) Any_Stream {
	if as == nil {
		return nil
	} else {
		h, t := as()
		return f(h).FollowedBy(t.Binded_Any(f))
	}
}
