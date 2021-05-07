package stream

// type B_Int = int

func (as String_Stream) Mapped_Int(f func(e A_String) B_Int) Int_Stream {
	if as == nil {
		return nil
	} else {
		return func() (B_Int, Int_Stream) {
			h, t := as()
			return f(h), t.Mapped_Int(f)
		}
	}
}