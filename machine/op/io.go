package op

import "Go-SAP3/machine/types"

const (
	IN = 0xdb
	OUT = 0xd3
)

var mcIN = []types.OctupleWord{
	PC_we | MAR_re, // T4
	PC_ce,          // T5
	MDR_we | IN_ps, // T6
	IN_we | ACC_re, // T7
	Noop,           // T8
}

var mcOUT = []types.OctupleWord{
	PC_we | MAR_re,  // T4
	PC_ce,           // T5
	MDR_we | OUT_ps, // T6
	ACC_we | OUT_re, // T7
	Noop,            // T8
}
