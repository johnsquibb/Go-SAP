package op

import "Go-SAP/machine/types"

const (
	CALL = 0xcd
	CNZ  = 0xc4
	CZ   = 0xcc
	CNC  = 0xd4
	CC   = 0xdc
	CPO  = 0xe4
	CPE  = 0xec
	CP   = 0xf4
	CM   = 0xfc
)

var mcCALL = []types.OctupleWord{
	PC_we | MAR_re,                // T4 Get LSB for CALL destination address
	MDR_we | MAB_lsbre,            // T5
	PC_ce,                         // T6
	PC_we | MAR_re,                // T7 GET MSB for CALL destination address
	MDR_we | MAB_msbre,            // T8
	PC_ce,                         // T9
	PC_we | GPB_msbre | GPB_lsbre, // T10 Save next instruction address to stack for RET.
	SP_accen | ALU_dcx,            // T11
	SP_we | MAR_re,                // T12
	GPB_msbwe | MDR_re,            // T13
	SP_accen | ALU_dcx,            // T14
	SP_we | MAR_re,                // T15
	GPB_lsbwe | MDR_re,            // T16
	MAB_msbwe | MAB_lsbwe | PC_re, // T17 Set PC to CALL destination address
	Noop,                          // T18
}

var mcCNZ = []types.OctupleWord{
	PC_we | MAR_re,                // T4 Get LSB for CALL destination address
	MDR_we | MAB_lsbre,            // T5
	PC_ce,                         // T6
	PC_we | MAR_re,                // T7 GET MSB for CALL destination address
	MDR_we | MAB_msbre,            // T8
	PC_ce | FLG_czfne,             // T9 (Abort here if flag condition)
	PC_we | GPB_msbre | GPB_lsbre, // T10 Save next instruction address to stack for RET.
	SP_accen | ALU_dcx,            // T11
	SP_we | MAR_re,                // T12
	GPB_msbwe | MDR_re,            // T13
	SP_accen | ALU_dcx,            // T14
	SP_we | MAR_re,                // T15
	GPB_lsbwe | MDR_re,            // T16
	MAB_msbwe | MAB_lsbwe | PC_re, // T17 Set PC to CALL destination address
	Noop,                          // T18
}

var mcCZ = []types.OctupleWord{
	PC_we | MAR_re,                // T4 Get LSB for CALL destination address
	MDR_we | MAB_lsbre,            // T5
	PC_ce,                         // T6
	PC_we | MAR_re,                // T7 GET MSB for CALL destination address
	MDR_we | MAB_msbre,            // T8
	PC_ce | FLG_czfie,             // T9 (Abort here if flag condition)
	PC_we | GPB_msbre | GPB_lsbre, // T10 Save next instruction address to stack for RET.
	SP_accen | ALU_dcx,            // T11
	SP_we | MAR_re,                // T12
	GPB_msbwe | MDR_re,            // T13
	SP_accen | ALU_dcx,            // T14
	SP_we | MAR_re,                // T15
	GPB_lsbwe | MDR_re,            // T16
	MAB_msbwe | MAB_lsbwe | PC_re, // T17 Set PC to CALL destination address
	Noop,                          // T18
}

var mcCNC = []types.OctupleWord{
	PC_we | MAR_re,                // T4 Get LSB for CALL destination address
	MDR_we | MAB_lsbre,            // T5
	PC_ce,                         // T6
	PC_we | MAR_re,                // T7 GET MSB for CALL destination address
	MDR_we | MAB_msbre,            // T8
	PC_ce | FLG_ccfne,             // T9 (Abort here if flag condition)
	PC_we | GPB_msbre | GPB_lsbre, // T10 Save next instruction address to stack for RET.
	SP_accen | ALU_dcx,            // T11
	SP_we | MAR_re,                // T12
	GPB_msbwe | MDR_re,            // T13
	SP_accen | ALU_dcx,            // T14
	SP_we | MAR_re,                // T15
	GPB_lsbwe | MDR_re,            // T16
	MAB_msbwe | MAB_lsbwe | PC_re, // T17 Set PC to CALL destination address
	Noop,                          // T18
}

