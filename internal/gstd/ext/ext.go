package ext

import "github.com/misnaged/gear-go-workshop/internal/gcore"

func Panic(message string) {
	gcore.PanicString(message)
}

func PanicBytes(data []byte) {
	gcore.Panic(data)
}
