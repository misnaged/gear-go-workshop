package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	destination := msg.Source()

	_, err := msg.SendBytes(
		destination,
		[]byte("HELLO"),
		gcore.Uint128{},
	)
	if err != nil {
		msg.ReplyBytes([]byte("SEND_ERROR"))
		return
	}

	msg.ReplyBytes([]byte("SENT"))
}

func main() {}
