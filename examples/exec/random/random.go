package main

import (
	"encoding/binary"

	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	var subject1 [32]byte
	var subject2 [32]byte

	for i := range subject1 {
		subject1[i] = byte(i + 1)
		subject2[i] = byte(i + 2)
	}

	seed1, block1 := exec.Random(subject1)
	seed2, block2 := exec.Random(subject2)

	var result [72]byte

	copy(result[0:32], seed1[:])
	binary.LittleEndian.PutUint32(result[32:36], block1)

	copy(result[36:68], seed2[:])
	binary.LittleEndian.PutUint32(result[68:72], block2)

	msg.ReplyBytes(result[:])
}

func main() {}
