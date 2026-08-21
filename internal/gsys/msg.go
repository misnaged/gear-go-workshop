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
