package stream

type B_String = string

func (as Int_Stream) Mapped_String(f func(e A_Int) B_String) String_Stream {
	if as == nil {
		return nil
	} else {
		h, t := as()
		return t.Mapped_String(f).PrecededBy(f(h))
	}
}

func (as Int_Stream) Binded_String(f func(A_Int) String_Stream) String_Stream {
	if as == nil {
		return nil
	} else {
		h, t := as()
		return f(h).FollowedBy(t.Binded_String(f))
	}
}
