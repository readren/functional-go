package fung

// #dependsOn {"typeCtor":"Stream", "baseTArgs":[], "funcTArgs":[{"type":"eType"}]}

// #importAnchor

// #excludeSectionBegin These lines are not included in the generated source files. They exist to make the template file compiler friendly.
import "fmt"

// #excludeSectionEnd

////

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

// #startOfFuncsWithNoInternalDependants

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es Stream_eType) PrecededBy(a eType) Stream_eType {
	return func() (eType, Stream_eType) {
		return a, es
	}
}

// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es Stream_eType) SuccedeedBy(e eType) Stream_eType {
	return es.FollowedBy(Stream__Single__eType(e))
}

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

// #dependsOn {"typeCtor":"Stream", "baseTArgs": [{"type":"eType"}], "funcTArgs": [{"type":"eType"}]}
func (es1 Stream_eType) IsEqualTo(es2 Stream_eType, elemEquality func(eType, eType) bool) bool {
	return es1.Corresponds__eType(es2, elemEquality)
}

// `ApplyOrElse` gives the element at the specified index or the value returned by the supplied `defaultValue` function if said index is out of the bounds [0, size).
// The purpose of this combinator is to make `Stream` implement the `PartialFunction[int, eType]` interface.
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

// `Apply` gives the element at the specified index or panics if said index is out of the bounds [0, size).
// The purpose of this combinator is to make `Stream` implement the `Func1[int, eType]` interface.
// CAUTION: this method traverses `index + 1` elements of this stream. Avoid it for long stream if possible.
// #usesExternalPackage {"path":"fmt"}
func (es Stream_eType) Apply(index int) eType {
	return es.ApplyOrElse(index, func() eType {
		panic(fmt.Errorf("index out of bounds: %v", index))
	})
}
