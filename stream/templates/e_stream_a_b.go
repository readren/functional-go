package templates

// DO NOT COPY contained lines - BEGIN. They exist to make the compiler happy.

type someTypeB struct{}

// The type of the stream whose elements are of type `a_type`
type BStream func() (b_type, BStream)
type MapFromAToSliceOfBsStream func() (map[a_type][]b_type, MapFromAToSliceOfBsStream)

func (bs BStream) FollowedBy(BStream) BStream { panic("This template line should have been removed") }
func (bs BStream) SucceddedBy(b_type) BStream { panic("This template line should have been removed") }

// DO NOT COPY contained lines - END

// The first type parameter of the methods
//type a_type = string

// The second type parameter of the methods
type b_type = someTypeB

// Returns a stream of elements of type `map[a_type][]b_type` (map from `a_type` to slices of `b_type`) where the element at index `i` is the grouped accumulation of the first `i` elements of this stream.
// In other words: Returns a stream of the same size than this stream, and whose element at index `i` is equivalent to `this.Take(i+1).GroupMap_a_b(accumulator, g)`, for all `i` between 0 and the size of this stream.
// Note: this method is fully lazy. Applying it traverses no element of this stream.
func (es EStream) GroupMapped_a_b(accumulator map[a_type][]b_type, g func(elem e_type) (key a_type, value b_type)) MapFromAToSliceOfBsStream {
	if es == nil {
		return nil
	} else {
		return func() (map[a_type][]b_type, MapFromAToSliceOfBsStream) {
			h, t := es()
			k, v := g(h)
			accumulator[k] = append(accumulator[k], v)
			return accumulator, t.GroupMapped_a_b(accumulator, g)
		}
	}
}

// For each element of this stream, obtain a key-value pair by applying the function `g` to the element and append the value to the slice associated with the key in the map `accumulator`.
// Returns the accumulator.
func (es EStream) GroupMap_a_b(accumulator map[a_type][]b_type, g func(elem e_type) (key a_type, value b_type)) map[a_type][]b_type {
	var e e_type
	for es != nil {
		e, es = es()
		k, v := g(e)
		accumulator[k] = append(accumulator[k], v)
	}
	return accumulator
}

// For each element of this stream, apply the function `g` to the element and put the resulting key-value pair into the map `m`. If a value already exists at the key, it is replaced with the result of applying the reducing function `r` to the old and new values.
func (es EStream) GroupMapReduce_a_b(accumulator map[a_type]b_type, g func(elem e_type) (key a_type, value b_type), r func(b_type, b_type) b_type) map[a_type]b_type {
	var e e_type
	for es != nil {
		e, es = es()
		k, v1 := g(e)
		v0, exists := accumulator[k]
		if exists {
			accumulator[k] = r(v0, v1)
		} else {
			accumulator[k] = v1
		}
	}
	return accumulator
}

func (es EStream) Combined_a_b(as AStream, indexBase int, f func(e e_type, a a_type, index int) b_type) BStream {
	if es == nil || as == nil {
		return nil
	} else {
		return func() (b_type, BStream) {
			he, te := es()
			ha, ta := as()
			return f(he, ha, indexBase), te.Combined_a_b(ta, indexBase+1, f)
		}
	}
}
