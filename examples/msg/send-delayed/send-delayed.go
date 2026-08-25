package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	_, err := msg.SendBytesDelayed(
		msg.Source(),
		[]byte("SEND_DELAYED"),
		gcore.Uint128{},
		5,
	)
	if err != nil {
		ext.Panic("send delayed failed")
	}
}
