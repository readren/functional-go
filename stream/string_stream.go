package stream

// The type of the elements in the stream
type e_string = string

// The type of the stream itself
type String func() (e_string, String)

func String_Empty() String {
	return nil
}
func String_Unit(a e_string) String {
	return func() (e_string, String) {
		return a, nil
	}
}
func String_FromSlice(slice []e_string) String {
	if len(slice) == 0 {
		return nil
	} else {
		return func() (e_string, String) {
			return slice[0], String_FromSlice(slice[1:])
		}
	}
}
func String_FromSet(m map[e_string]bool) String {
	slice := make([]e_string, len(m))
	for k := range m {
		slice = append(slice, k)
	}
	return String_FromSlice(slice)
}

////

func (es String) IsEmpty() bool {
	return es == nil
}

func (es String) Filtered(p func(e_string) bool) String {
	var h e_string
	for es != nil {
		h, es = es()
		if p(h) {
			return func() (e_string, String) {
				return h, es.Filtered(p)
			}
		}
	}
	return nil
}

func (es String) PrecededBy(a e_string) String {
	return func() (e_string, String) {
		return a, es
	}
}
func (es String) SuccedeedBy(a e_string) String {
	return es.FollowedBy(String_Unit(a))
}
func (as1 String) FollowedBy(as2 String) String {
	if as1 != nil {
		return func() (e_string, String) {
			h, t := as1()
			return h, t.FollowedBy(as2)
		}
	} else {
		return as2
	}
}

func (es String) ForAll(p func(e_string) bool) bool {
	z := true
	var h e_string
	for es != nil && z {
		h, es = es()
		z = p(h)
	}
	return z
}

func (es String) ForAny(p func(e_string) bool) bool {
	return es.ForAll(func(e e_string) bool {
		return !p(e)
	})
}

func (es1 String) IsEqualTo(es2 String) bool {
	var h1, h2 e_string
	for es1 != nil && es2 != nil {
		h1, es1 = es1()
		h2, es2 = es2()
		if h1 != h2 {
			return false
		}
	}
	return es1 == nil && es2 == nil
}

func (es String) AppendToSlice(s []e_string) []e_string {
	if es != nil {
		h, t := es()
		// All the following lines could be replaced by this << return t.AppendToSlice(append(s, h)) >> if the golang compiler supported tail recursion optimization.
		s = append(s, h)
		for t != nil {
			h, t = t()
			s = append(s, h)
		}
	}
	return s
}

func (es String) ToSlice(initialCapacity int) []e_string {
	slice := make([]e_string, 0, initialCapacity)
	return es.AppendToSlice(slice)
}
