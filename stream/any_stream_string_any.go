package stream

// The first type parameter of the methods
// type a_string = string

// The second type parameter of the methods
type b_any = interface{}

func (es Any) AppendToMap_string_any(m map[a_string]b_any, g func(e_any) (a_string, b_any)) map[a_string]b_any {
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
