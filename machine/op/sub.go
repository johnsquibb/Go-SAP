package op

import "Go-SAP/machine/types"

const (
	SUB_A = 0x97
	SUB_B = 0x90
	SUB_C = 0x91
	SUB_D = 0x92
	SUB_E = 0x93
	SUB_H = 0x94
	SUB_L = 0x95
	SUB_M = 0x96
)

var mcSUB_A = []types.OctupleWord{
	ACC_we | TMP_re,                       // T4
	ALU_we | ALU_sub | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcSUB_B = []types.OctupleWord{
	BREG_we | TMP_re,                      // T4
	ALU_we | ALU_sub | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcSUB_C = []types.OctupleWord{
	CREG_we | TMP_re,                      // T4
	ALU_we | ALU_sub | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcSUB_D = []types.OctupleWord{
	DREG_we | TMP_re,                      // T4
	ALU_we | ALU_sub | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcSUB_E = []types.OctupleWord{
	EREG_we | TMP_re,                      // T4
	ALU_we | ALU_sub | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcSUB_H = []types.OctupleWord{
	HREG_we | TMP_re,                      // T4
	ALU_we | ALU_sub | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcSUB_L = []types.OctupleWord{
	LREG_we | TMP_re,                      // T4
	ALU_we | ALU_sub | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcSUB_M = []types.OctupleWord{
	HREG_we | MAB_msbre,                   // T4
	LREG_we | MAB_lsbre,                   // T5
	MAB_lsbwe | MAB_msbwe | MAR_re,        // T6
	MDR_we | TMP_re,                       // T7
	ALU_we | ALU_sub | FLG_enall | ACC_re, // T8
	Noop,                                  // T9
}
