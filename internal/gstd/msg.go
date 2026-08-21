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

func Source() gcore.ActorID {
	return gcore.Source()
}
