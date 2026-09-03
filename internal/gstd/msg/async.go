package msg

import (
	"github.com/misnaged/gear-go-workshop/internal/gcore"
	"github.com/misnaged/gear-go-workshop/internal/gstd/async_runtime"
	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
)

type MessageFuture struct {
	WaitingReplyTo gcore.MessageID
	ReplyDeposit   uint64
}

func SendBytesForReply(destination gcore.ActorID, payload []byte, value gcore.Uint128, replyDeposit uint64) (*MessageFuture, error) {
	waitingReplyTo, err := SendBytes(
		destination,
		payload,
		value,
	)
	if err != nil {
		return nil, err
	}

	if replyDeposit != 0 {
		if err = exec.ReplyDeposit(
			waitingReplyTo,
			replyDeposit,
		); err != nil {
			return nil, err
		}
	}

	owner := ID()

	if err := async_runtime.Register(
		waitingReplyTo,
		owner,
	); err != nil {
		return nil, err
	}

	return &MessageFuture{
		WaitingReplyTo: waitingReplyTo,
		ReplyDeposit:   replyDeposit,
	}, nil
}
func SendBytesWithGasForReply(
	destination gcore.ActorID,
	payload []byte,
	gasLimit uint64,
	value gcore.Uint128,
	replyDeposit uint64,
) (*MessageFuture, error) {
	waitingReplyTo, err := SendBytesWithGas(
		destination,
		payload,
		gasLimit,
		value,
	)
	if err != nil {
		return nil, err
	}

	if replyDeposit != 0 {
		if err := exec.ReplyDeposit(
			waitingReplyTo,
			replyDeposit,
		); err != nil {
			return nil, err
		}
	}

	owner := ID()

	if err := async_runtime.Register(
		waitingReplyTo,
		owner,
	); err != nil {
		return nil, err
	}

	return &MessageFuture{
		WaitingReplyTo: waitingReplyTo,
		ReplyDeposit:   replyDeposit,
	}, nil
}
