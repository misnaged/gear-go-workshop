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

	h, err := msg.SendInit()
	if err != nil {
		ext.Panic("send init failed")
	}

	err = msg.SendPushInput(
		h,
		8,
		11,
	)
	if err != nil {
		ext.Panic("send push input failed")
	}

	_, err = msg.SendCommitFromReservation(
		reservationID,
		h,
		msg.Source(),
		gcore.Uint128{},
	)
	if err != nil {
		ext.Panic("send commit from reservation failed")
	}
}
func main() {}
