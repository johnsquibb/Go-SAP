package op

import "Go-SAP3/machine/types"

const (
	ANI = 0xe6
	ORI = 0xf6
	XRI = 0xee
	ADI = 0xc6
	ACI = 0xce
	SUI = 0xd6
	SBI = 0xde
	CPI = 0xfe
)

var mcANI = []types.OctupleWord{
	PC_we | MAR_re,                        // T4
	PC_ce,                                 // T5
	MDR_we | TMP_re,                       // T6
	ALU_ana | ALU_we | FLG_enall | ACC_re, // T7
	Noop,                                  // T8
}

var mcORI = []types.OctupleWord{
	PC_we | MAR_re,                        // T4
	PC_ce,                                 // T5
	MDR_we | TMP_re,                       // T6
	ALU_ora | ALU_we | FLG_enall | ACC_re, // T7
	Noop,                                  // T8
}

var mcXRI = []types.OctupleWord{
	PC_we | MAR_re,                        // T4
	PC_ce,                                 // T5
	MDR_we | TMP_re,                       // T6
	ALU_xra | ALU_we | FLG_enall | ACC_re, // T7
	Noop,                                  // T8
}

var mcADI = []types.OctupleWord{
	PC_we | MAR_re,                        // T4
	PC_ce,                                 // T5
	MDR_we | TMP_re,                       // T6
	ALU_add | ALU_we | FLG_enall | ACC_re, // T7
	Noop,                                  // T8
}

var mcACI = []types.OctupleWord{
	PC_we | MAR_re,                        // T4
	PC_ce,                                 // T5
	MDR_we | TMP_re,                       // T6
	ALU_adc | ALU_we | FLG_enall | ACC_re, // T7
	Noop,                                  // T8
}

var mcSUI = []types.OctupleWord{
	PC_we | MAR_re,                        // T4
	PC_ce,                                 // T5
	MDR_we | TMP_re,                       // T6
	ALU_sub | ALU_we | FLG_enall | ACC_re, // T7
	Noop,                                  // T8
}

var mcSBI = []types.OctupleWord{
	PC_we | MAR_re,                        // T4
	PC_ce,                                 // T5
	MDR_we | TMP_re,                       // T6
	ALU_sbb | ALU_we | FLG_enall | ACC_re, // T7
	Noop,                                  // T8
}

var mcCPI = []types.OctupleWord{
	ACC_we | TMP_re,              // T4
	PC_we | MAR_re,               // T5
	PC_ce,                        // T6
	MDR_we | ACC_re,              // T7
	ALU_cmp | ALU_we | FLG_enall, // T8
	TMP_we | ACC_re,              // T9
	Noop,                         // T10
}
