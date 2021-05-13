package stream

// The first type parameter of the methods
type a_any = interface{}

func (es Any) Collected_any(f func(e_any) (bool, a_any)) Any {
	var e e_any
	for es != nil {
		e, es = es()
		isIncluded, a := f(e)
		if isIncluded {
			return func() (a_any, Any) {
				return a, es.Collected_any(f)
			}
		}
	}
	return nil
}

func (es Any) CollectedKEI_any(baseIndex int, f func(e e_any, index int) (bool, a_any)) Any {
	var e e_any
	for es != nil {
		e, es = es()
		c, a := f(e, baseIndex)
		if c {
			return func() (a_any, Any) {
				return a, es.CollectedKEI_any(baseIndex+1, f)
			}
		}
	}
	return nil
}

func (es Any) Mapped_any(f func(e_any) a_any) Any {
	if es == nil {
		return nil
	} else {
		return func() (a_any, Any) {
			h, t := es()
			return f(h), t.Mapped_any(f)
		}
	}
}

// KEI stands for knowing elements indexes
func (es Any) MappedKEI_any(indexBase int, f func(e e_any, index int) a_any) Any {
	if es == nil {
		return nil
	} else {
		return func() (a_any, Any) {
			h, t := es()
			return f(h, indexBase), t.MappedKEI_any(indexBase+1, f)
		}
	}
}

func (es Any) Bound_any(f func(e_any) Any) Any {
	if es == nil {
		return nil
	} else {
		return func() (a_any, Any) {
			he, te := es()
			ha, ta := f(he)()
			return ha, ta.FollowedBy(te.Bound_any(f))
		}
	}
}

// KEI stands for knowing elements indexes
func (es Any) BoundKEI_any(indexBase int, f func(e e_any, index int) Any) Any {
	if es == nil {
		return nil
	} else {
		return func() (a_any, Any) {
			he, te := es()
			ha, ta := f(he, indexBase)()
			return ha, ta.FollowedBy(te.BoundKEI_any(indexBase+1, f))
		}
	}
}

func (es Any) FoldLeft_any(z a_any, f func(a_any, e_any) a_any) a_any {
	var h e_any
	for es != nil {
		h, es = es()
		z = f(z, h)
	}
	return z
}

func (es Any) FoldRight_any(f func(e_any, a_any) a_any, z a_any) a_any {
	if es == nil {
		return z
	} else {
		h, t := es()
		return f(h, t.FoldRight_any(f, z))
	}
}

func (es Any) Corresponds_any(as Any, f func(e e_any, a a_any) bool) bool {
	var e e_any
	var a a_any
	for es != nil && as != nil {
		e, es = es()
		a, as = as()
		if !f(e, a) {
			return false
		}
	}
	return es == nil && as == nil

}
