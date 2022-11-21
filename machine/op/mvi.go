package op

import "Go-SAP/machine/types"

const (
	MVI_A = 0x3e
	MVI_B = 0x06
	MVI_C = 0x0e
	MVI_D = 0x16
	MVI_E = 0x1e
	MVI_H = 0x26
	MVI_L = 0x2e
	MVI_M = 0x36
)

var mcMVI_A = []types.OctupleWord{
	PC_we | MAR_re,  // T4
	PC_ce,           // T5
	MDR_we | ACC_re, // T6
	Noop,            // T7
}

var mcMVI_B = []types.OctupleWord{
	PC_we | MAR_re,   // T4
	PC_ce,            // T5
	MDR_we | BREG_re, // T6
	Noop,             // T7
}

var mcMVI_C = []types.OctupleWord{
	PC_we | MAR_re,   // T4
	PC_ce,            // T5
	MDR_we | CREG_re, // T6
	Noop,             // T7
}

var mcMVI_D = []types.OctupleWord{
	PC_we | MAR_re,   // T4
	PC_ce,            // T5
	MDR_we | DREG_re, // T6
	Noop,             // T7
}

var mcMVI_E = []types.OctupleWord{
	PC_we | MAR_re,   // T4
	PC_ce,            // T5
	MDR_we | EREG_re, // T6
	Noop,             // T7
}

var mcMVI_H = []types.OctupleWord{
	PC_we | MAR_re,   // T4
	PC_ce,            // T5
	MDR_we | HREG_re, // T6
	Noop,             // T7
}

var mcMVI_L = []types.OctupleWord{
	PC_we | MAR_re,   // T4
	PC_ce,            // T5
	MDR_we | LREG_re, // T6
	Noop,             // T7
}

var mcMVI_M = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	PC_we | MAR_re,                 // T6
	PC_ce,                          // T7
	MDR_we | TMP_re,                // T8
	MAB_lsbwe | MAB_msbwe | MAR_re, // T9
	TMP_we | MDR_re,                // T10
	Noop,                           // T11
}
