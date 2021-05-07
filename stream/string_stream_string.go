package stream

// type B_String = string

func (as String_Stream) Mapped_String(f func(e A_String) B_String) String_Stream {
	if as == nil {
		return nil
	} else {
		return func() (B_String, String_Stream) {
			h, t := as()
			return f(h), t.Mapped_String(f)
		}
	}
}
