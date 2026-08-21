package gsys

//go:wasmimport env gr_program_id
func ProgramID(resultPtr uint32)

//go:wasmimport env gr_gas_available
func GasAvailable(resultPtr uint32)

//go:wasmimport env gr_wait
func Wait()

//go:wasmimport env gr_wake
func Wake(messageIDPtr uint32, delay uint32, errorPtr uint32)
