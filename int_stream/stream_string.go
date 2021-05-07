package int_stream

import "github.com/readren/go-stream/generic_stream"

type B_String = string

func (as Stream) Mapped_String(f func(e A) B_String) generic_stream.Stream {
	if as == nil {
		return nil
	} else {
		return func() (B_String, generic_stream.Stream) {
			h, t := as()
			return f(h), t.Mapped_String(f)
		}
	}
}
