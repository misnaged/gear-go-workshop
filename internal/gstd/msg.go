package msg

import "github.com/misnaged/gear-go-workshop/internal/gcore"

func ReplyBytes(payload []byte) {
	gcore.Reply(payload)
}
