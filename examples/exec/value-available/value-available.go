package main

import (
	"encoding/binary"

	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	value := exec.ValueAvailable()

	var result [16]byte

	binary.LittleEndian.PutUint64(result[0:8], value.Lo)
	binary.LittleEndian.PutUint64(result[8:16], value.Hi)

	msg.ReplyBytes(result[:])
}

func main() {}
