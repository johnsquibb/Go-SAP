package op

import "SystemOnePoc/sap3/machine/types"

const (
	SBB_A = 0x9f
	SBB_B = 0x98
	SBB_C = 0x99
	SBB_D = 0x9a
	SBB_E = 0x9b
	SBB_H = 0x9c
	SBB_L = 0x9d
	SBB_M = 0x9e
)

var mcSBB_A = []types.OctupleWord{
	ACC_we | TMP_re,                       // T4
	ALU_we | ALU_sbb | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcSBB_B = []types.OctupleWord{
	BREG_we | TMP_re,                   // T4
	ALU_we | ALU_sbb | FLG_enall | ACC_re, // T5
	Noop,                               // T6
}

var mcSBB_C = []types.OctupleWord{
	CREG_we | TMP_re,                   // T4
	ALU_we | ALU_sbb | FLG_enall | ACC_re, // T5
	Noop,                               // T6
}

var mcSBB_D = []types.OctupleWord{
	DREG_we | TMP_re,                   // T4
	ALU_we | ALU_sbb | FLG_enall | ACC_re, // T5
	Noop,                               // T6
}

var mcSBB_E = []types.OctupleWord{
	EREG_we | TMP_re,                   // T4
	ALU_we | ALU_sbb | FLG_enall | ACC_re, // T5
	Noop,                               // T6
}

var mcSBB_H = []types.OctupleWord{
	HREG_we | TMP_re,                   // T4
	ALU_we | ALU_sbb | FLG_enall | ACC_re, // T5
	Noop,                               // T6
}

var mcSBB_L = []types.OctupleWord{
	LREG_we | TMP_re,                   // T4
	ALU_we | ALU_sbb | FLG_enall | ACC_re, // T5
	Noop,                               // T6
}

var mcSBB_M = []types.OctupleWord{
	HREG_we | MAB_msbre,                   // T4
	LREG_we | MAB_lsbre,                   // T5
	MAB_lsbwe | MAB_msbwe | MAR_re,        // T6
	MDR_we | TMP_re,                       // T7
	ALU_we | ALU_sbb | FLG_enall | ACC_re, // T8
	Noop,                                  // T9
}

