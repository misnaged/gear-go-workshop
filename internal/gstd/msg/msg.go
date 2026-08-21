package msg

import "github.com/misnaged/gear-go-workshop/internal/gcore"

func ReplyBytes(payload []byte) {
	gcore.Reply(payload)
}
func LoadBytes() ([]byte, error) {
	payload := make([]byte, gcore.Size())

	if err := gcore.Read(payload); err != nil {
		return nil, err
	}

	return payload, nil
}
func SendBytes(destination gcore.ActorID, payload []byte, value gcore.Uint128) (gcore.MessageID, error) {
	return gcore.Send(destination, payload, value)
}
func ReplyTo() (gcore.MessageID, error) {
	return gcore.ReplyTo()
}
func ID() gcore.MessageID {
	return gcore.ID()
}

func Source() gcore.ActorID {
	return gcore.Source()
}
func Value() gcore.Uint128 {
	return gcore.Value()
}
