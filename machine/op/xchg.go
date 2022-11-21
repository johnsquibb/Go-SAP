package op

import "Go-SAP/machine/types"

const (
	XCHG = 0xEB
)

var mcXCHG = []types.OctupleWord{
	// Move DE to GP buffer
	DREG_we | GPB_msbre, // T4
	EREG_we | GPB_lsbre, // T5

	// Move HL to DE
	HREG_we | DREG_re, // T6
	LREG_we | EREG_re, // T7

	// Move GP to HL
	GPB_msbwe | HREG_re, // T8
	GPB_lsbwe | LREG_re, // T9

	Noop, // T10
}
