package op

import "SystemOnePoc/sap3/machine/types"

const (
	POP_B   = 0xc1
	POP_D   = 0xd1
	POP_H   = 0xe1
	POP_PSW = 0xf1
)

var mcPOP_B = []types.OctupleWord{
	SP_we | MAR_re,     // T4
	CREG_re | MDR_we,   // T5
	SP_accen | ALU_inr, // T6
	SP_we | MAR_re,     // T7
	BREG_re | MDR_we,   // T8
	SP_accen | ALU_inr, // T9
	Noop,               // T10
}

var mcPOP_D = []types.OctupleWord{
	SP_we | MAR_re,     // T4
	EREG_re | MDR_we,   // T5
	SP_accen | ALU_inr, // T6
	SP_we | MAR_re,     // T7
	DREG_re | MDR_we,   // T8
	SP_accen | ALU_inr, // T9
	Noop,               // T10
}

var mcPOP_H = []types.OctupleWord{
	SP_we | MAR_re,     // T4
	LREG_re | MDR_we,   // T5
	SP_accen | ALU_inr, // T6
	SP_we | MAR_re,     // T7
	HREG_re | MDR_we,   // T8
	SP_accen | ALU_inr, // T9
	Noop,               // T10
}

var mcPOP_PSW = []types.OctupleWord{
	SP_we | MAR_re,     // T4
	FREG_re | MDR_we,   // T5
	SP_accen | ALU_inr, // T6
	SP_we | MAR_re,     // T7
	ACC_re | MDR_we,    // T8
	SP_accen | ALU_inr, // T9
	Noop,               // T10
}
