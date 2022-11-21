package op

import "Go-SAP/machine/types"

const (
	STC = 0x37
	CMC = 0x3f
)

var mcSTC = []types.OctupleWord{
	ALU_we | ALU_stc, // T4
	Noop,             // T5
}

var mcCMC = []types.OctupleWord{
	ALU_we | ALU_cmc, // T4
	Noop,             // T5
}
