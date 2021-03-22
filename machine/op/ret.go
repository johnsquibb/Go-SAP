package op

import "SystemOnePoc/sap3/machine/types"

const (
	RET = 0xc9
	RNZ = 0xc0
	RZ  = 0xc8
	RNC = 0xd0
	RC  = 0xd8
	RPO = 0xe0
	RPE = 0xe8
	RP  = 0xf0
	RM  = 0xf8
)

var mcRET = []types.OctupleWord{
	SP_we | MAR_re,                // T4
	MAB_lsbre | MDR_we,            // T5
	SP_accen | ALU_inr,            // T6
	SP_we | MAR_re,                // T7
	MAB_msbre | MDR_we,            // T8
	SP_accen | ALU_inr,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}

var mcRNZ = []types.OctupleWord{
	FLG_czfne,                     // T4
	SP_we | MAR_re,                // T5
	MAB_lsbre | MDR_we,            // T6
	SP_accen | ALU_inr,            // T7
	SP_we | MAR_re,                // T8
	MAB_msbre | MDR_we,            // T9
	SP_accen | ALU_inr,            // T10
	MAB_lsbwe | MAB_msbwe | PC_re, // T11
	Noop,                          // T12
}

var mcRZ = []types.OctupleWord{
	FLG_czfie,                     // T4
	SP_we | MAR_re,                // T5
	MAB_lsbre | MDR_we,            // T6
	SP_accen | ALU_inr,            // T7
	SP_we | MAR_re,                // T8
	MAB_msbre | MDR_we,            // T9
	SP_accen | ALU_inr,            // T10
	MAB_lsbwe | MAB_msbwe | PC_re, // T11
	Noop,                          // T12
}

var mcRNC = []types.OctupleWord{
	FLG_ccfne,                     // T4
	SP_we | MAR_re,                // T5
	MAB_lsbre | MDR_we,            // T6
	SP_accen | ALU_inr,            // T7
	SP_we | MAR_re,                // T8
	MAB_msbre | MDR_we,            // T9
	SP_accen | ALU_inr,            // T10
	MAB_lsbwe | MAB_msbwe | PC_re, // T11
	Noop,                          // T12
}

var mcRC = []types.OctupleWord{
	FLG_ccfie,                     // T4
	SP_we | MAR_re,                // T5
	MAB_lsbre | MDR_we,            // T6
	SP_accen | ALU_inr,            // T7
	SP_we | MAR_re,                // T8
	MAB_msbre | MDR_we,            // T9
	SP_accen | ALU_inr,            // T10
	MAB_lsbwe | MAB_msbwe | PC_re, // T11
	Noop,                          // T12
}

var mcRPO = []types.OctupleWord{
	FLG_cpfne,                     // T4
	SP_we | MAR_re,                // T5
	MAB_lsbre | MDR_we,            // T6
	SP_accen | ALU_inr,            // T7
	SP_we | MAR_re,                // T8
	MAB_msbre | MDR_we,            // T9
	SP_accen | ALU_inr,            // T10
	MAB_lsbwe | MAB_msbwe | PC_re, // T11
	Noop,                          // T12
}

var mcRPE = []types.OctupleWord{
	FLG_cpfie,                     // T4
	SP_we | MAR_re,                // T5
	MAB_lsbre | MDR_we,            // T6
	SP_accen | ALU_inr,            // T7
	SP_we | MAR_re,                // T8
	MAB_msbre | MDR_we,            // T9
	SP_accen | ALU_inr,            // T10
	MAB_lsbwe | MAB_msbwe | PC_re, // T11
	Noop,                          // T12
}

var mcRP = []types.OctupleWord{
	FLG_csfne,                     // T4
	SP_we | MAR_re,                // T5
	MAB_lsbre | MDR_we,            // T6
	SP_accen | ALU_inr,            // T7
	SP_we | MAR_re,                // T8
	MAB_msbre | MDR_we,            // T9
	SP_accen | ALU_inr,            // T10
	MAB_lsbwe | MAB_msbwe | PC_re, // T11
	Noop,                          // T12
}

var mcRM = []types.OctupleWord{
	FLG_csfie,                     // T4
	SP_we | MAR_re,                // T5
	MAB_lsbre | MDR_we,            // T6
	SP_accen | ALU_inr,            // T7
	SP_we | MAR_re,                // T8
	MAB_msbre | MDR_we,            // T9
	SP_accen | ALU_inr,            // T10
	MAB_lsbwe | MAB_msbwe | PC_re, // T11
	Noop,                          // T12
}