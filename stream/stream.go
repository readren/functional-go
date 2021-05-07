package stream

type Any interface{}

func Empty() Stream {
	return nil
}

type Stream interface {
	isEmpty() bool
	Filtered(p func(Any) bool) Stream
	PrecededBy(Any) Stream
	SuccedeedBy(Any) Stream
	FollowedBy(as2 Stream) Stream
	IsEqualTo(as2 Stream) bool
	AppendToSlice([]Any) []Any
	ToSlice(initialCapacity int) []Any
}
