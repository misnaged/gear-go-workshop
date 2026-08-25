package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	_, err := msg.SendInputWithGasDelayed(
		msg.Source(),
		50_000_000,
		gcore.Uint128{},
		8,
		11,
		5,
	)
	if err != nil {
		ext.Panic("send input with gas delayed failed")
	}
}
func main() {}
