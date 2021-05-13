package stream

// The first type parameter of the methods
type a_sliceOfInts = []int

func (es SliceOfInts) Collected_sliceOfInts(f func(e_sliceOfInts) (bool, a_sliceOfInts)) SliceOfInts {
	var e e_sliceOfInts
	for es != nil {
		e, es = es()
		isIncluded, a := f(e)
		if isIncluded {
			return func() (a_sliceOfInts, SliceOfInts) {
				return a, es.Collected_sliceOfInts(f)
			}
		}
	}
	return nil
}

func (es SliceOfInts) CollectedKEI_sliceOfInts(baseIndex int, f func(e e_sliceOfInts, index int) (bool, a_sliceOfInts)) SliceOfInts {
	var e e_sliceOfInts
	for es != nil {
		e, es = es()
		c, a := f(e, baseIndex)
		if c {
			return func() (a_sliceOfInts, SliceOfInts) {
				return a, es.CollectedKEI_sliceOfInts(baseIndex+1, f)
			}
		}
	}
	return nil
}

func (es SliceOfInts) Mapped_sliceOfInts(f func(e_sliceOfInts) a_sliceOfInts) SliceOfInts {
	if es == nil {
		return nil
	} else {
		return func() (a_sliceOfInts, SliceOfInts) {
			h, t := es()
			return f(h), t.Mapped_sliceOfInts(f)
		}
	}
}

// KEI stands for knowing elements indexes
func (es SliceOfInts) MappedKEI_sliceOfInts(indexBase int, f func(e e_sliceOfInts, index int) a_sliceOfInts) SliceOfInts {
	if es == nil {
		return nil
	} else {
		return func() (a_sliceOfInts, SliceOfInts) {
			h, t := es()
			return f(h, indexBase), t.MappedKEI_sliceOfInts(indexBase+1, f)
		}
	}
}

func (es SliceOfInts) Bound_sliceOfInts(f func(e_sliceOfInts) SliceOfInts) SliceOfInts {
	if es == nil {
		return nil
	} else {
		return func() (a_sliceOfInts, SliceOfInts) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Bound_sliceOfInts(f))
		}
	}
}

// KEI stands for knowing elements indexes
func (es SliceOfInts) BoundKEI_sliceOfInts(indexBase int, f func(e e_sliceOfInts, index int) SliceOfInts) SliceOfInts {
	if es == nil {
		return nil
	} else {
		return func() (a_sliceOfInts, SliceOfInts) {
			he, te := es()
			ha, ta := f(he, indexBase)()
			return ha, ta.FollowedBy(te.BoundKEI_sliceOfInts(indexBase+1, f))
		}
	}
}

func (es SliceOfInts) FoldLeft_sliceOfInts(z a_sliceOfInts, f func(a_sliceOfInts, e_sliceOfInts) a_sliceOfInts) a_sliceOfInts {
	var h e_sliceOfInts
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es SliceOfInts) FoldRight_sliceOfInts(f func(e_sliceOfInts, a_sliceOfInts) a_sliceOfInts, z a_sliceOfInts) a_sliceOfInts {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_sliceOfInts(f, z))
	}
}

func (es SliceOfInts) Corresponds_sliceOfInts(as SliceOfInts, f func(e e_sliceOfInts, a a_sliceOfInts) bool) bool {
	var e e_sliceOfInts
	var a a_sliceOfInts
	for es != nil && as != nil {
		e, es = es()
		a, as = as()
		if !f(e, a) {
			return false
		}
	}
	return es == nil && as == nil

}
