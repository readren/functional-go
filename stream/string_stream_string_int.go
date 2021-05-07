package stream

// type K_String = string
// type V_Int = int

func (as String_Stream) AppendToMap_string_int(m map[K_String]V_Int, g func(A_String) (K_String, V_Int)) map[K_String]V_Int {
	if as != nil {
		h, t := as()
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
