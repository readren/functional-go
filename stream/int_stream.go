package stream

type A_Int = int
type Int_Stream func() (A_Int, Int_Stream)

func Int_Unit(a A_Int) Int_Stream {
	return func() (A_Int, Int_Stream) {
		return a, nil
	}
}
func (as Int_Stream) IsEmpty() bool {
	return as == nil
}

func (as Int_Stream) Filtered(p func(A_Int) bool) Int_Stream {
	if as == nil {
		return nil
	} else {
		h, t := as()
		pass := p(h)
		for !pass && t != nil {
			h, t = t()
			pass = p(h)
		}
		if pass {
			return func() (A_Int, Int_Stream) {
				return h, t.Filtered(p)
			}
		} else {
			return nil
		}
	}
}

func (as Int_Stream) PrecededBy(a A_Int) Int_Stream {
	return func() (A_Int, Int_Stream) {
		return a, as
	}
}
func (as Int_Stream) SuccedeedBy(a A_Int) Int_Stream {
	return as.FollowedBy(Int_Unit(a))
}
func (as1 Int_Stream) FollowedBy(as2 Int_Stream) Int_Stream {
	if as1 != nil {
		h, t := as1()
		return func() (A_Int, Int_Stream) { return h, t.FollowedBy(as2) }
	} else {
		return as2
	}
}
func (as1 Int_Stream) IsEqualTo(as2 Int_Stream) bool {
	if as1 == nil {
		return as2 == nil
	} else if as2 == nil {
		return false
	} else {
		h1, t1 := as1()
		h2, t2 := as2()
		return h1 == h2 && t1.IsEqualTo(t2)
	}
}

func (as Int_Stream) AppendToSlice(s []A_Int) []A_Int {
	if as != nil {
		h, t := as()
		// All the following lines could be replaced by this << return t.AppendToSlice(append(s, h)) >> if the golang compiler supported tail recursion optimization.
		s = append(s, h)
		for t != nil {
			h, t = t()
			s = append(s, h)
		}
	}
	return s
}

func (as Int_Stream) ToSlice(initialCapacity int) []A_Int {
	slice := make([]A_Int, 0, initialCapacity)
	return as.AppendToSlice(slice)
}
