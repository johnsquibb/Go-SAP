package op

import "SystemOnePoc/sap3/machine/types"

const (
	CMP_A = 0xbf
	CMP_B = 0xb8
	CMP_C = 0xb9
	CMP_D = 0xba
	CMP_E = 0xbb
	CMP_H = 0xbc
	CMP_L = 0xbd
	CMP_M = 0xbe
)

var mcCMP_A = []types.OctupleWord{
	ACC_we | TMP_re,              // T4
	ALU_cmp | ALU_we | FLG_enall, // T5
	TMP_we | ACC_re,              // T6
	Noop,                         // T7
}

var mcCMP_B = []types.OctupleWord{
	ACC_we | TMP_re,              // T4
	BREG_we | ACC_re,             // T5
	ALU_cmp | ALU_we | FLG_enall, // T6
	TMP_we | ACC_re,              // T7
	Noop,                         // T8
}

var mcCMP_C = []types.OctupleWord{
	ACC_we | TMP_re,              // T4
	CREG_we | ACC_re,             // T5
	ALU_cmp | ALU_we | FLG_enall, // T6
	TMP_we | ACC_re,              // T7
	Noop,                         // T8
}

var mcCMP_D = []types.OctupleWord{
	ACC_we | TMP_re,              // T4
	DREG_we | ACC_re,             // T5
	ALU_cmp | ALU_we | FLG_enall, // T6
	TMP_we | ACC_re,              // T7
	Noop,                         // T8
}

var mcCMP_E = []types.OctupleWord{
	ACC_we | TMP_re,              // T4
	EREG_we | ACC_re,             // T5
	ALU_cmp | ALU_we | FLG_enall, // T6
	TMP_we | ACC_re,              // T7
	Noop,                         // T8
}

var mcCMP_H = []types.OctupleWord{
	ACC_we | TMP_re,              // T4
	HREG_we | ACC_re,             // T5
	ALU_cmp | ALU_we | FLG_enall, // T6
	TMP_we | ACC_re,              // T7
	Noop,                         // T8
}

var mcCMP_L = []types.OctupleWord{
	ACC_we | TMP_re,              // T4
	LREG_we | ACC_re,             // T5
	ALU_cmp | ALU_we | FLG_enall, // T6
	TMP_we | ACC_re,              // T7
	Noop,                         // T8
}

var mcCMP_M = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | TMP_re,                // T7
	ALU_we | ALU_cmp | FLG_enall,   // T8
	Noop,                           // T9
}
