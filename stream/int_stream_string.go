package stream

// The first type parameter of the methods
type a_string = string

func (es Int) Collected_string(f func(e_int) (bool, a_string)) String {
	var e e_int
	for es != nil {
		e, es = es()
		isIncluded, a := f(e)
		if isIncluded {
			return func() (a_string, String) {
				return a, es.Collected_string(f)
			}
		}
	}
	return nil
}

func (es Int) CollectedKEI_string(baseIndex int, f func(e e_int, index int) (bool, a_string)) String {
	var e e_int
	for es != nil {
		e, es = es()
		c, a := f(e, baseIndex)
		if c {
			return func() (a_string, String) {
				return a, es.CollectedKEI_string(baseIndex+1, f)
			}
		}
	}
	return nil
}

func (es Int) Mapped_string(f func(e_int) a_string) String {
	if es == nil {
		return nil
	} else {
		return func() (a_string, String) {
			h, t := es()
			return f(h), t.Mapped_string(f)
		}
	}
}

// KEI stands for knowing elements indexes
func (es Int) MappedKEI_string(indexBase int, f func(e e_int, index int) a_string) String {
	if es == nil {
		return nil
	} else {
		return func() (a_string, String) {
			h, t := es()
			return f(h, indexBase), t.MappedKEI_string(indexBase+1, f)
		}
	}
}

func (es Int) Bound_string(f func(e_int) String) String {
	if es == nil {
		return nil
	} else {
		return func() (a_string, String) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Bound_string(f))
		}
	}
}

// KEI stands for knowing elements indexes
func (es Int) BoundKEI_string(indexBase int, f func(e e_int, index int) String) String {
	if es == nil {
		return nil
	} else {
		return func() (a_string, String) {
			he, te := es()
			ha, ta := f(he, indexBase)()
			return ha, ta.FollowedBy(te.BoundKEI_string(indexBase+1, f))
		}
	}
}

func (es Int) FoldLeft_string(z a_string, f func(a_string, e_int) a_string) a_string {
	var h e_int
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es Int) FoldRight_string(f func(e_int, a_string) a_string, z a_string) a_string {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_string(f, z))
	}
}

func (es Int) Corresponds_string(as String, f func(e e_int, a a_string) bool) bool {
	var e e_int
	var a a_string
	for es != nil && as != nil {
		e, es = es()
		a, as = as()
		if !f(e, a) {
			return false
		}
	}
	return es == nil && as == nil

}
