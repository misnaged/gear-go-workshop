package gcore

import (
	"unsafe"

	"github.com/misnaged/gear-go-workshop/internal/gsys"
)

func Panic(data []byte) {
	var dataPtr uint32

	if len(data) > 0 {
		dataPtr = uint32(
			uintptr(unsafe.Pointer(&data[0])),
		)
	}

	gsys.Panic(
		dataPtr,
		uint32(len(data)),
	)
}

func PanicString(data string) {
	var dataPtr uint32

	if len(data) > 0 {
		dataPtr = uint32(
			uintptr(unsafe.Pointer(unsafe.StringData(data))),
		)
	}

	gsys.Panic(
		dataPtr,
		uint32(len(data)),
	)
}
