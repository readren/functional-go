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
			return slice[0], Any_FromSlice(slice[1:])
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
	var h e_any
	for es != nil {
		h, es = es()
		if p(h) {
			return func() (e_any, Any) {
				return h, es.Filtered(p)
			}
		}
	}
	return nil
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
		return func() (e_any, Any) {
			h, t := es1()
			return h, t.FollowedBy(es2)
		}
	} else {
		return es2
	}
}

func (es Any) ForAll(p func(e_any) bool) bool {
	z := true
	var h e_any
	for es != nil && z {
		h, es = es()
		z = p(h)
	}
	return z
}

func (es Any) ForAny(p func(e_any) bool) bool {
	return es.ForAll(func(e e_any) bool {
		return !p(e)
	})
}

func (es1 Any) IsEqualTo(es2 Any) bool {
	var h1, h2 e_any
	for es1 != nil && es2 != nil {
		h1, es1 = es1()
		h2, es2 = es2()
		if h1 != h2 {
			return false
		}
	}
	return es1 == nil && es2 == nil
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
