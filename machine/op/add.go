package op

import "Go-SAP3/machine/types"

const (
	ADD_A = 0x87
	ADD_B = 0x80
	ADD_C = 0x81
	ADD_D = 0x82
	ADD_E = 0x83
	ADD_H = 0x84
	ADD_L = 0x85
	ADD_M = 0x86
)

var mcADD_A = []types.OctupleWord{
	ACC_we | TMP_re,                       // T4
	ALU_we | ALU_add | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADD_B = []types.OctupleWord{
	BREG_we | TMP_re,                      // T4
	ALU_we | ALU_add | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADD_C = []types.OctupleWord{
	CREG_we | TMP_re,                      // T4
	ALU_we | ALU_add | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADD_D = []types.OctupleWord{
	DREG_we | TMP_re,                      // T4
	ALU_we | ALU_add | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADD_E = []types.OctupleWord{
	EREG_we | TMP_re,                      // T4
	ALU_we | ALU_add | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADD_H = []types.OctupleWord{
	HREG_we | TMP_re,                      // T4
	ALU_we | ALU_add | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADD_L = []types.OctupleWord{
	LREG_we | TMP_re,                      // T4
	ALU_we | ALU_add | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADD_M = []types.OctupleWord{
	HREG_we | MAB_msbre,                   // T4
	LREG_we | MAB_lsbre,                   // T5
	MAB_lsbwe | MAB_msbwe | MAR_re,        // T6
	MDR_we | TMP_re,                       // T7
	ALU_we | ALU_add | FLG_enall | ACC_re, // T8
	Noop,                                  // T9
}
