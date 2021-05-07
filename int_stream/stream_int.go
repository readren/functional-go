package int_stream

type B_Int = int

func (as Stream) Mapped_Int(f func(e A) B_Int) Stream {
	if as == nil {
		return nil
	} else {
		return func() (B_Int, Stream) {
			h, t := as()
			return f(h), t.Mapped_Int(f)
		}
	}
}
