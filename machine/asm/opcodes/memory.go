package opcodes

import (
	"SystemOnePoc/sap3/machine/op"
	"SystemOnePoc/sap3/machine/types"
)

// MemoryReferenceInstructions is the list of OpCode(s) that supply an operand referencing an address in memory.
var MemoryReferenceInstructions = []types.Word{
	op.CALL,
	op.CC,
	op.CM,
	op.CNC,
	op.CNZ,
	op.CP,
	op.CPE,
	op.CPO,
	op.CZ,
	op.JC,
	op.JM,
	op.JMP,
	op.JNC,
	op.JNZ,
	op.JP,
	op.JPE,
	op.JPO,
	op.JZ,
	op.LDA,
	op.STA,
}
