package op

import "SystemOnePoc/sap3/machine/types"

var FetchInstructions = mcFetch

var mcFetch = []types.OctupleWord{
	PC_we | MAR_re, // T1
	PC_ce,          // T2
	MDR_we | IR_re, // T3
}
