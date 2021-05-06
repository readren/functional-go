package int_stream

type B = int

func (as Stream) Mapped_int(f func(e A) int) Stream {
	if as == nil {
		return nil
	} else {
		return func() (int, Stream) {
			h, t := as()
			return f(h), t.Mapped_int(f)
		}
	}
}
