package gcore

import (
	"unsafe"

	"github.com/misnaged/gear-go-workshop/internal/gsys"
)

func Wait() {
	gsys.Wait()
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
	var errorCode uint32

	gsys.Wake(
		uint32(uintptr(unsafe.Pointer(&messageID[0]))),
		0,
		uint32(uintptr(unsafe.Pointer(&errorCode))),
	)

	if errorCode != 0 {
		return ErrWakeFailed
	}

	return nil
}
