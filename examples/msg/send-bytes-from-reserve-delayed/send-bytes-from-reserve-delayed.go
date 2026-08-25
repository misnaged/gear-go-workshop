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
		ext.Panic("reserve gas failed")
	}

	_, err = msg.SendBytesDelayedFromReservation(
		reservationID,
		msg.Source(),
		[]byte("DELAYED_FROM_RESERVATION"),
		gcore.Uint128{},
		5,
	)
	if err != nil {
		ext.Panic("delayed send from reservation failed")
	}
}

func main() {}
