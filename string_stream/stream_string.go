package string_stream

type B_String = string

func (as Stream) Mapped_String(f func(e A) B_String) Stream {
	if as == nil {
		return nil
	} else {
		return func() (B_String, Stream) {
			h, t := as()
			return f(h), t.Mapped_String(f)
		}
	}
}
