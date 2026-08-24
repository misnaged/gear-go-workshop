package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	reservationID, err := exec.ReserveGas(
		50_000_000,
		100,
	)
	if err != nil {
		ext.Panic("failed to reserve gas")
	}

	_, err = msg.ReplyBytesFromReservation(
		reservationID,
		[]byte("REPLY_FROM_RESERVATION"),
		gcore.Uint128{},
	)
	if err != nil {
		ext.Panic("failed to reply from reservation")
	}
}

func main() {}
