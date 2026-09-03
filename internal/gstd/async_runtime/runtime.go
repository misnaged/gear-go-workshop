package async_runtime

import (
	"errors"

	"github.com/misnaged/gear-go-workshop/internal/gcore"
)

var (
	ErrReplyAlreadyRegistered = errors.New("reply already registered")
	ErrOwnerAlreadyWaiting    = errors.New("owner already waiting")
)
var waitingByOwner = make(map[gcore.MessageID]gcore.MessageID)

type ReplyState struct {
	Owner   gcore.MessageID
	Payload []byte
	Code    gcore.ReplyCodeBytes
	Ready   bool
}

var replies = make(map[gcore.MessageID]*ReplyState)

func Register(waitingReplyTo gcore.MessageID, owner gcore.MessageID) error {
	if _, exists := replies[waitingReplyTo]; exists {
		return ErrReplyAlreadyRegistered
	}

	if _, exists := waitingByOwner[owner]; exists {
		return ErrOwnerAlreadyWaiting
	}

	replies[waitingReplyTo] = &ReplyState{
		Owner: owner,
	}

	waitingByOwner[owner] = waitingReplyTo

	return nil
}

func Get(waitingReplyTo gcore.MessageID) (*ReplyState, bool) {
	state, ok := replies[waitingReplyTo]
	return state, ok
}

func Remove(waitingReplyTo gcore.MessageID) {
	state, ok := replies[waitingReplyTo]
	if ok {
		delete(waitingByOwner, state.Owner)
	}

	delete(replies, waitingReplyTo)
}

func RecordReply(waitingReplyTo gcore.MessageID, payload []byte, code gcore.ReplyCodeBytes) (gcore.MessageID, bool) {
	state, ok := replies[waitingReplyTo]
	if !ok {
		return gcore.MessageID{}, false
	}

	state.Payload = append([]byte(nil), payload...)
	state.Code = code
	state.Ready = true

	return state.Owner, true
}
func WaitingReply(owner gcore.MessageID) (gcore.MessageID, bool) {
	waitingReplyTo, ok := waitingByOwner[owner]
	return waitingReplyTo, ok
}
