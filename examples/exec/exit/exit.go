package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	source := msg.Source()

	exec.Exit(source)

	//must not be called in the runtime
	msg.ReplyBytes([]byte("AFTER_EXIT"))
}
func main() {}