var mcCC = []types.OctupleWord{
	PC_we | MAR_re,                // T4 Get LSB for CALL destination address
	MDR_we | MAB_lsbre,            // T5
	PC_ce,                         // T6
	PC_we | MAR_re,                // T7 GET MSB for CALL destination address
	MDR_we | MAB_msbre,            // T8
	PC_ce | FLG_ccfie,             // T9 (Abort here if flag condition)
	PC_we | GPB_msbre | GPB_lsbre, // T10 Save next instruction address to stack for RET.
	SP_accen | ALU_dcx,            // T11
	SP_we | MAR_re,                // T12
	GPB_msbwe | MDR_re,            // T13
	SP_accen | ALU_dcx,            // T14
	SP_we | MAR_re,                // T15
	GPB_lsbwe | MDR_re,            // T16
	MAB_msbwe | MAB_lsbwe | PC_re, // T17 Set PC to CALL destination address
	Noop,                          // T18
}

var mcCPO = []types.OctupleWord{
	PC_we | MAR_re,                // T4 Get LSB for CALL destination address
	MDR_we | MAB_lsbre,            // T5
	PC_ce,                         // T6
	PC_we | MAR_re,                // T7 GET MSB for CALL destination address
	MDR_we | MAB_msbre,            // T8
	PC_ce | FLG_cpfne,             // T9 (Abort here if flag condition)
	PC_we | GPB_msbre | GPB_lsbre, // T10 Save next instruction address to stack for RET.
	SP_accen | ALU_dcx,            // T11
	SP_we | MAR_re,                // T12
	GPB_msbwe | MDR_re,            // T13
	SP_accen | ALU_dcx,            // T14
	SP_we | MAR_re,                // T15
	GPB_lsbwe | MDR_re,            // T16
	MAB_msbwe | MAB_lsbwe | PC_re, // T17 Set PC to CALL destination address
	Noop,                          // T18
}

var mcCPE = []types.OctupleWord{
	PC_we | MAR_re,                // T4 Get LSB for CALL destination address
	MDR_we | MAB_lsbre,            // T5
	PC_ce,                         // T6
	PC_we | MAR_re,                // T7 GET MSB for CALL destination address
	MDR_we | MAB_msbre,            // T8
	PC_ce | FLG_cpfie,             // T9 (Abort here if flag condition)
	PC_we | GPB_msbre | GPB_lsbre, // T10 Save next instruction address to stack for RET.
	SP_accen | ALU_dcx,            // T11
	SP_we | MAR_re,                // T12
	GPB_msbwe | MDR_re,            // T13
	SP_accen | ALU_dcx,            // T14
	SP_we | MAR_re,                // T15
	GPB_lsbwe | MDR_re,            // T16
	MAB_msbwe | MAB_lsbwe | PC_re, // T17 Set PC to CALL destination address
	Noop,                          // T18
}

var mcCP = []types.OctupleWord{
	PC_we | MAR_re,                // T4 Get LSB for CALL destination address
	MDR_we | MAB_lsbre,            // T5
	PC_ce,                         // T6
	PC_we | MAR_re,                // T7 GET MSB for CALL destination address
	MDR_we | MAB_msbre,            // T8
	PC_ce | FLG_csfne,             // T9 (Abort here if flag condition)
	PC_we | GPB_msbre | GPB_lsbre, // T10 Save next instruction address to stack for RET.
	SP_accen | ALU_dcx,            // T11
	SP_we | MAR_re,                // T12
	GPB_msbwe | MDR_re,            // T13
	SP_accen | ALU_dcx,            // T14
	SP_we | MAR_re,                // T15
	GPB_lsbwe | MDR_re,            // T16
	MAB_msbwe | MAB_lsbwe | PC_re, // T17 Set PC to CALL destination address
	Noop,                          // T18
}

var mcCM = []types.OctupleWord{
	PC_we | MAR_re,                // T4 Get LSB for CALL destination address
	MDR_we | MAB_lsbre,            // T5
	PC_ce,                         // T6
	PC_we | MAR_re,                // T7 GET MSB for CALL destination address
	MDR_we | MAB_msbre,            // T8
	PC_ce | FLG_csfie,             // T9 (Abort here if flag condition)
	PC_we | GPB_msbre | GPB_lsbre, // T10 Save next instruction address to stack for RET.
	SP_accen | ALU_dcx,            // T11
	SP_we | MAR_re,                // T12
	GPB_msbwe | MDR_re,            // T13
	SP_accen | ALU_dcx,            // T14
	SP_we | MAR_re,                // T15
	GPB_lsbwe | MDR_re,            // T16
	MAB_msbwe | MAB_lsbwe | PC_re, // T17 Set PC to CALL destination address
	Noop,                          // T18
}
