package stream

import "reflect"

// The type of the elements in the stream
type e_sliceOfInts = []int

func sliceOfInts_equality(e1, e2 e_sliceOfInts) bool {
	return reflect.DeepEqual(e1, e2)
}

// The type of the stream whose elements are of type `e_sliceOfInts`
type EStream func() (e_sliceOfInts, EStream)

func SliceOfInts_Empty() EStream {
	return nil
}

func SliceOfInts_Single(e e_sliceOfInts) EStream {
	return func() (e_sliceOfInts, EStream) {
		return e, nil
	}
}

func SliceOfInts_Forever(e e_sliceOfInts) EStream {
	return func() (e_sliceOfInts, EStream) {
		return e, SliceOfInts_Forever(e)
	}
}

func SliceOfInts_FromSlice(slice []e_sliceOfInts) EStream {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_sliceOfInts, EStream) {
			return slice[0], SliceOfInts_FromSlice(slice[1:])
		}
	}
}

// func SliceOfInts_FromSet(m map[e_sliceOfInts]bool) EStream {
// 	slice := make([]e_sliceOfInts, len(m))
// 	for k := range m {
// 		slice = append(slice, k)
// 	}
// 	return SliceOfInts_FromSlice(slice)
// }

////

func (es EStream) IsEmpty() bool {
	return es == nil
}

func (es EStream) TakeWhile(indexBase int, p func(elem e_sliceOfInts, index int) bool) EStream {
	if es == nil {
		return nil
	} else {
		h, t := es()
		if p(h, indexBase) {
			return func() (e_sliceOfInts, EStream) {
				return h, t.TakeWhile(indexBase+1, p)
			}
		} else {
			return nil
		}
	}
}

func (es EStream) DropWhile(indexBase int, p func(elem e_sliceOfInts, index int) bool) EStream {
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

func (es EStream) Filtered(p func(elem e_sliceOfInts) bool) EStream {
	var h e_sliceOfInts
	for es != nil {
		h, es = es()
		if p(h) {
			return func() (e_sliceOfInts, EStream) {
				return h, es.Filtered(p)
			}
		}
	}
	return nil
}

func (es EStream) PrecededBy(a e_sliceOfInts) EStream {
	return func() (e_sliceOfInts, EStream) {
		return a, es
	}
}

func (es EStream) SuccedeedBy(a e_sliceOfInts) EStream {
	return es.FollowedBy(SliceOfInts_Single(a))
}

func (es1 EStream) FollowedBy(es2 EStream) EStream {
	if es1 != nil {
		return func() (e_sliceOfInts, EStream) {
			h, t := es1()
			return h, t.FollowedBy(es2)
		}
	} else {
		return es2
	}
}

func (es EStream) ForAll(p func(e_sliceOfInts) bool) bool {
	z := true
	var h e_sliceOfInts
	for es != nil && z {
		h, es = es()
		z = p(h)
	}
	return z
}

func (es EStream) ForAny(p func(e_sliceOfInts) bool) bool {
	return es.ForAll(func(e e_sliceOfInts) bool {
		return !p(e)
	})
}

func (es1 EStream) IsEqualTo(es2 EStream) bool {
	var h1, h2 e_sliceOfInts
	for es1 != nil && es2 != nil {
		h1, es1 = es1()
		h2, es2 = es2()
		if !sliceOfInts_equality(h1, h2) {
			return false
		}
	}
	return es1 == nil && es2 == nil
}

func (es EStream) AppendToSlice(s []e_sliceOfInts) []e_sliceOfInts {
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

func (es EStream) ToSlice(initialCapacity int) []e_sliceOfInts {
	slice := make([]e_sliceOfInts, 0, initialCapacity)
	return es.AppendToSlice(slice)
}
