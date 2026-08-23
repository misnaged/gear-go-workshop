package gcore

import (
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
