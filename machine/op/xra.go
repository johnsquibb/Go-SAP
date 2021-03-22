package op

import "SystemOnePoc/sap3/machine/types"

const (
	XRA_A = 0xaf
	XRA_B = 0xa8
	XRA_C = 0xa9
	XRA_D = 0xaa
	XRA_E = 0xab
	XRA_H = 0xac
	XRA_L = 0xad
	XRA_M = 0xae
)

var mcXRA_A = []types.OctupleWord{
	BREG_we | TMP_re,                      // T4
	ALU_xra | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcXRA_B = []types.OctupleWord{
	BREG_we | TMP_re,                      // T4
	ALU_xra | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcXRA_C = []types.OctupleWord{
	CREG_we | TMP_re,                      // T4
	ALU_xra | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcXRA_D = []types.OctupleWord{
	DREG_we | TMP_re,                      // T4
	ALU_xra | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcXRA_E = []types.OctupleWord{
	EREG_we | TMP_re,                      // T4
	ALU_xra | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcXRA_H = []types.OctupleWord{
	HREG_we | TMP_re,                      // T4
	ALU_xra | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcXRA_L = []types.OctupleWord{
	LREG_we | TMP_re,                      // T4
	ALU_xra | ALU_we | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcXRA_M = []types.OctupleWord{
	HREG_we | MAB_msbre,                   // T4
	LREG_we | MAB_lsbre,                   // T5
	MAB_lsbwe | MAB_msbwe | MAR_re,        // T6
	MDR_we | TMP_re,                       // T7
	ALU_we | ALU_xra | FLG_enall | ACC_re, // T8
	Noop,                                  // T9
}
