package op

import "Go-SAP3/machine/types"

const (
	DCR_A = 0x3d
	DCR_B = 0x05
	DCR_C = 0x0d
	DCR_D = 0x15
	DCR_E = 0x1d
	DCR_H = 0x25
	DCR_L = 0x2d
	DCR_M = 0x35
)

var mcDCR_A = []types.OctupleWord{
	ALU_dcr | ALU_we | FLG_nc | ACC_re, // T4
	Noop,                               // T5
}

var mcDCR_B = []types.OctupleWord{
	ACC_we | TMP_re,  // T4
	BREG_we | ACC_re, // T5
	ALU_dcr | ALU_we | FLG_nc | BREG_re, // T6
	TMP_we | ACC_re, // T7
	Noop,            // T8
}

var mcDCR_C = []types.OctupleWord{
	ACC_we | TMP_re,  // T4
	CREG_we | ACC_re, // T5
	ALU_dcr | ALU_we | FLG_nc | CREG_re, // T6
	TMP_we | ACC_re, // T7
	Noop,            // T8
}

var mcDCR_D = []types.OctupleWord{
	ACC_we | TMP_re,  // T4
	DREG_we | ACC_re, // T5
	ALU_dcr | ALU_we | FLG_nc | DREG_re, // T6
	TMP_we | ACC_re, // T7
	Noop,            // T8
}

var mcDCR_E = []types.OctupleWord{
	ACC_we | TMP_re,  // T4
	EREG_we | ACC_re, // T5
	ALU_dcr | ALU_we | FLG_nc | EREG_re, // T6
	TMP_we | ACC_re, // T7
	Noop,            // T8
}

var mcDCR_H = []types.OctupleWord{
	ACC_we | TMP_re,  // T4
	HREG_we | ACC_re, // T5
	ALU_dcr | ALU_we | FLG_nc | HREG_re, // T6
	TMP_we | ACC_re, // T7
	Noop,            // T8
}

var mcDCR_L = []types.OctupleWord{
	ACC_we | TMP_re,  // T4
	LREG_we | ACC_re, // T5
	ALU_dcr | ALU_we | FLG_nc | LREG_re, // T6
	TMP_we | ACC_re, // T7
	Noop,            // T8
}

var mcDCR_M = []types.OctupleWord{
	ACC_we | TMP_re,                    // T4
	HREG_we | MAB_msbre,                // T5
	LREG_we | MAB_lsbre,                // T6
	MAB_lsbwe | MAB_msbwe | MAR_re,     // T7
	MDR_we | ACC_re,                    // T8
	ALU_dcr | ALU_we | FLG_nc | MDR_re, // T9
	TMP_we | ACC_re,                    // T10
	Noop,                               // T11
}
