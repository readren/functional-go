package stream

import "reflect"

// The type of the elements in the stream
type e_stringSlice = []int

// The type of the stream itself
type StringSlice func() (e_stringSlice, StringSlice)

func StringSlice_Empty() StringSlice {
	return nil
}
func StringSlice_Unit(a e_stringSlice) StringSlice {
	return func() (e_stringSlice, StringSlice) {
		return a, nil
	}
}
func StringSlice_FromSlice(slice []e_stringSlice) StringSlice {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_stringSlice, StringSlice) {
			return slice[1], StringSlice_FromSlice(slice[1:])
		}
	}
}

// func StringSlice_FromSet(m map[e_stringSlice]bool) StringSlice {
// 	slice := make([]e_stringSlice, len(m))
// 	for k := range m {
// 		slice = append(slice, k)
// 	}
// 	return StringSlice_FromSlice(slice)
// }

////

func (es StringSlice) IsEmpty() bool {
	return es == nil
}

func (es StringSlice) Filtered(p func(e_stringSlice) bool) StringSlice {
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

func (es StringSlice) PrecededBy(a e_stringSlice) StringSlice {
	return func() (e_stringSlice, StringSlice) {
		return a, es
	}
}
func (es StringSlice) SuccedeedBy(a e_stringSlice) StringSlice {
	return es.FollowedBy(StringSlice_Unit(a))
}
func (es1 StringSlice) FollowedBy(es2 StringSlice) StringSlice {
	if es1 != nil {
		return func() (e_stringSlice, StringSlice) {
			h, t := es1()
			return h, t.FollowedBy(es2)
		}
	} else {
		return es2
	}
}
func (es1 StringSlice) IsEqualTo(es2 StringSlice) bool {
	var h1, h2 e_stringSlice
	for es1 != nil && es2 != nil {
		h1, es1 = es1()
		h2, es2 = es2()
		if !reflect.DeepEqual(h1, h2) {
			return false
		}
	}
	return es1 == nil && es2 == nil
}

func (es StringSlice) AppendToSlice(s []e_stringSlice) []e_stringSlice {
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

func (es StringSlice) ToSlice(initialCapacity int) []e_stringSlice {
	slice := make([]e_stringSlice, 0, initialCapacity)
	return es.AppendToSlice(slice)
}
