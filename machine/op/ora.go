package op

import "SystemOnePoc/sap3/machine/types"

const (
	ORA_A = 0xb7
	ORA_B = 0xb0
	ORA_C = 0xb1
	ORA_D = 0xb2
	ORA_E = 0xb3
	ORA_H = 0xb4
	ORA_L = 0xb5
	ORA_M = 0xb6
)

var mcORA_A = []types.OctupleWord{
	ACC_we | TMP_re,                      // T4
	ALU_ora | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcORA_B = []types.OctupleWord{
	BREG_we | TMP_re,                      // T4
	ALU_ora | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcORA_C = []types.OctupleWord{
	CREG_we | TMP_re,                      // T4
	ALU_ora | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcORA_D = []types.OctupleWord{
	DREG_we | TMP_re,                      // T4
	ALU_ora | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcORA_E = []types.OctupleWord{
	EREG_we | TMP_re,                      // T4
	ALU_ora | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcORA_H = []types.OctupleWord{
	HREG_we | TMP_re,                      // T4
	ALU_ora | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcORA_L = []types.OctupleWord{
	LREG_we | TMP_re,                      // T4
	ALU_ora | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcORA_M = []types.OctupleWord{
	HREG_we | MAB_msbre,                   // T4
	LREG_we | MAB_lsbre,                   // T5
	MAB_lsbwe | MAB_msbwe | MAR_re,        // T6
	MDR_we | TMP_re,                       // T7
	ALU_we | ALU_ora | FLG_enall | ACC_re, // T8
	Noop,                                  // T9
}