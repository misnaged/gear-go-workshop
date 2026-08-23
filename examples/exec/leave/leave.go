package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	payload, err := msg.LoadBytes()
	if err != nil {
		return
	}

	if string(payload) == "LEAVE" {
		exec.Leave()

		//must not be called in the runtime
		msg.ReplyBytes([]byte("AFTER_LEAVE"))
	}
}

func main() {}
