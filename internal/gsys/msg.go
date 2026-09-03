package gsys

//go:wasmimport env gr_reply
func Reply(
	payloadPtr uint32,
	payloadLen uint32,
	valuePtr uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_read
func Read(
	offset uint32,
	length uint32,
	bufferPtr uint32,
	errorPtr uint32,
)

//go:wasmimport env gr_send
func Send(
	pidValuePtr uint32,
	payloadPtr uint32,
	payloadLen uint32,
	delay uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_reply_deposit
func ReplyDeposit(
	messageIDPtr uint32,
	gas uint64,
	errorPtr uint32,
)

//go:wasmimport env gr_size
func Size(lengthPtr uint32)

//go:wasmimport env gr_source
func Source(resultPtr uint32)

//go:wasmimport env gr_message_id
func MessageID(resultPtr uint32)

//go:wasmimport env gr_value
func Value(resultPtr uint32)

//go:wasmimport env gr_reply_to
func ReplyTo(resultPtr uint32)

//go:wasmimport env gr_reply_code
func ReplyCode(resultPtr uint32)

//go:wasmimport env gr_reply_wgas
func ReplyWithGas(
	payloadPtr uint32,
	payloadLen uint32,
	gasLimit uint64,
	valuePtr uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_reply_commit
func ReplyCommit(valuePtr uint32, resultPtr uint32)

//go:wasmimport env gr_signal_code
func SignalCode(resultPtr uint32)

//go:wasmimport env gr_signal_from
func SignalFrom(resultPtr uint32)

//go:wasmimport env gr_reservation_reply
func ReservationReply(
	reservationValuePtr uint32,
	payloadPtr uint32,
	payloadLen uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_reply_input
func ReplyInput(
	offset uint32,
	length uint32,
	valuePtr uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_reply_push_input
func ReplyPushInput(
	offset uint32,
	length uint32,
	errorPtr uint32,
)

//go:wasmimport env gr_reservation_reply_commit
func ReservationReplyCommit(
	reservationValuePtr uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_reservation_send_commit
func ReservationSendCommit(
	handle uint32,
	reservationDestinationValuePtr uint32,
	delay uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_reply_commit_wgas
func ReplyCommitWithGas(
	gasLimit uint64,
	valuePtr uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_reply_push
func ReplyPush(
	payloadPtr uint32,
	payloadLen uint32,
	errorPtr uint32,
)

//go:wasmimport env gr_reply_input_wgas
func ReplyInputWithGas(
	offset uint32,
	length uint32,
	gasLimit uint64,
	valuePtr uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_send_input
func SendInput(
	destinationValuePtr uint32,
	offset uint32,
	length uint32,
	delay uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_reservation_send
func ReservationSend(
	reservationDestinationValuePtr uint32,
	payloadPtr uint32,
	payloadLen uint32,
	delay uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_send_init
func SendInit(resultPtr uint32)

//go:wasmimport env gr_send_push_input
func SendPushInput(
	handle uint32,
	offset uint32,
	length uint32,
	errorPtr uint32,
)

//go:wasmimport env gr_send_push
func SendPush(
	handle uint32,
	payloadPtr uint32,
	payloadLen uint32,
	errorPtr uint32,
)

//go:wasmimport env gr_send_commit
func SendCommit(
	handle uint32,
	destinationValuePtr uint32,
	delay uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_send_input_wgas
func SendInputWithGas(
	destinationValuePtr uint32,
	offset uint32,
	length uint32,
	gasLimit uint64,
	delay uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_send_wgas
func SendWithGas(
	destinationValuePtr uint32,
	payloadPtr uint32,
	payloadLen uint32,
	gasLimit uint64,
	delay uint32,
	resultPtr uint32,
)

//go:wasmimport env gr_send_commit_wgas
func SendCommitWithGas(
	handle uint32,
	destinationValuePtr uint32,
	gasLimit uint64,
	delay uint32,
	resultPtr uint32,
)
