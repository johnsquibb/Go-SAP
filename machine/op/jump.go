package op

import "Go-SAP3/machine/types"

const (
	JMP = 0xc3
	JM  = 0xfa
	JP  = 0xf2
	JZ  = 0xca
	JNZ = 0xc2
	JC  = 0xda
	JNC = 0xd2
	JPE = 0xea
	JPO = 0xe2
)

var mcJMP = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce,                         // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}

var mcJM = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce | FLG_csfie,             // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}

var mcJP = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce | FLG_csfne,             // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}

var mcJZ = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce | FLG_czfie,             // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}

var mcJNZ = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce | FLG_czfne,             // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}

var mcJC = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce | FLG_ccfie,             // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}

var mcJNC = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce | FLG_ccfne,             // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}

var mcJPE = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce | FLG_cpfie,             // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}

var mcJPO = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce | FLG_cpfne,             // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | PC_re, // T10
	Noop,                          // T11
}
