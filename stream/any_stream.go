package stream

// The type of the elements in the stream
type e_any = interface{}

// The type of the stream itself
type Any func() (e_any, Any)

func Any_Empty() Any {
	return nil
}
func Any_Unit(a e_any) Any {
	return func() (e_any, Any) {
		return a, nil
	}
}
func Any_FromSlice(slice []e_any) Any {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_any, Any) {
			return slice[1], Any_FromSlice(slice[1:])
		}
	}
}
func Any_FromSet(m map[e_any]bool) Any {
	slice := make([]e_any, len(m))
	for k := range m {
		slice = append(slice, k)
	}
	return Any_FromSlice(slice)
}

////

func (es Any) IsEmpty() bool {
	return es == nil
}

func (es Any) Filtered(p func(e_any) bool) Any {
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

func (es Any) PrecededBy(a e_any) Any {
	return func() (e_any, Any) {
		return a, es
	}
}
func (es Any) SuccedeedBy(a e_any) Any {
	return es.FollowedBy(Any_Unit(a))
}
func (es1 Any) FollowedBy(es2 Any) Any {
	if es1 != nil {
		h, t := es1()
		return func() (e_any, Any) { return h, t.FollowedBy(es2) }
	} else {
		return es2
	}
}
func (es1 Any) IsEqualTo(es2 Any) bool {
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

func (es Any) AppendToSlice(s []e_any) []e_any {
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

func (es Any) ToSlice(initialCapacity int) []e_any {
	slice := make([]e_any, 0, initialCapacity)
	return es.AppendToSlice(slice)
}
