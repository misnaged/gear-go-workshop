package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	err := exec.SystemReserveGas(1_000_000)
	if err != nil {
		msg.ReplyBytes([]byte("SYSTEM_RESERVE_ERROR"))
		return
	}

	msg.ReplyBytes([]byte("SYSTEM_RESERVE_OK"))
}

func main() {}
