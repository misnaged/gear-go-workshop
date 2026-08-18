package gcore

import (
	"unsafe"

	"github.com/misnaged/gear-go-workshop/internal/gsys"
)

const zeroValuePtr uint32 = 0xffffffff

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
