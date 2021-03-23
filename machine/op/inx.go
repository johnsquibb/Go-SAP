package op

import "Go-SAP3/machine/types"

const (
	INX_B  = 0x03
	INX_D  = 0x13
	INX_H  = 0x23
	INX_SP = 0x33
)

var mcINX_B = []types.OctupleWord{
	BREG_accen | CREG_accen | ALU_we | ALU_inx, // T4
	Noop, // T5
}

var mcINX_D = []types.OctupleWord{
	DREG_accen | EREG_accen | ALU_we | ALU_inx, // T4
	Noop, // T5
}

var mcINX_H = []types.OctupleWord{
	HREG_accen | LREG_accen | ALU_we | ALU_inx, // T4
	Noop, // T5
}
