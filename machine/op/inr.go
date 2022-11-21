package op

import "Go-SAP/machine/types"

const (
	INR_A = 0x3c
	INR_B = 0x04
	INR_C = 0x0c
	INR_D = 0x14
	INR_E = 0x1C
	INR_H = 0x24
	INR_L = 0x2C
	INR_M = 0x34
)

var mcINR_A = []types.OctupleWord{
	ALU_inr | ALU_we | FLG_nc | ACC_re, // T4
	Noop,                               // T5
}

var mcINR_B = []types.OctupleWord{
	ACC_we | TMP_re,                     // T4
	BREG_we | ACC_re,                    // T5
	ALU_inr | ALU_we | FLG_nc | BREG_re, // T6
	TMP_we | ACC_re,                     // T7
	Noop,                                // T8
}

var mcINR_C = []types.OctupleWord{
	ACC_we | TMP_re,                     // T4
	CREG_we | ACC_re,                    // T5
	ALU_inr | ALU_we | FLG_nc | CREG_re, // T6
	TMP_we | ACC_re,                     // T7
	Noop,                                // T8
}

var mcINR_D = []types.OctupleWord{
	ACC_we | TMP_re,                     // T4
	DREG_we | ACC_re,                    // T5
	ALU_inr | ALU_we | FLG_nc | DREG_re, // T6
	TMP_we | ACC_re,                     // T7
	Noop,                                // T8
}

var mcINR_E = []types.OctupleWord{
	ACC_we | TMP_re,                     // T4
	EREG_we | ACC_re,                    // T5
	ALU_inr | ALU_we | FLG_nc | EREG_re, // T6
	TMP_we | ACC_re,                     // T7
	Noop,                                // T8
}

var mcINR_H = []types.OctupleWord{
	ACC_we | TMP_re,                     // T4
	HREG_we | ACC_re,                    // T5
	ALU_inr | ALU_we | FLG_nc | HREG_re, // T6
	TMP_we | ACC_re,                     // T7
	Noop,                                // T8
}

var mcINR_L = []types.OctupleWord{
	ACC_we | TMP_re,                     // T4
	LREG_we | ACC_re,                    // T5
	ALU_inr | ALU_we | FLG_nc | LREG_re, // T6
	TMP_we | ACC_re,                     // T7
	Noop,                                // T8
}

var mcINR_M = []types.OctupleWord{
	ACC_we | TMP_re,                    // T4
	HREG_we | MAB_msbre,                // T5
	LREG_we | MAB_lsbre,                // T6
	MAB_lsbwe | MAB_msbwe | MAR_re,     // T7
	MDR_we | ACC_re,                    // T8
	ALU_inr | ALU_we | FLG_nc | MDR_re, // T9
	TMP_we | ACC_re,                    // T10
	Noop,                               // T11
}
