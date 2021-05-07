package stream

type A_Any = interface{}
type Any_Stream func() (A_Any, Any_Stream)

func Any_Empty() Any_Stream {
	return nil
}
func Any_Unit(a A_Any) Any_Stream {
	return func() (A_Any, Any_Stream) {
		return a, nil
	}
}
func (as Any_Stream) IsEmpty() bool {
	return as == nil
}

func (as Any_Stream) Filtered(p func(A_Any) bool) Any_Stream {
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
			return t.Filtered(p).PrecededBy(h)
		} else {
			return nil
		}
	}
}

func (as Any_Stream) PrecededBy(a A_Any) Any_Stream {
	return func() (A_Any, Any_Stream) {
		return a, as
	}
}
func (as Any_Stream) SuccedeedBy(a A_Any) Any_Stream {
	return as.FollowedBy(Any_Unit(a))
}
func (as1 Any_Stream) FollowedBy(as2 Any_Stream) Any_Stream {
	if as1 != nil {
		h, t := as1()
		return func() (A_Any, Any_Stream) { return h, t.FollowedBy(as2) }
	} else {
		return as2
	}
}
func (as1 Any_Stream) IsEqualTo(as2 Any_Stream) bool {
	if as1 == nil {
		return as2 == nil
	} else if as2 == nil {
		return false
	} else {
		h1, t1 := as1()
		h2, t2 := as2()
		return h1 == h2 && t1.IsEqualTo(t2)
	}
}

func (as Any_Stream) AppendToSlice(s []A_Any) []A_Any {
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

func (as Any_Stream) ToSlice(initialCapacity int) []A_Any {
	slice := make([]A_Any, 0, initialCapacity)
	return as.AppendToSlice(slice)
}
