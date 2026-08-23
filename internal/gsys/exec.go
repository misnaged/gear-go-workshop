package gsys

//go:wasmimport env gr_program_id
func ProgramID(resultPtr uint32)

//go:wasmimport env gr_gas_available
func GasAvailable(resultPtr uint32)

//go:wasmimport env gr_wait
func Wait()

//go:wasmimport env gr_wait_for
func WaitFor(duration uint32)

//go:wasmimport env gr_wait_up_to
func WaitUpTo(duration uint32)

//go:wasmimport env gr_wake
func Wake(messageIDPtr uint32, delay uint32, errorPtr uint32)

//go:wasmimport env gr_leave
func Leave()

//go:wasmimport env gr_value_available
func ValueAvailable(resultPtr uint32)

//go:wasmimport env gr_block_height
func BlockHeight(resultPtr uint32)

//go:wasmimport env gr_block_timestamp
func BlockTimestamp(resultPtr uint32)

//go:wasmimport env gr_random
func Random(subjectPtr uint32, resultPtr uint32)

//go:wasmimport env gr_reserve_gas
func ReserveGas(amount uint64, duration uint32, resultPtr uint32)

//go:wasmimport env gr_unreserve_gas
func UnreserveGas(idPtr uint32, resultPtr uint32)

//go:wasmimport env gr_system_reserve_gas
func SystemReserveGas(amount uint64, errorPtr uint32)

//go:wasmimport env gr_exit
func Exit(inheritorIDPtr uint32)

//go:wasmimport env gr_env_vars
func EnvVars(version uint32, varsPtr uint32)
