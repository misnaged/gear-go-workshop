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

	if string(payload) != "WAIT_FOR" {
		return
	}

	if waited {
		msg.ReplyBytes([]byte("AFTER_WAIT_FOR"))
		return
	}

	waited = true

	exec.WaitFor(3)
}

func main() {}
