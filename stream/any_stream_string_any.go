package stream

// type K_String = string
type V_Any = interface{}

func (as Any_Stream) AppendToMap_string_any(m map[K_String]V_Any, g func(A_Any) (K_String, V_Any)) map[K_String]V_Any {
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
