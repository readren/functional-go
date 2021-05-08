package stream

// The type of the elements in the stream
type e_int = int

// The type of the stream itself
type Int_Stream func() (e_int, Int_Stream)

func Int_Empty() Int_Stream {
	return nil
}
func Int_Unit(a e_int) Int_Stream {
	return func() (e_int, Int_Stream) {
		return a, nil
	}
}

func Int_FromSlice(slice []e_int) Int_Stream {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_int, Int_Stream) {
			return slice[1], Int_FromSlice(slice[1:])
		}
	}
}

func Int_FromSet(m map[e_int]bool) Int_Stream {
	slice := make([]e_int, len(m))
	for k := range m {
		slice = append(slice, k)
	}
	return Int_FromSlice(slice)
}

////

func (es Int_Stream) IsEmpty() bool {
	return es == nil
}

func (es Int_Stream) Filtered(p func(e_int) bool) Int_Stream {
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

func (es Int_Stream) PrecededBy(a e_int) Int_Stream {
	return func() (e_int, Int_Stream) {
		return a, es
	}
}
func (es Int_Stream) SuccedeedBy(a e_int) Int_Stream {
	return es.FollowedBy(Int_Unit(a))
}
func (es1 Int_Stream) FollowedBy(es2 Int_Stream) Int_Stream {
	if es1 != nil {
		h, t := es1()
		return func() (e_int, Int_Stream) { return h, t.FollowedBy(es2) }
	} else {
		return es2
	}
}
func (es1 Int_Stream) IsEqualTo(es2 Int_Stream) bool {
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

func (es Int_Stream) AppendToSlice(s []e_int) []e_int {
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

func (es Int_Stream) ToSlice(initialCapacity int) []e_int {
	slice := make([]e_int, 0, initialCapacity)
	return es.AppendToSlice(slice)
}
