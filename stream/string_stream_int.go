package stream

// The first type parameter of the methods
// type a_int = int

func (es String) Collected_int(f func(e_string) (bool, a_int)) Int {
	var e e_string
	for es != nil {
		e, es = es()
		isIncluded, a := f(e)
		if isIncluded {
			return func() (a_int, Int) {
				return a, es.Collected_int(f)
			}
		}
	}
	return nil
}

func (es String) CollectedKEI_int(baseIndex int, f func(e e_string, index int) (bool, a_int)) Int {
	var e e_string
	for es != nil {
		e, es = es()
		c, a := f(e, baseIndex)
		if c {
			return func() (a_int, Int) {
				return a, es.CollectedKEI_int(baseIndex+1, f)
			}
		}
	}
	return nil
}

func (es String) Mapped_int(f func(e_string) a_int) Int {
	if es == nil {
		return nil
	} else {
		return func() (a_int, Int) {
			h, t := es()
			return f(h), t.Mapped_int(f)
		}
	}
}

// KEI stands for knowing elements indexes
func (es String) MappedKEI_int(indexBase int, f func(e e_string, index int) a_int) Int {
	if es == nil {
		return nil
	} else {
		return func() (a_int, Int) {
			h, t := es()
			return f(h, indexBase), t.MappedKEI_int(indexBase+1, f)
		}
	}
}

func (es String) Bound_int(f func(e_string) Int) Int {
	if es == nil {
		return nil
	} else {
		return func() (a_int, Int) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Bound_int(f))
		}
	}
}

// KEI stands for knowing elements indexes
func (es String) BoundKEI_int(indexBase int, f func(e e_string, index int) Int) Int {
	if es == nil {
		return nil
	} else {
		return func() (a_int, Int) {
			he, te := es()
			ha, ta := f(he, indexBase)()
			return ha, ta.FollowedBy(te.BoundKEI_int(indexBase+1, f))
		}
	}
}

func (es String) FoldLeft_int(z a_int, f func(a_int, e_string) a_int) a_int {
	var h e_string
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es String) FoldRight_int(f func(e_string, a_int) a_int, z a_int) a_int {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_int(f, z))
	}
}

func (es String) Corresponds_int(as Int, f func(e e_string, a a_int) bool) bool {
	var e e_string
	var a a_int
	for es != nil && as != nil {
		e, es = es()
		a, as = as()
		if !f(e, a) {
			return false
		}
	}
	return es == nil && as == nil

}
