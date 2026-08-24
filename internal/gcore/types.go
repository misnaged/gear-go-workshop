package gcore

const zeroValuePtr uint32 = 0xffffffff

type SignalCodeValue uint32

const (
	SignalUserspacePanic         SignalCodeValue = 100
	SignalRanOutOfGas            SignalCodeValue = 101
	SignalBackendError           SignalCodeValue = 102
	SignalMemoryOverflow         SignalCodeValue = 103
	SignalUnreachableInstruction SignalCodeValue = 104
	SignalStackLimitExceeded     SignalCodeValue = 105
	SignalRemovedFromWaitlist    SignalCodeValue = 200

	SignalUnsupported SignalCodeValue = 0xffffffff
)

type ActorID [32]byte

type MessageID [32]byte

type ReservationID [32]byte

type ReplyCodeBytes [4]byte

func (r ReplyCodeBytes) IsSuccess() bool {
	return r[0] == 0
}

func (r ReplyCodeBytes) IsError() bool {
	return r[0] == 1
}

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
