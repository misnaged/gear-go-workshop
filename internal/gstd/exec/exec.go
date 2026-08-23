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

func Leave() {
	gcore.Leave()
}

func ValueAvailable() gcore.Uint128 {
	return gcore.ValueAvailable()
}

func BlockHeight() uint32 {
	return gcore.BlockHeight()
}

func BlockTimestamp() uint64 {
	return gcore.BlockTimestamp()
}
func Random(subject [32]byte) ([32]byte, uint32) {
	return gcore.Random(subject)
}

func ReserveGas(amount uint64, duration uint32) (gcore.ReservationID, error) {
	return gcore.ReserveGas(amount, duration)
}

func UnreserveGas(id gcore.ReservationID) (uint64, error) {
	return gcore.UnreserveGas(id)
}

func SystemReserveGas(amount uint64) error {
	return gcore.SystemReserveGas(amount)
}

func Exit(inheritorID gcore.ActorID) {
	gcore.Exit(inheritorID)
}
