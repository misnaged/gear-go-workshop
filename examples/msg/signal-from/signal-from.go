package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	if err := exec.SystemReserveGas(1_000_000_000); err != nil {
		ext.Panic("failed to reserve system gas")
	}

	ext.Panic("test gear panic")
}

//go:wasmexport handle_signal
func handleSignal() {
	if _, err := msg.SignalFrom(); err != nil {
		ext.Panic("failed to get signal source")
	}
}

func main() {}
