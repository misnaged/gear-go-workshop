package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	_, err := msg.SendBytesWithGasDelayed(
		msg.Source(),
		[]byte("SEND_WITH_GAS_DELAYED"),
		50_000_000,
		gcore.Uint128{},
		5,
	)
	if err != nil {
		ext.Panic("send with gas delayed failed")
	}
}

func main() {}
