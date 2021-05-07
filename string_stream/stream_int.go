package string_stream

import "github.com/readren/go-stream/generic_stream"

type B_Int = int

func (as Stream) Mapped_Int(f func(e A) B_Int) generic_stream.Stream {
	if as == nil {
		return nil
	} else {
		return func() (B_Int, generic_stream.Stream) {
			h, t := as()
			return f(h), t.Mapped_Int(f)
		}
	}
}
