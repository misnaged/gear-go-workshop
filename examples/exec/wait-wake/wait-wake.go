package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

var awakened bool

//go:wasmexport handle
func handle() {
	payload, err := msg.LoadBytes()
	if err != nil {
		return
	}

	if string(payload) == "WAIT" {
		if awakened {
			msg.ReplyBytes([]byte("AFTER_WAKE"))
			return
		}

		exec.Wait()
	}

	if len(payload) == 32 {
		var id gcore.MessageID
		copy(id[:], payload)

		awakened = true

		if err := exec.Wake(id); err != nil {
			msg.ReplyBytes([]byte("WAKE_ERROR"))
			return
		}

		msg.ReplyBytes([]byte("WOKEN"))
	}
}

func main() {}
