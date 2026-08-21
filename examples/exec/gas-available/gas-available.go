package main

import (
	"unsafe"

	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	gas := exec.GasAvailable()

	var result [8]byte
	*(*uint64)(unsafe.Pointer(&result[0])) = gas

	msg.ReplyBytes(result[:])
}
