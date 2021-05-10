package stream

// The first type parameter of the methods
//type a_string = string

// The second type parameter of the methods
// type b_int = int

func (es Int) AppendToMap_string_int(m map[a_string]b_int, g func(e_int) (a_string, b_int)) map[a_string]b_int {
	if es != nil {
		h, t := es()
		k, v := g(h)
		m[k] = v
		for t != nil {
			h, t = t()
			k, v = g(h)
			m[k] = v
		}
	}
	return m
}

func (es Int) Combined_string_int(as String, indexBase int, f func(e e_int, a a_string, index int) b_int) Int {
	if es == nil || as == nil {
		return nil
	} else {
		return func() (b_int, Int) {
			he, te := es()
			ha, ta := as()
			return f(he, ha, indexBase), te.Combined_string_int(ta, indexBase+1, f)
		}
	}
}
