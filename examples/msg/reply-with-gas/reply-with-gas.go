package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	_, err := msg.ReplyBytesWithGas(
		[]byte("REPLY_WITH_GAS"),
		50_000_000,
		gcore.Uint128{},
	)
	if err != nil {
		ext.Panic("reply with gas failed")
	}
}

func main() {}
