package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

var waited bool

//go:wasmexport handle
func handle() {
	payload, err := msg.LoadBytes()
	if err != nil {
		return
	}

	if string(payload) != "WAIT_UP_TO" {
		return
	}

	if waited {
		msg.ReplyBytes([]byte("AFTER_WAIT_UP_TO"))
		return
	}

	waited = true

	exec.WaitUpTo(3)
}

func main() {}
