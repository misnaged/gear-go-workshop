package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	if err := msg.ReplyPush([]byte("Hello, ")); err != nil {
		ext.Panic("first reply push failed")
	}

	if err := msg.ReplyPush([]byte("world!")); err != nil {
		ext.Panic("second reply push failed")
	}

	_, err := msg.ReplyCommit(gcore.Uint128{})
	if err != nil {
		ext.Panic("reply commit failed")
	}
}
func main() {}
