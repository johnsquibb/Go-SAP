package op

import "Go-SAP/machine/types"

const (
	LDA    = 0x3a
	LDAX_B = 0x0a
	LDAX_D = 0x1a
	LHLD   = 0x2a
)

var mcLDA = []types.OctupleWord{
	PC_we | MAR_re,                 // T4
	PC_ce,                          // T5
	MDR_we | MAB_lsbre,             // T6
	PC_we | MAR_re,                 // T7
	PC_ce,                          // T8
	MDR_we | MAB_msbre,             // T9
	MAB_lsbwe | MAB_msbwe | MAR_re, // T10
	MDR_we | ACC_re,                // T11
	Noop,                           // T12
}

var mcLDAX_B = []types.OctupleWord{
	BREG_we | MAB_msbre,            // T4
	CREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | ACC_re,                // T7
	Noop,                           // T8
}

var mcLDAX_D = []types.OctupleWord{
	DREG_we | MAB_msbre,            // T4
	EREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | ACC_re,                // T7
	Noop,                           // T8
}

var mcLHLD = []types.OctupleWord{
	// Read next two memory bytes
	PC_we | MAR_re,     // T4
	PC_ce,              // T5
	MDR_we | MAB_lsbre, // T6
	PC_we | MAR_re,     // T7
	PC_ce,              // T8
	MDR_we | MAB_msbre, // T9

	// Store bytes into GP buffer LSB
	MAB_lsbwe | MAB_msbwe | MAR_re, // T10
	MDR_we | GPB_lsbre,             // T11

	// Move the MAB into HL register pair and increment
	MAB_lsbwe | LREG_re,                        // T12
	MAB_msbwe | HREG_re,                        // T13
	HREG_accen | LREG_accen | ALU_we | ALU_inx, // T14

	// Move the HL register pair values back to MAB
	MAB_lsbre | LREG_we, // T15
	MAB_msbre | HREG_we, // T16

	// Get the next memory address into GP buffer MSB
	MAB_lsbwe | MAB_msbwe | MAR_re, // T17
	MDR_we | GPB_msbre,             // T18

	// Move the GP buffer to HL register pair.
	GPB_lsbwe | LREG_re, //T19
	GPB_msbwe | HREG_re, // T20

	Noop, // T21
}
