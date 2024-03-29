package opcodes

import (
	"Go-SAP/machine/op"
	"Go-SAP/machine/types"
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
	op.LHLD,
	op.SHLD,
	op.STA,
}
