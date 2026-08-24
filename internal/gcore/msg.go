package gcore

import (
	"unsafe"

	"github.com/misnaged/gear-go-workshop/internal/gsys"
)

func Source() ActorID {
	var source ActorID

	gsys.Source(
		uint32(uintptr(unsafe.Pointer(&source[0]))),
	)

	return source
}
func ID() MessageID {
	var id MessageID

	gsys.MessageID(
		uint32(uintptr(unsafe.Pointer(&id[0]))),
	)

	return id
}

func Value() Uint128 {
	var value Uint128

	gsys.Value(
		uint32(uintptr(unsafe.Pointer(&value))),
	)

	return value
}

func ReplyTo() (MessageID, error) {
	var result [36]byte

	gsys.ReplyTo(
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrReplyTo
	}

	var id MessageID
	copy(id[:], result[4:])

	return id, nil
}

func Reply(payload []byte) {
	var result [36]byte

	var payloadPtr uint32

	if len(payload) > 0 {
		payloadPtr = uint32(
			uintptr(unsafe.Pointer(&payload[0])),
		)
	}

	resultPtr := uint32(
		uintptr(unsafe.Pointer(&result[0])),
	)

	gsys.Reply(
		payloadPtr,
		uint32(len(payload)),
		zeroValuePtr,
		resultPtr,
	)
}

func Read(buffer []byte) error {
	size := Size()

	if size > len(buffer) {
		return ErrBufferTooSmall
	}

	if size == 0 {
		return nil
	}

	var errorCode uint32

	gsys.Read(
		0,
		uint32(size),
		uint32(uintptr(unsafe.Pointer(&buffer[0]))),
		uint32(uintptr(unsafe.Pointer(&errorCode))),
	)

	if errorCode != 0 {
		return ErrReadFailed
	}

	return nil
}
func ReadAt(offset int, buffer []byte) error {
	if len(buffer) == 0 {
		return nil
	}

	size := Size()

	if size > len(buffer)+offset {
		return ErrBufferTooSmall
	}

	var errorCode uint32

	gsys.Read(
		uint32(offset),
		uint32(len(buffer)),
		uint32(uintptr(unsafe.Pointer(&buffer[0]))),
		uint32(uintptr(unsafe.Pointer(&errorCode))),
	)

	if errorCode != 0 {
		return ErrReadFailed
	}

	return nil
}
func Size() int {
	var size uint32

	gsys.Size(
		uint32(uintptr(unsafe.Pointer(&size))),
	)

	return int(size)
}
func Send(destination ActorID, payload []byte, value Uint128) (MessageID, error) {
	pidValue := hashWithValue{
		Hash:  destination,
		Value: value,
	}

	var result [36]byte

	var payloadPtr uint32
	if len(payload) > 0 {
		payloadPtr = uint32(uintptr(unsafe.Pointer(&payload[0])))
	}

	gsys.Send(
		uint32(uintptr(unsafe.Pointer(&pidValue))),
		payloadPtr,
		uint32(len(payload)),
		0,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrSendFailed
	}

	var id MessageID
	copy(id[:], result[4:])

	return id, nil
}

func ReplyDeposit(messageID MessageID, gas uint64) error {
	var errorCode uint32

	gsys.ReplyDeposit(
		uint32(uintptr(unsafe.Pointer(&messageID[0]))),
		gas,
		uint32(uintptr(unsafe.Pointer(&errorCode))),
	)

	if errorCode != 0 {
		return ErrReplyDepositFailed
	}

	return nil
}

func ReplyCode() (ReplyCodeBytes, error) {
	var result [8]byte

	gsys.ReplyCode(
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return ReplyCodeBytes{}, ErrReplyCodeFailed
	}

	var code ReplyCodeBytes
	copy(code[:], result[4:8])

	return code, nil
}

func SignalCode() (SignalCodeValue, error) {
	var result [8]byte

	gsys.SignalCode(
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return 0, ErrSignalCodeFailed
	}

	code := SignalCodeValue(
		*(*uint32)(unsafe.Pointer(&result[4])),
	)

	switch code {
	case SignalUserspacePanic,
		SignalRanOutOfGas,
		SignalBackendError,
		SignalMemoryOverflow,
		SignalUnreachableInstruction,
		SignalStackLimitExceeded,
		SignalRemovedFromWaitlist,
		SignalUnsupported:
		return code, nil
	default:
		return 0, ErrUnsupportedSignalCode
	}
}

func ReplyFromReservation(reservationID ReservationID, payload []byte, value Uint128) (MessageID, error) {
	ridValue := hashWithValue{
		Hash:  ActorID(reservationID),
		Value: value,
	}

	var result [36]byte

	var payloadPtr uint32
	if len(payload) > 0 {
		payloadPtr = uint32(
			uintptr(unsafe.Pointer(&payload[0])),
		)
	}

	gsys.ReservationReply(
		uint32(uintptr(unsafe.Pointer(&ridValue))),
		payloadPtr,
		uint32(len(payload)),
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrReplyFromReservationFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}
