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
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	handle, err := msg.SendInit()
	if err != nil {
		ext.Panic("send init failed")
	}

	err = msg.SendPushInput(
		handle,
		8,
		11,
	)
	if err != nil {
		ext.Panic("send push input failed")
	}

	_, err = msg.SendCommitDelayed(
		handle,
		msg.Source(),
		gcore.Uint128{},
		5,
	)
	if err != nil {
		ext.Panic("send commit delayed failed")
	}
}

func main() {}
