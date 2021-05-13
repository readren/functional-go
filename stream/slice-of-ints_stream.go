package stream

import (
	"fmt"
	"reflect"
)

// The type of the elements in the stream
type e_sliceOfInts = []int

func type_equality(e1, e2 e_sliceOfInts) bool {
	return reflect.DeepEqual(e1, e2)
}

// The type of the stream whose elements are of type `e_sliceOfInts`
type SliceOfInts func() (e_sliceOfInts, SliceOfInts)

func SliceOfInts_Empty() SliceOfInts {
	return nil
}

func SliceOfInts_Single(e e_sliceOfInts) SliceOfInts {
	return func() (e_sliceOfInts, SliceOfInts) {
		return e, nil
	}
}

func SliceOfInts_Forever(e e_sliceOfInts) SliceOfInts {
	return func() (e_sliceOfInts, SliceOfInts) {
		return e, SliceOfInts_Forever(e)
	}
}

func SliceOfInts_FromSlice(slice []e_sliceOfInts) SliceOfInts {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_sliceOfInts, SliceOfInts) {
			return slice[0], SliceOfInts_FromSlice(slice[1:])
		}
	}
}

// func SliceOfInts_FromSet(m map[e_sliceOfInts]bool) SliceOfInts {
// 	slice := make([]e_sliceOfInts, len(m))
// 	for k := range m {
// 		slice = append(slice, k)
// 	}
// 	return SliceOfInts_FromSlice(slice)
// }

////

func (es SliceOfInts) IsEmpty() bool {
	return es == nil
}

func (es SliceOfInts) Size() int {
	count := 0
	for es != nil {
		count += 1
		_, es = es()
	}
	return count
}

func (es SliceOfInts) TakeWhile(indexBase int, p func(elem e_sliceOfInts, index int) bool) SliceOfInts {
	if es == nil {
		return nil
	} else {
		h, t := es()
		if p(h, indexBase) {
			return func() (e_sliceOfInts, SliceOfInts) {
				return h, t.TakeWhile(indexBase+1, p)
			}
		} else {
			return nil
		}
	}
}

func (es SliceOfInts) DropWhile(indexBase int, p func(elem e_sliceOfInts, index int) bool) SliceOfInts {
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

func (es SliceOfInts) Filtered(p func(elem e_sliceOfInts) bool) SliceOfInts {
	var h e_sliceOfInts
	for es != nil {
		h, es = es()
		if p(h) {
			return func() (e_sliceOfInts, SliceOfInts) {
				return h, es.Filtered(p)
			}
		}
	}
	return nil
}

func (es SliceOfInts) PrecededBy(a e_sliceOfInts) SliceOfInts {
	return func() (e_sliceOfInts, SliceOfInts) {
		return a, es
	}
}

func (es SliceOfInts) SuccedeedBy(a e_sliceOfInts) SliceOfInts {
	return es.FollowedBy(SliceOfInts_Single(a))
}

func (es1 SliceOfInts) FollowedBy(es2 SliceOfInts) SliceOfInts {
	if es1 != nil {
		return func() (e_sliceOfInts, SliceOfInts) {
			h, t := es1()
			return h, t.FollowedBy(es2)
		}
	} else {
		return es2
	}
}

func (es SliceOfInts) ForAll(p func(e_sliceOfInts) bool) bool {
	z := true
	var h e_sliceOfInts
	for es != nil && z {
		h, es = es()
		z = p(h)
	}
	return z
}

func (es SliceOfInts) ForAny(p func(e_sliceOfInts) bool) bool {
	return es.ForAll(func(e e_sliceOfInts) bool {
		return !p(e)
	})
}

func (es1 SliceOfInts) IsEqualTo(es2 SliceOfInts) bool {
	return es1.Corresponds_sliceOfInts(es2, type_equality)
	// var h1, h2 e_sliceOfInts
	// for es1 != nil && es2 != nil {
	// 	h1, es1 = es1()
	// 	h2, es2 = es2()
	// 	if !type_equality(h1, h2) {
	// 		return false
	// 	}
	// }
	// return es1 == nil && es2 == nil
}

func (es SliceOfInts) AppendToSlice(s []e_sliceOfInts) []e_sliceOfInts {
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

func (es SliceOfInts) ToSlice(initialCapacity int) []e_sliceOfInts {
	slice := make([]e_sliceOfInts, 0, initialCapacity)
	return es.AppendToSlice(slice)
}

//// implementation of PartialFunction[int, e_sliceOfInts] ////

func (es SliceOfInts) ApplyOrElse(index int, defaultValue func() e_sliceOfInts) e_sliceOfInts {
	if index < 0 {
		return defaultValue()
	}
	var h e_sliceOfInts
	for es != nil {
		h, es = es()
		if index == 0 {
			return h
		}
		index -= 1
	}
	return defaultValue()
}

func (es SliceOfInts) Apply(index int) (e_sliceOfInts, error) {
	var err error
	v := es.ApplyOrElse(index, func() e_sliceOfInts {
		err = fmt.Errorf("index out of bounds: %v", index)
		var zero e_sliceOfInts
		return zero
	})
	return v, err
}
