/*
package main

//go:wasmexport handle
func handle() {
	//default entrypoint. for examples see ./examples
}

func main() {}
*/

package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

func messageIDHex(id gcore.MessageID) string {
	const hex = "0123456789abcdef"

	var out [64]byte

	for i, b := range id {
		out[i*2] = hex[b>>4]
		out[i*2+1] = hex[b&0x0f]
	}

	return string(out[:])
}

//go:wasmexport handle
func handle() {
	if err := exec.SystemReserveGas(1_000_000_000); err != nil {
		ext.Panic("failed to reserve system gas")
	}

	ext.Panic("test gear panic")
}

//go:wasmexport handle_signal
func handleSignal() {
	id, err := msg.SignalFrom()
	if err != nil {
		ext.Panic("failed to get signal source")
	}

	ext.Panic("signal_from=0x" + messageIDHex(id))
}

func main() {}
