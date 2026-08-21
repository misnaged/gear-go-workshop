package main

import (
	msg "github.com/misnaged/gear-go-workshop/internal/gstd"
)

/*
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
*/

//go:wasmexport handle
func handle() {
	source := msg.Source()

	msg.ReplyBytes(source[:])
}
func main() {}
