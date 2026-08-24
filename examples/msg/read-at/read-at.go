package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	size := msg.Size()

	if size < 8 {
		ext.Panic("payload too small")
	}

	payload := make([]byte, size-8)

	if err := gcore.ReadAt(8, payload); err != nil {
		ext.Panic("failed to read payload")
	}

	msg.ReplyBytes(payload)
}
func main() {}
