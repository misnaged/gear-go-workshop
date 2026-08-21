package exec

import "github.com/misnaged/gear-go-workshop/internal/gcore"

func GasAvailable() uint64 {
	return gcore.GasAvailable()
}

func ProgramID() gcore.ActorID {
	return gcore.ProgramID()
}

func Wait() {
	gcore.Wait()
}

func Wake(messageID gcore.MessageID) error {
	return gcore.Wake(messageID)
}
