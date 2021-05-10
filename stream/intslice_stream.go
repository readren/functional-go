package stream

import "reflect"

// The type of the elements in the stream
type e_intSlice = []int

// The type of the stream itself
type IntSlice func() (e_intSlice, IntSlice)

func IntSlice_Empty() IntSlice {
	return nil
}
func IntSlice_Unit(a e_intSlice) IntSlice {
	return func() (e_intSlice, IntSlice) {
		return a, nil
	}
}
func IntSlice_FromSlice(slice []e_intSlice) IntSlice {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_intSlice, IntSlice) {
			return slice[0], IntSlice_FromSlice(slice[1:])
		}
	}
}

// func IntSlice_FromSet(m map[e_intSlice]bool) IntSlice {
// 	slice := make([]e_intSlice, len(m))
// 	for k := range m {
// 		slice = append(slice, k)
// 	}
// 	return IntSlice_FromSlice(slice)
// }

////

func (es IntSlice) IsEmpty() bool {
	return es == nil
}

func (es IntSlice) Filtered(p func(e_intSlice) bool) IntSlice {
	var h e_intSlice
	for es != nil {
		h, es = es()
		if p(h) {
			return func() (e_intSlice, IntSlice) {
				return h, es.Filtered(p)
			}
		}
	}
	return nil
}

func (es IntSlice) PrecededBy(a e_intSlice) IntSlice {
	return func() (e_intSlice, IntSlice) {
		return a, es
	}
}
func (es IntSlice) SuccedeedBy(a e_intSlice) IntSlice {
	return es.FollowedBy(IntSlice_Unit(a))
}
func (es1 IntSlice) FollowedBy(es2 IntSlice) IntSlice {
	if es1 != nil {
		return func() (e_intSlice, IntSlice) {
			h, t := es1()
			return h, t.FollowedBy(es2)
		}
	} else {
		return es2
	}
}

func (es IntSlice) ForAll(p func(e_intSlice) bool) bool {
	z := true
	var h e_intSlice
	for es != nil && z {
		h, es = es()
		z = p(h)
	}
	return z
}

func (es IntSlice) ForAny(p func(e_intSlice) bool) bool {
	return es.ForAll(func(e e_intSlice) bool {
		return !p(e)
	})
}

func (es1 IntSlice) IsEqualTo(es2 IntSlice) bool {
	var h1, h2 e_intSlice
	for es1 != nil && es2 != nil {
		h1, es1 = es1()
		h2, es2 = es2()
		if !reflect.DeepEqual(h1, h2) {
			return false
		}
	}
	return es1 == nil && es2 == nil
}

func (es IntSlice) AppendToSlice(s []e_intSlice) []e_intSlice {
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

func (es IntSlice) ToSlice(initialCapacity int) []e_intSlice {
	slice := make([]e_intSlice, 0, initialCapacity)
	return es.AppendToSlice(slice)
}
