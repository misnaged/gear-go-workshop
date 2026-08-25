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
	return SendDelayed(
		destination,
		payload,
		value,
		0,
	)
}

func SendDelayed(destination ActorID, payload []byte, value Uint128, delay uint32) (MessageID, error) {
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
		delay,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrSendFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

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

func ReplyWithGas(payload []byte, gasLimit uint64, value Uint128) (MessageID, error) {
	var result [36]byte

	var payloadPtr uint32
	if len(payload) > 0 {
		payloadPtr = uint32(
			uintptr(unsafe.Pointer(&payload[0])),
		)
	}

	valuePtr := zeroValuePtr

	if value.Lo != 0 || value.Hi != 0 {
		valuePtr = uint32(
			uintptr(unsafe.Pointer(&value)),
		)
	}

	gsys.ReplyWithGas(
		payloadPtr,
		uint32(len(payload)),
		gasLimit,
		valuePtr,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrReplyWithGasFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func ReplyCommit(value Uint128) (MessageID, error) {
	var result [36]byte

	valuePtr := zeroValuePtr

	if value.Lo != 0 || value.Hi != 0 {
		valuePtr = uint32(
			uintptr(unsafe.Pointer(&value)),
		)
	}

	gsys.ReplyCommit(
		valuePtr,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrReplyCommitFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func ReplyPush(payload []byte) error {
	var errorCode uint32

	var payloadPtr uint32
	if len(payload) > 0 {
		payloadPtr = uint32(
			uintptr(unsafe.Pointer(&payload[0])),
		)
	}

	gsys.ReplyPush(
		payloadPtr,
		uint32(len(payload)),
		uint32(uintptr(unsafe.Pointer(&errorCode))),
	)

	if errorCode != 0 {
		return ErrReplyPushFailed
	}

	return nil
}

func ReplyCommitWithGas(gasLimit uint64, value Uint128) (MessageID, error) {
	var result [36]byte

	valuePtr := zeroValuePtr

	if value.Lo != 0 || value.Hi != 0 {
		valuePtr = uint32(
			uintptr(unsafe.Pointer(&value)),
		)
	}

	gsys.ReplyCommitWithGas(
		gasLimit,
		valuePtr,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrReplyCommitWithGasFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func ReplyCommitFromReservation(reservationID ReservationID, value Uint128) (MessageID, error) {
	ridValue := hashWithValue{
		Hash:  ActorID(reservationID),
		Value: value,
	}

	var result [36]byte

	gsys.ReservationReplyCommit(
		uint32(uintptr(unsafe.Pointer(&ridValue))),
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrReplyCommitFromReservationFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func ReplyInput(value Uint128, offset uint32, length uint32) (MessageID, error) {
	var result [36]byte

	valuePtr := zeroValuePtr

	if value.Lo != 0 || value.Hi != 0 {
		valuePtr = uint32(
			uintptr(unsafe.Pointer(&value)),
		)
	}

	gsys.ReplyInput(
		offset,
		length,
		valuePtr,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrReplyInputFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func ReplyPushInput(offset uint32, length uint32) error {
	var errorCode uint32

	gsys.ReplyPushInput(
		offset,
		length,
		uint32(uintptr(unsafe.Pointer(&errorCode))),
	)

	if errorCode != 0 {
		return ErrReplyPushInputFailed
	}

	return nil
}

func ReplyInputWithGas(gasLimit uint64, value Uint128, offset uint32, length uint32) (MessageID, error) {
	var result [36]byte

	valuePtr := zeroValuePtr

	if value.Lo != 0 || value.Hi != 0 {
		valuePtr = uint32(
			uintptr(unsafe.Pointer(&value)),
		)
	}

	gsys.ReplyInputWithGas(
		offset,
		length,
		gasLimit,
		valuePtr,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrReplyInputWithGasFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func SendInput(destination ActorID, value Uint128, offset uint32, length uint32) (MessageID, error) {
	pidValue := hashWithValue{
		Hash:  destination,
		Value: value,
	}

	var result [36]byte

	gsys.SendInput(
		uint32(uintptr(unsafe.Pointer(&pidValue))),
		offset,
		length,
		0,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrSendInputFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func SendInputDelayed(destination ActorID, value Uint128, offset uint32, length uint32, delay uint32) (MessageID, error) {
	pidValue := hashWithValue{
		Hash:  destination,
		Value: value,
	}

	var result [36]byte

	gsys.SendInput(
		uint32(uintptr(unsafe.Pointer(&pidValue))),
		offset,
		length,
		delay,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrSendInputFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func SendFromReservation(reservationID ReservationID, destination ActorID, payload []byte, value Uint128) (MessageID, error) {
	return SendDelayedFromReservation(
		reservationID,
		destination,
		payload,
		value,
		0,
	)
}

func SendDelayedFromReservation(reservationID ReservationID, destination ActorID, payload []byte, value Uint128, delay uint32) (MessageID, error) {
	ridPidValue := twoHashesWithValue{
		Hash1: ActorID(reservationID),
		Hash2: destination,
		Value: value,
	}

	var result [36]byte

	var payloadPtr uint32
	if len(payload) > 0 {
		payloadPtr = uint32(
			uintptr(unsafe.Pointer(&payload[0])),
		)
	}

	gsys.ReservationSend(
		uint32(uintptr(unsafe.Pointer(&ridPidValue))),
		payloadPtr,
		uint32(len(payload)),
		delay,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrSendFromReservationFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func SendInit() (MessageHandle, error) {
	var result [8]byte

	gsys.SendInit(
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return 0, ErrSendInitFailed
	}

	handle := *(*uint32)(unsafe.Pointer(&result[4]))

	return MessageHandle(handle), nil
}
func SendPushInput(handle MessageHandle, offset uint32, length uint32) error {
	var errorCode uint32

	gsys.SendPushInput(
		uint32(handle),
		offset,
		length,
		uint32(uintptr(unsafe.Pointer(&errorCode))),
	)

	if errorCode != 0 {
		return ErrSendPushInputFailed
	}

	return nil
}
func SendCommit(handle MessageHandle, destination ActorID, value Uint128) (MessageID, error) {
	pidValue := hashWithValue{
		Hash:  destination,
		Value: value,
	}

	var result [36]byte

	gsys.SendCommit(
		uint32(handle),
		uint32(uintptr(unsafe.Pointer(&pidValue))),
		0,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrSendCommitFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}
func SendInputWithGas(destination ActorID, gasLimit uint64, value Uint128, offset, length uint32) (MessageID, error) {
	return SendInputWithGasDelayed(
		destination,
		gasLimit,
		value,
		offset,
		length,
		0,
	)
}
func SendInputWithGasDelayed(destination ActorID, gasLimit uint64, value Uint128, offset, length, delay uint32) (MessageID, error) {
	pidValue := hashWithValue{
		Hash:  destination,
		Value: value,
	}

	var result [36]byte

	gsys.SendInputWithGas(
		uint32(uintptr(unsafe.Pointer(&pidValue))),
		offset,
		length,
		gasLimit,
		delay,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrSendInputWithGasFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}

func SendCommitFromReservation(
	reservationID ReservationID,
	handle MessageHandle,
	destination ActorID,
	value Uint128,
) (MessageID, error) {
	return SendCommitDelayedFromReservation(
		reservationID,
		handle,
		destination,
		value,
		0,
	)
}

func SendCommitDelayedFromReservation(
	reservationID ReservationID,
	handle MessageHandle,
	destination ActorID,
	value Uint128,
	delay uint32,
) (MessageID, error) {
	ridPidValue := twoHashesWithValue{
		Hash1: ActorID(reservationID),
		Hash2: destination,
		Value: value,
	}

	var result [36]byte

	gsys.ReservationSendCommit(
		uint32(handle),
		uint32(uintptr(unsafe.Pointer(&ridPidValue))),
		delay,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return MessageID{}, ErrSendCommitFromReservationFailed
	}

	var id MessageID
	copy(id[:], result[4:36])

	return id, nil
}
