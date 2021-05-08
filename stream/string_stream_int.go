package stream

// The first type parameter of the methods
// type a_int = int

func (es String_Stream) Mapped_Int(f func(e e_string) a_int) Int_Stream {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return t.Mapped_Int(f).PrecededBy(f(h))
	}
}

func (es String_Stream) Binded_Int(f func(e_string) Int_Stream) Int_Stream {
	if es == nil {
		return nil
	} else {
		h, t := es()
		return f(h).FollowedBy(t.Binded_Int(f))
	}
}
