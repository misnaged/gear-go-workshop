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

func WaitFor(duration uint32) {
	gcore.WaitFor(duration)
}

func WaitUpTo(duration uint32) {
	gcore.WaitUpTo(duration)
}

func Wake(messageID gcore.MessageID) error {
	return gcore.Wake(messageID)
}

func WakeDelayed(messageID gcore.MessageID, delay uint32) error {
	return gcore.WakeDelayed(messageID, delay)
}
