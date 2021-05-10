package stream

// The type of the elements in the stream
type e_int = int

func int_equality(e1, e2 e_int) bool {
	return e1 == e2
}

// The type of the stream whose elements are of type `e_int`
type Int func() (e_int, Int)

func Int_Empty() Int {
	return nil
}

func Int_Single(e e_int) Int {
	return func() (e_int, Int) {
		return e, nil
	}
}

func Int_Forever(e e_int) Int {
	return func() (e_int, Int) {
		return e, Int_Forever(e)
	}
}

func Int_FromSlice(slice []e_int) Int {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_int, Int) {
			return slice[0], Int_FromSlice(slice[1:])
		}
	}
}
func Int_FromSet(m map[e_int]bool) Int {
	slice := make([]e_int, len(m))
	for k := range m {
		slice = append(slice, k)
	}
	return Int_FromSlice(slice)
}

////

func (es Int) IsEmpty() bool {
	return es == nil
}

func (es Int) TakeWhile(indexBase int, p func(elem e_int, index int) bool) Int {
	if es == nil {
		return nil
	} else {
		h, t := es()
		if p(h, indexBase) {
			return func() (e_int, Int) {
				return h, t.TakeWhile(indexBase+1, p)
			}
		} else {
			return nil
		}
	}
}

func (es Int) DropWhile(indexBase int, p func(elem e_int, index int) bool) Int {
	for es != nil {
		h, t := es()
		if !p(h, indexBase) {
			return es
		}
		indexBase += 1
		es = t
	}
	return nil
}

func (es Int) Filtered(p func(elem e_int) bool) Int {
	var h e_int
	for es != nil {
		h, es = es()
		if p(h) {
			return func() (e_int, Int) {
				return h, es.Filtered(p)
			}
		}
	}
	return nil
}

func (es Int) PrecededBy(a e_int) Int {
	return func() (e_int, Int) {
		return a, es
	}
}

func (es Int) SuccedeedBy(a e_int) Int {
	return es.FollowedBy(Int_Single(a))
}

func (es1 Int) FollowedBy(es2 Int) Int {
	if es1 != nil {
		return func() (e_int, Int) {
			h, t := es1()
			return h, t.FollowedBy(es2)
		}
	} else {
		return es2
	}
}

func (es Int) ForAll(p func(e_int) bool) bool {
	z := true
	var h e_int
	for es != nil && z {
		h, es = es()
		z = p(h)
	}
	return z
}

func (es Int) ForAny(p func(e_int) bool) bool {
	return es.ForAll(func(e e_int) bool {
		return !p(e)
	})
}

func (es1 Int) IsEqualTo(es2 Int) bool {
	var h1, h2 e_int
	for es1 != nil && es2 != nil {
		h1, es1 = es1()
		h2, es2 = es2()
		if !int_equality(h1, h2) {
			return false
		}
	}
	return es1 == nil && es2 == nil
}

func (es Int) AppendToSlice(s []e_int) []e_int {
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

func (es Int) ToSlice(initialCapacity int) []e_int {
	slice := make([]e_int, 0, initialCapacity)
	return es.AppendToSlice(slice)
}
