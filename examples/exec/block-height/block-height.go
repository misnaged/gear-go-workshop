package main

import (
	"encoding/binary"

	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	height := exec.BlockHeight()

	var result [4]byte
	binary.LittleEndian.PutUint32(result[:], height)

	msg.ReplyBytes(result[:])
}

func main() {}
