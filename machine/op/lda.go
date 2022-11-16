package op

import "Go-SAP3/machine/types"

const (
	LDA    = 0x3a
	LDAX_B = 0x0a
	LDAX_D = 0x1a
)

var mcLDA = []types.OctupleWord{
	PC_we | MAR_re,                 // T4
	PC_ce,                          // T5
	MDR_we | MAB_lsbre,             // T6
	PC_we | MAR_re,                 // T7
	PC_ce,                          // T8
	MDR_we | MAB_msbre,             // T9
	MAB_lsbwe | MAB_msbwe | MAR_re, // T10
	MDR_we | ACC_re,                // T11
	Noop,                           // T12
}

var mcLDAX_B = []types.OctupleWord{
	BREG_we | MAB_msbre,            // T4
	CREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | ACC_re,                // T7
	Noop,                           // T8
}

var mcLDAX_D = []types.OctupleWord{
	DREG_we | MAB_msbre,            // T4
	EREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | ACC_re,                // T7
	Noop,                           // T8
}
