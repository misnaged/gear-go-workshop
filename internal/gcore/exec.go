package gcore

import (
	"encoding/binary"
	"unsafe"

	"github.com/misnaged/gear-go-workshop/internal/gsys"
)

func Wait() {
	gsys.Wait()
}

func WaitFor(duration uint32) {
	gsys.WaitFor(duration)
}

func WaitUpTo(duration uint32) {
	gsys.WaitUpTo(duration)
}

func ProgramID() ActorID {
	var id ActorID

	gsys.ProgramID(
		uint32(uintptr(unsafe.Pointer(&id[0]))),
	)

	return id
}

func GasAvailable() uint64 {
	var gas uint64

	gsys.GasAvailable(
		uint32(uintptr(unsafe.Pointer(&gas))),
	)

	return gas
}

func Wake(messageID MessageID) error {
	return WakeDelayed(messageID, 0)
}

func WakeDelayed(messageID MessageID, delay uint32) error {
	var errorCode uint32

	gsys.Wake(
		uint32(uintptr(unsafe.Pointer(&messageID[0]))),
		delay,
		uint32(uintptr(unsafe.Pointer(&errorCode))),
	)

	if errorCode != 0 {
		return ErrWakeFailed
	}

	return nil
}

func Leave() {
	gsys.Leave()
}

func ValueAvailable() Uint128 {
	var value Uint128

	gsys.ValueAvailable(
		uint32(uintptr(unsafe.Pointer(&value))),
	)

	return value
}

func BlockHeight() uint32 {
	var height uint32

	gsys.BlockHeight(
		uint32(uintptr(unsafe.Pointer(&height))),
	)

	return height
}

func BlockTimestamp() uint64 {
	var timestamp uint64

	gsys.BlockTimestamp(
		uint32(uintptr(unsafe.Pointer(&timestamp))),
	)

	return timestamp
}
func Random(subject [32]byte) ([32]byte, uint32) {
	var result blockNumberWithHash

	gsys.Random(
		uint32(uintptr(unsafe.Pointer(&subject[0]))),
		uint32(uintptr(unsafe.Pointer(&result))),
	)

	return result.Hash, result.BlockNumber
}

func ReserveGas(amount uint64, duration uint32) (ReservationID, error) {
	var result [36]byte

	gsys.ReserveGas(
		amount,
		duration,
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := *(*uint32)(unsafe.Pointer(&result[0]))
	if errorCode != 0 {
		return ReservationID{}, ErrReserveGasFailed
	}

	var id ReservationID
	copy(id[:], result[4:])

	return id, nil
}

func UnreserveGas(id ReservationID) (uint64, error) {
	var result [12]byte

	gsys.UnreserveGas(
		uint32(uintptr(unsafe.Pointer(&id[0]))),
		uint32(uintptr(unsafe.Pointer(&result[0]))),
	)

	errorCode := binary.LittleEndian.Uint32(result[0:4])
	if errorCode != 0 {
		return 0, ErrUnreserveGasFailed
	}

	gas := binary.LittleEndian.Uint64(result[4:12])

	return gas, nil
}

func SystemReserveGas(amount uint64) error {
	var errorCode uint32

	gsys.SystemReserveGas(
		amount,
		uint32(uintptr(unsafe.Pointer(&errorCode))),
	)

	if errorCode != 0 {
		return ErrSystemReserveGasFailed
	}

	return nil
}
