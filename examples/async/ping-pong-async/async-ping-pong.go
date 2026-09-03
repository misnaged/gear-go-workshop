package main

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/async_runtime"
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/ext"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

var destination gcore.ActorID

//go:wasmexport init
func initProgram() {
	payload, err := msg.LoadBytes()
	if err != nil {
		ext.Panic("failed to load init payload")
	}

	if len(payload) != 32 {
		ext.Panic("invalid destination program id")
	}

	copy(destination[:], payload)
}

//go:wasmexport handle
func handle() {
	owner := msg.ID()

	waitingReplyTo, waiting := async_runtime.WaitingReply(owner)

	if waiting {
		state, ok := async_runtime.Get(waitingReplyTo)
		if !ok {
			ext.Panic("async state not found")
		}

		if !state.Ready {
			exec.Wait()
			return
		}

		payload := state.Payload
		code := state.Code

		async_runtime.Remove(waitingReplyTo)

		if !code.IsSuccess() {
			ext.Panic("reply returned error")
		}

		msg.ReplyBytes(payload)
		return
	}

	_, err := msg.SendBytesWithGasForReply(
		destination,
		[]byte("PING"),
		1_500_000_000,
		gcore.Uint128{},
		1_500_000_000,
	)

	if err != nil {
		ext.Panic("failed to send message")
	}

	exec.Wait()
	return
}

//go:wasmexport handle_reply
func handleReply() {
	waitingReplyTo, err := msg.ReplyTo()
	if err != nil {
		ext.Panic("failed to get reply_to")
	}

	payload, err := msg.LoadBytes()
	if err != nil {
		ext.Panic("failed to load reply")
	}

	code, err := msg.ReplyCode()
	if err != nil {
		ext.Panic("failed to get reply code")
	}

	owner, ok := async_runtime.RecordReply(
		waitingReplyTo,
		payload,
		code,
	)
	if !ok {
		return
	}

	if err := exec.Wake(owner); err != nil {
		ext.Panic("failed to wake owner")
	}
}
func main() {}
