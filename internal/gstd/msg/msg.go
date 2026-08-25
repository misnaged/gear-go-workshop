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

func SendBytesDelayed(destination gcore.ActorID, payload []byte, value gcore.Uint128, delay uint32) (gcore.MessageID, error) {
	return gcore.SendDelayed(
		destination,
		payload,
		value,
		delay,
	)
}

func Size() int {
	return gcore.Size()
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

func ReplyCode() (gcore.ReplyCodeBytes, error) {
	return gcore.ReplyCode()
}

func SignalCode() (gcore.SignalCodeValue, error) {
	return gcore.SignalCode()
}

func ReplyBytesFromReservation(reservationID gcore.ReservationID, payload []byte, value gcore.Uint128) (gcore.MessageID, error) {
	return gcore.ReplyFromReservation(
		reservationID,
		payload,
		value,
	)
}
func ReplyBytesWithGas(payload []byte, gasLimit uint64, value gcore.Uint128) (gcore.MessageID, error) {
	return gcore.ReplyWithGas(
		payload,
		gasLimit,
		value,
	)
}
func ReplyCommit(value gcore.Uint128) (gcore.MessageID, error) {
	return gcore.ReplyCommit(value)
}
func ReplyPush(payload []byte) error {
	return gcore.ReplyPush(payload)
}

func ReplyCommitWithGas(gasLimit uint64, value gcore.Uint128) (gcore.MessageID, error) {
	return gcore.ReplyCommitWithGas(
		gasLimit,
		value,
	)
}

func ReplyCommitFromReservation(reservationID gcore.ReservationID, value gcore.Uint128) (gcore.MessageID, error) {
	return gcore.ReplyCommitFromReservation(
		reservationID,
		value,
	)
}

func ReplyInput(value gcore.Uint128, offset, length uint32) (gcore.MessageID, error) {
	return gcore.ReplyInput(
		value,
		offset,
		length,
	)
}

func ReplyPushInput(offset, length uint32) error {
	return gcore.ReplyPushInput(
		offset,
		length,
	)
}

func ReplyInputWithGas(gasLimit uint64, value gcore.Uint128, offset, length uint32) (gcore.MessageID, error) {
	return gcore.ReplyInputWithGas(
		gasLimit,
		value,
		offset,
		length,
	)
}
func SendInput(destination gcore.ActorID, value gcore.Uint128, offset, length uint32) (gcore.MessageID, error) {
	return gcore.SendInput(
		destination,
		value,
		offset,
		length,
	)
}
func SendInputDelayed(destination gcore.ActorID, value gcore.Uint128, offset, length, delay uint32) (gcore.MessageID, error) {
	return gcore.SendInputDelayed(
		destination,
		value,
		offset,
		length,
		delay,
	)
}
func SendBytesFromReservation(reservationID gcore.ReservationID, destination gcore.ActorID, payload []byte, value gcore.Uint128) (gcore.MessageID, error) {
	return gcore.SendFromReservation(
		reservationID,
		destination,
		payload,
		value,
	)
}

func SendBytesDelayedFromReservation(
	reservationID gcore.ReservationID,
	destination gcore.ActorID, payload []byte,
	value gcore.Uint128, delay uint32) (gcore.MessageID, error) {
	return gcore.SendDelayedFromReservation(
		reservationID,
		destination,
		payload,
		value,
		delay,
	)
}

func SendInit() (gcore.MessageHandle, error) {
	return gcore.SendInit()
}

func SendPushInput(handle gcore.MessageHandle, offset, length uint32) error {
	return gcore.SendPushInput(handle, offset, length)
}

func SendCommit(handle gcore.MessageHandle, destination gcore.ActorID, value gcore.Uint128) (gcore.MessageID, error) {
	return gcore.SendCommit(handle, destination, value)
}
func SendInputWithGas(destination gcore.ActorID, gasLimit uint64, value gcore.Uint128, offset, length uint32) (gcore.MessageID, error) {
	return gcore.SendInputWithGas(
		destination,
		gasLimit,
		value,
		offset,
		length,
	)
}
func SendInputWithGasDelayed(destination gcore.ActorID, gasLimit uint64, value gcore.Uint128, offset, length, delay uint32) (gcore.MessageID, error) {
	return gcore.SendInputWithGasDelayed(
		destination,
		gasLimit,
		value,
		offset,
		length,
		delay,
	)
}
func SendCommitFromReservation(reservationID gcore.ReservationID, handle gcore.MessageHandle, destination gcore.ActorID, value gcore.Uint128) (gcore.MessageID, error) {
	return gcore.SendCommitFromReservation(
		reservationID,
		handle,
		destination,
		value,
	)
}

func SendCommitDelayedFromReservation(reservationID gcore.ReservationID, handle gcore.MessageHandle,
	destination gcore.ActorID, value gcore.Uint128, delay uint32) (gcore.MessageID, error) {
	return gcore.SendCommitDelayedFromReservation(
		reservationID,
		handle,
		destination,
		value,
		delay,
	)
}
