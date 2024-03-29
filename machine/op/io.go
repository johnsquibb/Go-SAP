package op

import "Go-SAP/machine/types"

const (
	IN  = 0xdb
	OUT = 0xd3
)

var mcIN = []types.OctupleWord{
	PC_we | MAR_re, // T4
	PC_ce,          // T5
	MDR_we | IO_ps, // T6
	IO_we | ACC_re, // T7
	Noop,           // T8
}

var mcOUT = []types.OctupleWord{
	PC_we | MAR_re, // T4
	PC_ce,          // T5
	MDR_we | IO_ps, // T6
	ACC_we | IO_re, // T7
	Noop,           // T8
}
