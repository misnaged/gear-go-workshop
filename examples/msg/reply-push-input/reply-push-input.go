package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	if err := msg.ReplyPush([]byte("copied: ")); err != nil {
		ext.Panic("reply push failed")
	}

	if err := msg.ReplyPushInput(8, 11); err != nil {
		ext.Panic("reply push input failed")
	}

	_, err := msg.ReplyCommit(gcore.Uint128{})
	if err != nil {
		ext.Panic("reply commit failed")
	}
}

func main() {}
