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

	if err := msg.ReplyPush([]byte("Hello, ")); err != nil {
		ext.Panic("first reply push failed")
	}

	if err := msg.ReplyPush([]byte("reservation!")); err != nil {
		ext.Panic("second reply push failed")
	}

	_, err = msg.ReplyCommitFromReservation(
		reservationID,
		gcore.Uint128{},
	)
	if err != nil {
		ext.Panic("reply commit from reservation failed")
	}
}

func main() {}
