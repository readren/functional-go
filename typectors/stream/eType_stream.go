package functional

import (
	"fmt"
)

// #exclude-section-begin These lines are not included in the generated source files. They exist to make the template file compiler friendly.

// The type of the elements in the EStream
type eType = struct{}

func (es1 Stream_eType) Corresponds_eType(es2 Stream_eType, f func(e1 eType, e2 eType) bool) bool {
	panic("This template line should have been removed")
}

// #exclude-section-end

// The type of the stream whose elements are of type `eType`
type Stream_eType func() (eType, Stream_eType)

func Stream_eType__Empty() Stream_eType {
	return nil
}

func Stream_eType__Single(e eType) Stream_eType {
	return func() (eType, Stream_eType) {
		return e, nil
	}
}

func Stream_eType__Forever(e eType) Stream_eType {
	return func() (eType, Stream_eType) {
		return e, Stream_eType__Forever(e)
	}
}

func Stream_eType__FromSlice(slice []eType) Stream_eType {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (eType, Stream_eType) {
			return slice[0], Stream_eType__FromSlice(slice[1:])
		}
	}
}

func Stream_eType__FromSet(m map[eType]bool) Stream_eType {
	slice := make([]eType, len(m))
	for k := range m {
		slice = append(slice, k)
	}
	return Stream_eType__FromSlice(slice)
}

////

func (es Stream_eType) IsEmpty() bool {
	return es == nil
}

// CAUTION: this method traverses all the elements of the stream. Avoid it if possible for long streams.
func (es Stream_eType) SizeAndLast() (int, eType) {
	count := 0
	var e eType
	for es != nil {
		count += 1
		e, es = es()
	}
	return count, e
}

// Note: this method is almost lazy. Applying it only traverses the first element of this stream.
func (es Stream_eType) TakeWhile(indexBase int, p func(elem eType, index int) bool) Stream_eType {
	if es == nil {
		return nil
	} else {
		h, t := es()
		if p(h, indexBase) {
			return func() (eType, Stream_eType) {
				return h, t.TakeWhile(indexBase+1, p)
			}
		} else {
			return nil
		}
	}
}

// Note: this method is almost lazy. Applying it only traverses the first element of this stream.
func (es Stream_eType) Take(num int) Stream_eType {
	return es.TakeWhile(0, func(_ eType, index int) bool {
		return index < num
	})
}

// Note: this method is partially lazy. Applying it traverses all the droped and the first non-droped elements of this stream.
func (es Stream_eType) DropWhile(indexBase int, p func(elem eType, index int) bool) Stream_eType {
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

// Note: this method is partially lazy. Applying it traverses all the droped and the first non-droped elements of this stream.
func (es Stream_eType) Drop(num int) Stream_eType {
	return es.DropWhile(0, func(_ eType, index int) bool {
		return index < num
	})
}

// Note: this method is partially lazy. Applying it traverses the first excluded elements of this stream until the first included element inclusive.
func (es Stream_eType) Filter(p func(elem eType) bool) Stream_eType {
	var h eType
	for es != nil {
		h, es = es()
		if p(h) {
			return func() (eType, Stream_eType) {
				return h, es.Filter(p)
			}
		}
	}
	return nil
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es Stream_eType) PrecededBy(a eType) Stream_eType {
	return func() (eType, Stream_eType) {
		return a, es
	}
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es Stream_eType) SuccedeedBy(a eType) Stream_eType {
	return es.FollowedBy(Stream_eType__Single(a))
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es1 Stream_eType) FollowedBy(es2 Stream_eType) Stream_eType {
	if es1 != nil {
		return func() (eType, Stream_eType) {
			h, t := es1()
			return h, t.FollowedBy(es2)
		}
	} else {
		return es2
	}
}

func (es Stream_eType) ForAll(p func(eType) bool) bool {
	z := true
	var h eType
	for es != nil && z {
		h, es = es()
		z = p(h)
	}
	return z
}

func (es Stream_eType) ForAny(p func(eType) bool) bool {
	return es.ForAll(func(e eType) bool {
		return !p(e)
	})
}

// #requires {"typeCtor":"stream", "baseTArgs": [{"type":"eType"}], "methodTArgs": [{"type":"eType"}]}
func (es1 Stream_eType) IsEqualTo(es2 Stream_eType, elemEquality func(eType, eType) bool) bool {
	return es1.Corresponds_eType(es2, elemEquality)
}

func (es Stream_eType) AppendToSlice(s []eType) []eType {
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

func (es Stream_eType) ToSlice(initialCapacity int) []eType {
	slice := make([]eType, 0, initialCapacity)
	return es.AppendToSlice(slice)
}

//// implementation of PartialFunction[int, eType] ////

// CAUTION: this method traverses `index + 1` elements of this stream. Avoid it for long stream if possible.
func (es Stream_eType) ApplyOrElse(index int, defaultValue func() eType) eType {
	if index < 0 {
		return defaultValue()
	}
	var h eType
	for es != nil {
		h, es = es()
		if index == 0 {
			return h
		}
		index -= 1
	}
	return defaultValue()
}

// CAUTION: this method traverses `index + 1` elements of this stream. Avoid it for long stream if possible.
func (es Stream_eType) Apply(index int) (eType, error) {
	var err error
	v := es.ApplyOrElse(index, func() eType {
		err = fmt.Errorf("index out of bounds: %v", index)
		var zero eType
		return zero
	})
	return v, err
}
