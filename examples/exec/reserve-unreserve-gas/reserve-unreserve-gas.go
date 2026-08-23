package main

import (
	"encoding/binary"

	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	payload, _ := msg.LoadBytes()

	if string(payload) == "RESERVE" {
		id, err := exec.ReserveGas(50_000_000, 100)
		if err != nil {
			msg.ReplyBytes([]byte("RESERVE_ERROR"))
			return
		}

		msg.ReplyBytes(id[:])
		return
	}

	if len(payload) == 32 {
		var id gcore.ReservationID
		copy(id[:], payload)

		gas, err := exec.UnreserveGas(id)
		if err != nil {
			msg.ReplyBytes([]byte("UNRESERVE_ERROR"))
			return
		}

		var result [8]byte
		binary.LittleEndian.PutUint64(result[:], gas)
		msg.ReplyBytes(result[:])
	}
}

func main() {}
