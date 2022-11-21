package op

import "Go-SAP/machine/types"

const (
	LXI_B  = 0x01
	LXI_D  = 0x11
	LXI_H  = 0x21
	LXI_SP = 0x31
)

// BC Register pair in Little-Endian order.
//
// e.g. When ASM Instruction is LXI B,0x90FF, the generated machine code will be 0x1,0xFF,0x90.
// The register values are: B = 0x90, C = 0xFF
var mcLXI_B = []types.OctupleWord{
	PC_we | MAR_re,   // T4
	PC_ce,            // T5
	MDR_we | CREG_re, // T6
	PC_we | MAR_re,   // T7
	PC_ce,            // T8
	MDR_we | BREG_re, // T9
	Noop,             // T10
}

var mcLXI_D = []types.OctupleWord{
	PC_we | MAR_re,   // T4
	PC_ce,            // T5
	MDR_we | EREG_re, // T6
	PC_we | MAR_re,   // T7
	PC_ce,            // T8
	MDR_we | DREG_re, // T9
	Noop,             // T10
}

var mcLXI_H = []types.OctupleWord{
	PC_we | MAR_re,   // T4
	PC_ce,            // T5
	MDR_we | LREG_re, // T6
	PC_we | MAR_re,   // T7
	PC_ce,            // T8
	MDR_we | HREG_re, // T9
	Noop,             // T10
}

var mcLXI_SP = []types.OctupleWord{
	PC_we | MAR_re,                // T4
	PC_ce,                         // T5
	MDR_we | MAB_lsbre,            // T6
	PC_we | MAR_re,                // T7
	PC_ce,                         // T8
	MDR_we | MAB_msbre,            // T9
	MAB_lsbwe | MAB_msbwe | SP_re, // T10
	Noop,                          // T11
}
