package fung

type aType struct{}
type bType struct{}
type eType struct{}

type Stream_eType func() (eType, Stream_eType)

func Stream__Single__eType(e eType) Stream_eType {
	panic(1)
}

type Stream_bType func() (bType, Stream_bType)

type Stream_mapFrom_aType_to_slice_bType func() (map[aType][]bType, Stream_mapFrom_aType_to_slice_bType)

func (as Stream_aType) FollowedBy(Stream_aType) Stream_aType {
	panic(1)
}
func (bs Stream_bType) FollowedBy(Stream_bType) Stream_bType {
	panic(1)
}
func (bs Stream_bType) SucceddedBy(bType) Stream_bType {
	panic(1)
}
func (es1 Stream_eType) Corresponds__eType(es2 Stream_eType, f func(e1 eType, e2 eType) bool) bool {
	panic(1)
}
