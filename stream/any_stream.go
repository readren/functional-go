package stream

// The type of the elements in the stream
type e_any = interface{}

// The type of the stream itself
type Any_Stream func() (e_any, Any_Stream)

func Any_Empty() Any_Stream {
	return nil
}
func Any_Unit(a e_any) Any_Stream {
	return func() (e_any, Any_Stream) {
		return a, nil
	}
}
func Any_FromSlice(slice []e_any) Any_Stream {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_any, Any_Stream) {
			return slice[1], Any_FromSlice(slice[1:])
		}
	}
}
func Any_FromSet(m map[e_any]bool) Any_Stream {
	slice := make([]e_any, len(m))
	for k := range m {
		slice = append(slice, k)
	}
	return Any_FromSlice(slice)
}

////

func (es Any_Stream) IsEmpty() bool {
	return es == nil
}

func (es Any_Stream) Filtered(p func(e_any) bool) Any_Stream {
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

func (es Any_Stream) PrecededBy(a e_any) Any_Stream {
	return func() (e_any, Any_Stream) {
		return a, es
	}
}
func (es Any_Stream) SuccedeedBy(a e_any) Any_Stream {
	return es.FollowedBy(Any_Unit(a))
}
func (es1 Any_Stream) FollowedBy(es2 Any_Stream) Any_Stream {
	if es1 != nil {
		h, t := es1()
		return func() (e_any, Any_Stream) { return h, t.FollowedBy(es2) }
	} else {
		return es2
	}
}
func (es1 Any_Stream) IsEqualTo(es2 Any_Stream) bool {
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

func (es Any_Stream) AppendToSlice(s []e_any) []e_any {
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

func (es Any_Stream) ToSlice(initialCapacity int) []e_any {
	slice := make([]e_any, 0, initialCapacity)
	return es.AppendToSlice(slice)
}
