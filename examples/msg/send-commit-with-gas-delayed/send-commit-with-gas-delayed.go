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

	if err := msg.SendPush(handle, []byte("HELLO")); err != nil {
		ext.Panic("send push failed")
	}

	_, err = msg.SendCommitWithGasDelayed(
		handle,
		msg.Source(),
		1_000_000,
		gcore.Uint128{},
		5,
	)
	if err != nil {
		ext.Panic("send commit with gas delayed failed")
	}
}

func main() {}
