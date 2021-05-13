package stream

// The first type parameter of the methods
type a_sliceOfStrings = []string

func (es SliceOfStrings) Collected_sliceOfStrings(f func(e_sliceOfStrings) (bool, a_sliceOfStrings)) SliceOfStrings {
	var e e_sliceOfStrings
	for es != nil {
		e, es = es()
		isIncluded, a := f(e)
		if isIncluded {
			return func() (a_sliceOfStrings, SliceOfStrings) {
				return a, es.Collected_sliceOfStrings(f)
			}
		}
	}
	return nil
}

func (es SliceOfStrings) CollectedKEI_sliceOfStrings(baseIndex int, f func(e e_sliceOfStrings, index int) (bool, a_sliceOfStrings)) SliceOfStrings {
	var e e_sliceOfStrings
	for es != nil {
		e, es = es()
		c, a := f(e, baseIndex)
		if c {
			return func() (a_sliceOfStrings, SliceOfStrings) {
				return a, es.CollectedKEI_sliceOfStrings(baseIndex+1, f)
			}
		}
	}
	return nil
}

func (es SliceOfStrings) Mapped_sliceOfStrings(f func(e_sliceOfStrings) a_sliceOfStrings) SliceOfStrings {
	if es == nil {
		return nil
	} else {
		return func() (a_sliceOfStrings, SliceOfStrings) {
			h, t := es()
			return f(h), t.Mapped_sliceOfStrings(f)
		}
	}
}

// KEI stands for knowing elements indexes
func (es SliceOfStrings) MappedKEI_sliceOfStrings(indexBase int, f func(e e_sliceOfStrings, index int) a_sliceOfStrings) SliceOfStrings {
	if es == nil {
		return nil
	} else {
		return func() (a_sliceOfStrings, SliceOfStrings) {
			h, t := es()
			return f(h, indexBase), t.MappedKEI_sliceOfStrings(indexBase+1, f)
		}
	}
}

func (es SliceOfStrings) Bound_sliceOfStrings(f func(e_sliceOfStrings) SliceOfStrings) SliceOfStrings {
	if es == nil {
		return nil
	} else {
		return func() (a_sliceOfStrings, SliceOfStrings) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Bound_sliceOfStrings(f))
		}
	}
}

// KEI stands for knowing elements indexes
func (es SliceOfStrings) BoundKEI_sliceOfStrings(indexBase int, f func(e e_sliceOfStrings, index int) SliceOfStrings) SliceOfStrings {
	if es == nil {
		return nil
	} else {
		return func() (a_sliceOfStrings, SliceOfStrings) {
			he, te := es()
			ha, ta := f(he, indexBase)()
			return ha, ta.FollowedBy(te.BoundKEI_sliceOfStrings(indexBase+1, f))
		}
	}
}

func (es SliceOfStrings) FoldLeft_sliceOfStrings(z a_sliceOfStrings, f func(a_sliceOfStrings, e_sliceOfStrings) a_sliceOfStrings) a_sliceOfStrings {
	var h e_sliceOfStrings
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es SliceOfStrings) FoldRight_sliceOfStrings(f func(e_sliceOfStrings, a_sliceOfStrings) a_sliceOfStrings, z a_sliceOfStrings) a_sliceOfStrings {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_sliceOfStrings(f, z))
	}
}

func (es SliceOfStrings) Corresponds_sliceOfStrings(as SliceOfStrings, f func(e e_sliceOfStrings, a a_sliceOfStrings) bool) bool {
	var e e_sliceOfStrings
	var a a_sliceOfStrings
	for es != nil && as != nil {
		e, es = es()
		a, as = as()
		if !f(e, a) {
			return false
		}
	}
	return es == nil && as == nil

}
