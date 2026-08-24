package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
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
	code, err := msg.SignalCode()
	if err != nil {
		ext.Panic("failed to get signal code")
	}

	if code != gcore.SignalUserspacePanic {
		ext.Panic("unexpected signal code")
	}

}
func main() {}
