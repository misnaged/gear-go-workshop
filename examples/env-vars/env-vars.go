package main

import (
	"encoding/binary"

	"github.com/misnaged/gear-go-workshop/internal/gstd/exec"
	"github.com/misnaged/gear-go-workshop/internal/gstd/msg"
)

//go:wasmexport handle
func handle() {
	vars := exec.EnvVars()

	var result [52]byte

	binary.LittleEndian.PutUint32(
		result[0:4],
		vars.PerformanceMultiplier,
	)

	binary.LittleEndian.PutUint64(
		result[4:12],
		vars.ExistentialDeposit.Lo,
	)

	binary.LittleEndian.PutUint64(
		result[12:20],
		vars.ExistentialDeposit.Hi,
	)

	binary.LittleEndian.PutUint64(
		result[20:28],
		vars.MailboxThreshold,
	)

	binary.LittleEndian.PutUint64(
		result[28:36],
		vars.GasMultiplier.GasPerValue,
	)

	binary.LittleEndian.PutUint64(
		result[36:44],
		vars.GasMultiplier.ValuePerGas.Lo,
	)

	binary.LittleEndian.PutUint64(
		result[44:52],
		vars.GasMultiplier.ValuePerGas.Hi,
	)

	msg.ReplyBytes(result[:])
}

func main() {}
