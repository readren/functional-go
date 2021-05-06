package int_stream

type K = string
type V = int

func (as Stream) AppendToMap_string_int(m map[K]V, g func(A) (K, V)) map[K]V {
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
