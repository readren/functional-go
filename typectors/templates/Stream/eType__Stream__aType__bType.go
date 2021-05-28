package fung

// #importAnchor

// #startOfFuncsWithNoInternalDependants

// Returns a stream of elements of type `map[aType][]bType` (map from `aType` to slices of `bType`) where the element at index `i` is the grouped accumulation of the first `i` elements of this stream.
// In other words: Returns a stream of the same size than this stream, and whose element at index `i` is equivalent to `this.Take(i+1).GroupMap_aType_bType(accumulator, g)`, for all `i` between 0 and the size of this stream.
// Note: this method is fully lazy. Applying it traverses no element of this stream.
// #dependsOn {"typeCtor":"Stream", "baseTArgs": [{"type":"map[aType][]bType", "typeName":"mapFrom_aType_to_slice_bType"}]}
func (es Stream_eType) GroupMapped__aType__bType(accumulator map[aType][]bType, g func(elem eType) (key aType, value bType)) Stream_mapFrom_aType_to_slice_bType {
	if es == nil {
		return nil
	} else {
		return func() (map[aType][]bType, Stream_mapFrom_aType_to_slice_bType) {
			h, t := es()
			k, v := g(h)
			accumulator[k] = append(accumulator[k], v)
			return accumulator, t.GroupMapped__aType__bType(accumulator, g)
		}
	}
}

// For each element of this stream, obtain a key-value pair by applying the function `g` to the element and append the value to the slice associated with the key in the map `accumulator`.
// Returns the accumulator.
func (es Stream_eType) GroupMap__aType__bType(accumulator map[aType][]bType, g func(elem eType) (key aType, value bType)) map[aType][]bType {
	var e eType
	for es != nil {
		e, es = es()
		k, v := g(e)
		accumulator[k] = append(accumulator[k], v)
	}
	return accumulator
}

// For each element of this stream, apply the function `g` to the element and put the resulting key-value pair into the map `m`. If a value already exists at the key, it is replaced with the result of applying the reducing function `r` to the old and new values.
func (es Stream_eType) GroupMapReduce__aType__bType(accumulator map[aType]bType, g func(elem eType) (key aType, value bType), r func(bType, bType) bType) map[aType]bType {
	var e eType
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

// #dependsOn {"typeCtor":"Stream", "baseTArgs": [{"type":"aType"}]}
// #dependsOn {"typeCtor":"Stream", "baseTArgs": [{"type":"bType"}]}
func (es Stream_eType) Combined__aType__bType(as Stream_aType, indexBase int, f func(e eType, a aType, index int) bType) Stream_bType {
	if es == nil || as == nil {
		return nil
	} else {
		return func() (bType, Stream_bType) {
			he, te := es()
			ha, ta := as()
			return f(he, ha, indexBase), te.Combined__aType__bType(ta, indexBase+1, f)
		}
	}
}
