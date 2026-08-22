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
