package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	_, err := msg.ReplyInput(
		gcore.Uint128{},
		8,
		11,
	)
	if err != nil {
		ext.Panic("reply input failed")
	}
}

func main() {}
