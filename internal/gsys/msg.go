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

//go:wasmimport env gr_size
func Size(lengthPtr uint32)

//go:wasmimport env gr_source
func Source(resultPtr uint32)
