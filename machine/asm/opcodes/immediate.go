package opcodes

import (
	"SystemOnePoc/sap3/machine/op"
	"SystemOnePoc/sap3/machine/types"
)

// ImmediateInstructions is the list of OpCode(s) that supply an operand as a byte for immediate use.
var ImmediateInstructions = []types.Word{
	op.ACI,
	op.ADI,
	op.ANI,
	op.CPI,
	op.IN,
	op.MVI_A,
	op.MVI_B,
	op.MVI_C,
	op.MVI_D,
	op.MVI_E,
	op.MVI_H,
	op.MVI_L,
	op.MVI_M,
	op.ORI,
	op.OUT,
	op.SBI,
	op.SUI,
	op.XRI,
}
