package stream

import (
	"fmt"
	"reflect"
)

// The type of the elements in the stream
type e_sliceOfStrings = []string

func sliceOfStrings_equality(e1, e2 e_sliceOfStrings) bool {
	return reflect.DeepEqual(e1, e2)
}

// The type of the stream whose elements are of type `e_sliceOfStrings`
type SliceOfStrings func() (e_sliceOfStrings, SliceOfStrings)

func SliceOfStrings_Empty() SliceOfStrings {
	return nil
}

func SliceOfStrings_Single(e e_sliceOfStrings) SliceOfStrings {
	return func() (e_sliceOfStrings, SliceOfStrings) {
		return e, nil
	}
}

func SliceOfStrings_Forever(e e_sliceOfStrings) SliceOfStrings {
	return func() (e_sliceOfStrings, SliceOfStrings) {
		return e, SliceOfStrings_Forever(e)
	}
}

func SliceOfStrings_FromSlice(slice []e_sliceOfStrings) SliceOfStrings {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_sliceOfStrings, SliceOfStrings) {
			return slice[0], SliceOfStrings_FromSlice(slice[1:])
		}
	}
}

// func SliceOfStrings_FromSet(m map[e_sliceOfStrings]bool) SliceOfStrings {
// 	slice := make([]e_sliceOfStrings, len(m))
// 	for k := range m {
// 		slice = append(slice, k)
// 	}
// 	return SliceOfStrings_FromSlice(slice)
// }

////

func (es SliceOfStrings) IsEmpty() bool {
	return es == nil
}

func (es SliceOfStrings) Size() int {
	count := 0
	for es != nil {
		count += 1
		_, es = es()
	}
	return count
}

func (es SliceOfStrings) TakeWhile(indexBase int, p func(elem e_sliceOfStrings, index int) bool) SliceOfStrings {
	if es == nil {
		return nil
	} else {
		h, t := es()
		if p(h, indexBase) {
			return func() (e_sliceOfStrings, SliceOfStrings) {
				return h, t.TakeWhile(indexBase+1, p)
			}
		} else {
			return nil
		}
	}
}

func (es SliceOfStrings) DropWhile(indexBase int, p func(elem e_sliceOfStrings, index int) bool) SliceOfStrings {
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

func (es SliceOfStrings) Filtered(p func(elem e_sliceOfStrings) bool) SliceOfStrings {
	var h e_sliceOfStrings
	for es != nil {
		h, es = es()
		if p(h) {
			return func() (e_sliceOfStrings, SliceOfStrings) {
				return h, es.Filtered(p)
			}
		}
	}
	return nil
}

func (es SliceOfStrings) PrecededBy(a e_sliceOfStrings) SliceOfStrings {
	return func() (e_sliceOfStrings, SliceOfStrings) {
		return a, es
	}
}

func (es SliceOfStrings) SuccedeedBy(a e_sliceOfStrings) SliceOfStrings {
	return es.FollowedBy(SliceOfStrings_Single(a))
}

func (es1 SliceOfStrings) FollowedBy(es2 SliceOfStrings) SliceOfStrings {
	if es1 != nil {
		return func() (e_sliceOfStrings, SliceOfStrings) {
			h, t := es1()
			return h, t.FollowedBy(es2)
		}
	} else {
		return es2
	}
}

func (es SliceOfStrings) ForAll(p func(e_sliceOfStrings) bool) bool {
	z := true
	var h e_sliceOfStrings
	for es != nil && z {
		h, es = es()
		z = p(h)
	}
	return z
}

func (es SliceOfStrings) ForAny(p func(e_sliceOfStrings) bool) bool {
	return es.ForAll(func(e e_sliceOfStrings) bool {
		return !p(e)
	})
}

func (es1 SliceOfStrings) IsEqualTo(es2 SliceOfStrings) bool {
	return es1.Corresponds_sliceOfStrings(es2, sliceOfStrings_equality)
	// var h1, h2 e_sliceOfStrings
	// for es1 != nil && es2 != nil {
	// 	h1, es1 = es1()
	// 	h2, es2 = es2()
	// 	if !sliceOfStrings_equality(h1, h2) {
	// 		return false
	// 	}
	// }
	// return es1 == nil && es2 == nil
}

func (es SliceOfStrings) AppendToSlice(s []e_sliceOfStrings) []e_sliceOfStrings {
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

func (es SliceOfStrings) ToSlice(initialCapacity int) []e_sliceOfStrings {
	slice := make([]e_sliceOfStrings, 0, initialCapacity)
	return es.AppendToSlice(slice)
}

//// implementation of PartialFunction[int, e_sliceOfStrings] ////

func (es SliceOfStrings) ApplyOrElse(index int, defaultValue func() e_sliceOfStrings) e_sliceOfStrings {
	if index < 0 {
		return defaultValue()
	}
	var h e_sliceOfStrings
	for es != nil {
		h, es = es()
		if index == 0 {
			return h
		}
		index -= 1
	}
	return defaultValue()
}

func (es SliceOfStrings) Apply(index int) (e_sliceOfStrings, error) {
	var err error
	v := es.ApplyOrElse(index, func() e_sliceOfStrings {
		err = fmt.Errorf("index out of bounds: %v", index)
		var zero e_sliceOfStrings
		return zero
	})
	return v, err
}
