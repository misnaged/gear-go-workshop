package gcore

const zeroValuePtr uint32 = 0xffffffff

type ActorID [32]byte
type MessageID [32]byte

type Uint128 struct {
	Lo uint64
	Hi uint64
}
type hashWithValue struct {
	Hash  ActorID
	Value Uint128
}
