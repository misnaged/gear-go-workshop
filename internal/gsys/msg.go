package gsys

//go:wasmimport env gr_reply
func Reply(
	payloadPtr uint32,
	payloadLen uint32,
	valuePtr uint32,
	resultPtr uint32,
)
