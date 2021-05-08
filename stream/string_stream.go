package stream

// The type of the elements in the stream
type e_string = string

// The type of the stream itself
type String_Stream func() (e_string, String_Stream)

func String_Empty() String_Stream {
	return nil
}
func String_Unit(a e_string) String_Stream {
	return func() (e_string, String_Stream) {
		return a, nil
	}
}
func String_FromSlice(slice []e_string) String_Stream {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_string, String_Stream) {
			return slice[1], String_FromSlice(slice[1:])
		}
	}
}
func String_FromSet(m map[e_string]bool) String_Stream {
	slice := make([]e_string, len(m))
	for k := range m {
		slice = append(slice, k)
	}
	return String_FromSlice(slice)
}

////

func (es String_Stream) IsEmpty() bool {
	return es == nil
}

func (es String_Stream) Filtered(p func(e_string) bool) String_Stream {
	if es == nil {
		return nil
	} else {
		h, t := es()
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

func (es String_Stream) PrecededBy(a e_string) String_Stream {
	return func() (e_string, String_Stream) {
		return a, es
	}
}
func (es String_Stream) SuccedeedBy(a e_string) String_Stream {
	return es.FollowedBy(String_Unit(a))
}
func (as1 String_Stream) FollowedBy(as2 String_Stream) String_Stream {
	if as1 != nil {
		h, t := as1()
		return func() (e_string, String_Stream) { return h, t.FollowedBy(as2) }
	} else {
		return as2
	}
}
func (es1 String_Stream) IsEqualTo(es2 String_Stream) bool {
	if es1 == nil {
		return es2 == nil
	} else if es2 == nil {
		return false
	} else {
		h1, t1 := es1()
		h2, t2 := es2()
		return h1 == h2 && t1.IsEqualTo(t2)
	}
}

func (es String_Stream) AppendToSlice(s []e_string) []e_string {
	if es != nil {
		h, t := es()
		// All the following lines could be replaced by this << return t.AppendToSlice(append(s, h)) >> if the golang compiler supported tail recursion optimization.
		s = append(s, h)
		for t != nil {
			h, t = t()
			s = append(s, h)
		}
	}
	return s
}

func (es String_Stream) ToSlice(initialCapacity int) []e_string {
	slice := make([]e_string, 0, initialCapacity)
	return es.AppendToSlice(slice)
}
