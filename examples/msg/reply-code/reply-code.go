package main

import (
	"bytes"

	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	payload, err := msg.LoadBytes()
	if err != nil {
		ext.Panic("failed to load payload")
	}

	if bytes.Equal(payload, []byte("START")) {
		_, err := msg.SendBytes(
			exec.ProgramID(),
			[]byte("PING"),
			gcore.Uint128{},
		)
		if err != nil {
			ext.Panic("failed to send PING")
		}

		return
	}

	if bytes.Equal(payload, []byte("PING")) {
		msg.ReplyBytes([]byte("PONG"))
		return
	}

	ext.Panic("unexpected payload")
}

//go:wasmexport handle_reply
func handleReply() {
	code, err := msg.ReplyCode()
	if err != nil {
		ext.Panic("failed to get reply code")
	}

	// Success(Manual) == 00 01 00 00
	if code[0] != 0 ||
		code[1] != 1 ||
		code[2] != 0 ||
		code[3] != 0 {
		ext.Panic("unexpected reply code")
	}
}
func main() {}
