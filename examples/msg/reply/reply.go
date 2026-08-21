package main

import "github.com/misnaged/gear-go-workshop/internal/gstd/msg"

//go:wasmexport handle
func handle() {

	b, err := msg.LoadBytes()
	if err != nil {
		return
	}

	var reply []byte

	switch string(b) {
	case "beach":
		reply = []byte("sea")

	case "sand":
		reply = []byte("desert")

	default:
		reply = []byte("unknown")
	}

	msg.ReplyBytes(reply)
}
