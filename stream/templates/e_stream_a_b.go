package templates

// DO NOT COPY contained lines - BEGIN. They exist to make the compiler happy.

type someTypeB struct{}

// The type of the stream whose elements are of type `a_type`
type BStream func() (b_type, BStream)

func (as BStream) FollowedBy(BStream) BStream { panic("This template line should have been removed") }

// DO NOT COPY contained lines - END

// The first type parameter of the methods
//type a_type = string

// The second type parameter of the methods
type b_type = someTypeB

func (es EStream) AppendToMap_a_b(m map[a_type]b_type, g func(e_type) (a_type, b_type)) map[a_type]b_type {
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

func (es EStream) Combined_a_b(as AStream, indexBase int, f func(e e_type, a a_type, index int) b_type) BStream {
	if es == nil || as == nil {
		return nil
	} else {
		return func() (b_type, BStream) {
			he, te := es()
			ha, ta := as()
			return f(he, ha, indexBase), te.Combined_a_b(ta, indexBase+1, f)
		}
	}
}
