package op

import "Go-SAP/machine/types"

const (
	CMA = 0x2f
)

var mcCMA = []types.OctupleWord{
	ALU_cma | ALU_we | ACC_re, // T4
	Noop,
}
