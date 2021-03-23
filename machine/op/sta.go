package op

import "Go-SAP3/machine/types"

const (
	STA = 0x32
)

var mcSTA = []types.OctupleWord{
	PC_we | MAR_re,                 // T4
	PC_ce,                          // T5
	MDR_we | MAB_lsbre,             // T6
	PC_we | MAR_re,                 // T7
	PC_ce,                          // T8
	MDR_we | MAB_msbre,             // T9
	MAB_lsbwe | MAB_msbwe | MAR_re, // T10
	MDR_re | ACC_we,                // T11
	Noop,                           // T12
}