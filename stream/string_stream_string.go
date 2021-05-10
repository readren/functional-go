package stream

// The first type parameter of the methods
// type a_string = string

func (es String) Mapped_String(f func(e e_string) a_string) String {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return t.Mapped_String(f).PrecededBy(f(h))
	}
}

func (es String) Binded_String(f func(e_string) String) String {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return f(h).FollowedBy(t)
	}
}
