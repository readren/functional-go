package stream

// type B_Int = int

func (as String_Stream) Mapped_Int(f func(e A_String) B_Int) Int_Stream {
	if as == nil {
		return nil
	} else {
		h, t := as()
		return t.Mapped_Int(f).PrecededBy(f(h))
	}
}

func (as String_Stream) Binded_Int(f func(A_String) Int_Stream) Int_Stream {
	if as == nil {
		return nil
	} else {
		h, t := as()
		return f(h).FollowedBy(t.Binded_Int(f))
	}
}
