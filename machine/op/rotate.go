package op

import "Go-SAP/machine/types"

const (
	RAL = 0x17
	RAR = 0x1f
	RLC = 0x07
	RRC = 0x0f
)

var mcRAL = []types.OctupleWord{
	ALU_ral | ALU_we | FLG_cen | ACC_re, // T4
	Noop,                                // T5
}

var mcRAR = []types.OctupleWord{
	ALU_rar | ALU_we | FLG_cen | ACC_re, // T4
	Noop,                                // T5
}

var mcRLC = []types.OctupleWord{
	ALU_rlc | ALU_we | FLG_cen | ACC_re, // T4
	Noop,                                // T5
}

var mcRRC = []types.OctupleWord{
	ALU_rrc | ALU_we | FLG_cen | ACC_re, // T4
	Noop,                                // T5
}
