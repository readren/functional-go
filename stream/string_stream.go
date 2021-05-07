package stream

type A_String = string
type String_Stream func() (A_String, String_Stream)

func String_Empty() String_Stream {
	return nil
}
func String_Unit(a A_String) String_Stream {
	return func() (A_String, String_Stream) {
		return a, nil
	}
}
func (as String_Stream) IsEmpty() bool {
	return as == nil
}

func (as String_Stream) Filtered(p func(A_String) bool) String_Stream {
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

func (as String_Stream) PrecededBy(a A_String) String_Stream {
	return func() (A_String, String_Stream) {
		return a, as
	}
}
func (as String_Stream) SuccedeedBy(a A_String) String_Stream {
	return as.FollowedBy(String_Unit(a))
}
func (as1 String_Stream) FollowedBy(as2 String_Stream) String_Stream {
	if as1 != nil {
		h, t := as1()
		return func() (A_String, String_Stream) { return h, t.FollowedBy(as2) }
	} else {
		return as2
	}
}
func (as1 String_Stream) IsEqualTo(as2 String_Stream) bool {
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

func (as String_Stream) AppendToSlice(s []A_String) []A_String {
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

func (as String_Stream) ToSlice(initialCapacity int) []A_String {
	slice := make([]A_String, 0, initialCapacity)
	return as.AppendToSlice(slice)
}
