package stream

// The first type parameter of the methods
// type a_string = string

func (es String) Collected_string(f func(e_string) (bool, a_string)) String {
	var e e_string
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

func (es String) CollectedKEI_string(baseIndex int, f func(e e_string, index int) (bool, a_string)) String {
	var e e_string
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

func (es String) Mapped_string(f func(e_string) a_string) String {
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
func (es String) MappedKEI_string(indexBase int, f func(e e_string, index int) a_string) String {
	if es == nil {
		return nil
	} else {
		return func() (a_string, String) {
			h, t := es()
			return f(h, indexBase), t.MappedKEI_string(indexBase+1, f)
		}
	}
}

func (es String) Bound_string(f func(e_string) String) String {
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
func (es String) BoundKEI_string(indexBase int, f func(e e_string, index int) String) String {
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

func (es String) FoldLeft_string(z a_string, f func(a_string, e_string) a_string) a_string {
	var h e_string
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es String) FoldRight_string(f func(e_string, a_string) a_string, z a_string) a_string {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_string(f, z))
	}
}

func (es String) Corresponds_string(as String, f func(e e_string, a a_string) bool) bool {
	var e e_string
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
