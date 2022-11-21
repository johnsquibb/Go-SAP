package op

import "Go-SAP/machine/types"

const (
	ANA_A = 0xa7
	ANA_B = 0xa0
	ANA_C = 0xa1
	ANA_D = 0xa2
	ANA_E = 0xa3
	ANA_H = 0xa4
	ANA_L = 0xa5
	ANA_M = 0xa6
)

var mcANA_A = []types.OctupleWord{
	ACC_we | TMP_re,                       // T4
	ALU_ana | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcANA_B = []types.OctupleWord{
	BREG_we | TMP_re,                      // T4
	ALU_ana | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcANA_C = []types.OctupleWord{
	CREG_we | TMP_re,                      // T4
	ALU_ana | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcANA_D = []types.OctupleWord{
	DREG_we | TMP_re,                      // T4
	ALU_ana | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcANA_E = []types.OctupleWord{
	EREG_we | TMP_re,                      // T4
	ALU_ana | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcANA_H = []types.OctupleWord{
	HREG_we | TMP_re,                      // T4
	ALU_ana | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcANA_L = []types.OctupleWord{
	LREG_we | TMP_re,                      // T4
	ALU_ana | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcANA_M = []types.OctupleWord{
	HREG_we | MAB_msbre,                   // T4
	LREG_we | MAB_lsbre,                   // T5
	MAB_lsbwe | MAB_msbwe | MAR_re,        // T6
	MDR_we | TMP_re,                       // T7
	ALU_we | ALU_ana | FLG_enall | ACC_re, // T8
	Noop,                                  // T9
}
