package opcodes

import (
	"Go-SAP3/machine/op"
	"Go-SAP3/machine/types"
)

// ExtendedInstructions is the list of OpCode(s) that supply an operand as two bytes (MSB|LSB) for immediate use.
var ExtendedInstructions = []types.Word{
	op.LXI_B,
	op.LXI_D,
	op.LXI_H,
	op.LXI_SP,
}
