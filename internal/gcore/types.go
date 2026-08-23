package gcore

const zeroValuePtr uint32 = 0xffffffff

type ActorID [32]byte

type MessageID [32]byte

type ReservationID [32]byte

type Uint128 struct {
	Lo uint64
	Hi uint64
}
type Percent uint32

type GasMultiplier struct {
	GasPerValue uint64
	ValuePerGas Uint128
}

type EnvVarsStruct struct {
	PerformanceMultiplier uint32
	ExistentialDeposit    Uint128
	MailboxThreshold      uint64
	GasMultiplier         GasMultiplier
}
type hashWithValue struct {
	Hash  ActorID
	Value Uint128
}
type blockNumberWithHash struct {
	BlockNumber uint32
	Hash        [32]byte
}
