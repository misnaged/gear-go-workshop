package main

import (
	"encoding/binary"

	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	timestamp := exec.BlockTimestamp()

	var result [8]byte
	binary.LittleEndian.PutUint64(result[:], timestamp)

	msg.ReplyBytes(result[:])
}

func main() {}
