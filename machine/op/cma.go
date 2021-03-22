package op

import "SystemOnePoc/sap3/machine/types"

const (
	CMA = 0x2f
)

var mcCMA = []types.OctupleWord{
	ALU_cma | ALU_we | ACC_re, // T4
	Noop,
}

