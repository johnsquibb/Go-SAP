package op

import "Go-SAP3/machine/types"

const (
	DCX_B  = 0x0b
	DCX_D  = 0x1b
	DCX_H  = 0x2b
	DCX_SP = 0x3b
)

var mcDCX_B = []types.OctupleWord{
	BREG_accen | CREG_accen | ALU_we | ALU_dcx, // T4
	Noop, // T5
}

var mcDCX_D = []types.OctupleWord{
	DREG_accen | EREG_accen | ALU_we | ALU_dcx, // T4
	Noop, // T5
}

var mcDCX_H = []types.OctupleWord{
	HREG_accen | LREG_accen | ALU_we | ALU_dcx, // T4
	Noop, // T5
}

var mcDCX_SP = []types.OctupleWord{
	SP_accen | ALU_we | ALU_dcx, // T4
	Noop,                        // T5
}
