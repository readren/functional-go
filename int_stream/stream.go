package int_stream

type A = int
type Stream func() (A, Stream)

func Unit(a A) Stream {
	return func() (A, Stream) {
		return a, nil
	}
}
func Empty() Stream {
	return nil
}
func (as Stream) IsEmpty() bool {
	return as == nil
}

func (as Stream) Filtered(p func(A) bool) Stream {
	if as == nil {
		return nil
	} else {
		h, t := as()
		pass := p(h)
		for !pass && t != nil {
			h, t = t()
			pass = p(h)
		}
		if pass {
			return func() (A, Stream) {
				return h, t.Filtered(p)
			}
		} else {
			return nil
		}
	}
}

func (as Stream) Prepend(a A) Stream {
	return func() (A, Stream) {
		return a, as
	}
}
func (as1 Stream) Concatenated(as2 Stream) Stream {
	if as1 != nil {
		h, t := as1()
		return func() (A, Stream) { return h, t.Concatenated(as2) }
	} else {
		return as2
	}
}
func (as1 Stream) EqualTo(as2 Stream) bool {
	if as1 == nil {
		return as2 == nil
	} else if as2 == nil {
		return false
	} else {
		h1, t1 := as1()
		h2, t2 := as2()
		return h1 == h2 && t1.EqualTo(t2)
	}
}

func (as Stream) AppendToSlice(s []A) []A {
	if as != nil {
		h, t := as()
		// All the following lines could be replaced by this << return t.AppendToSlice(append(s, h)) >> if the golang compiler supported tail recursion optimization.
		s = append(s, h)
		for t != nil {
			h, t = t()
			s = append(s, h)
		}
	}
	return s
}

func (as Stream) ToSlice(initialCapacity int) []A {
	slice := make([]A, 0, initialCapacity)
	return as.AppendToSlice(slice)
}
