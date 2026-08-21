package gcore

import (
	"errors"
	"unsafe"

	"github.com/misnaged/gear-go-workshop/internal/gsys"
)

const zeroValuePtr uint32 = 0xffffffff

type ActorID [32]byte

var (
	ErrBufferTooSmall = errors.New("buffer is too small")
	ErrReadFailed     = errors.New("gr_read failed")
)

func Source() ActorID {
	var source ActorID

	gsys.Source(
		uint32(uintptr(unsafe.Pointer(&source[0]))),
	)

	return source
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

func Size() int {
	var size uint32

	gsys.Size(
		uint32(uintptr(unsafe.Pointer(&size))),
	)

	return int(size)
}
