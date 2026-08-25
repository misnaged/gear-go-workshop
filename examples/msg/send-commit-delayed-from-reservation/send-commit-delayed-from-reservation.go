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

	handle, err := msg.SendInit()
	if err != nil {
		ext.Panic("send init failed")
	}

	err = msg.SendPushInput(
		handle,
		8,
		11,
	)
	if err != nil {
		ext.Panic("send push input failed")
	}

	_, err = msg.SendCommitDelayedFromReservation(
		reservationID,
		handle,
		msg.Source(),
		gcore.Uint128{},
		5,
	)
	if err != nil {
		ext.Panic("send commit delayed from reservation failed")
	}
}
func main() {}
