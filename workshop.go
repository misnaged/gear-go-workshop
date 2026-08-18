package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gstd"
)

//go:wasmexport handle
func handle() {
	msg.ReplyBytes([]byte("PONG"))
}
