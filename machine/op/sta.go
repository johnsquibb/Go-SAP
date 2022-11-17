package op

import "Go-SAP3/machine/types"

const (
	STA    = 0x32
	SHLD   = 0x22
	STAX_B = 0x02
	STAX_D = 0x12
)

var mcSTA = []types.OctupleWord{
	PC_we | MAR_re,                 // T4
	PC_ce,                          // T5
	MDR_we | MAB_lsbre,             // T6
	PC_we | MAR_re,                 // T7
	PC_ce,                          // T8
	MDR_we | MAB_msbre,             // T9
	MAB_lsbwe | MAB_msbwe | MAR_re, // T10
	MDR_re | ACC_we,                // T11
	Noop,                           // T12
}

var mcSTAX_B = []types.OctupleWord{
	BREG_we | MAB_msbre,            // T4
	CREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_re | ACC_we,                // T7
	Noop,                           // T8
}

var mcSTAX_D = []types.OctupleWord{
	DREG_we | MAB_msbre,            // T4
	EREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_re | ACC_we,                // T7
	Noop,                           // T8
}

var mcSHLD = []types.OctupleWord{
	// Move the HL register values in GP buffer
	GPB_lsbre | LREG_we, // T4
	GPB_msbre | HREG_we, // T5

	// Read next two memory bytes
	PC_we | MAR_re,     // T6
	PC_ce,              // T7
	MDR_we | MAB_lsbre, // T8
	PC_we | MAR_re,     // T9
	PC_ce,              // T10
	MDR_we | MAB_msbre, // T11

	// Store LSB in first memory address
	MAB_lsbwe | MAB_msbwe | MAR_re, // T12
	MDR_re | GPB_lsbwe,             // T13

	// Copy the MAB into HL register pair and increment
	MAB_lsbwe | LREG_re,                        // T14
	MAB_msbwe | HREG_re,                        // T15
	HREG_accen | LREG_accen | ALU_we | ALU_inx, // T16

	// Copy the HL register pair values back to MAB
	MAB_lsbre | LREG_we, // T17
	MAB_msbre | HREG_we, // T18

	// Store MSB in second memory address
	MAB_lsbwe | MAB_msbwe | MAR_re, // T19
	MDR_re | GPB_msbwe,             // T20

	Noop, // T21
}
