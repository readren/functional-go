package stream

// The first type parameter of the methods
//type a_string = string

// The second type parameter of the methods
type b_int = int

func (es Int_Stream) AppendToMap_string_int(m map[a_string]b_int, g func(e_int) (a_string, b_int)) map[a_string]b_int {
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
