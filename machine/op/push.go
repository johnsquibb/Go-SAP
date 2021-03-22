package op

import "SystemOnePoc/sap3/machine/types"

const (
	PUSH_B   = 0xc5
	PUSH_D   = 0xd5
	PUSH_H   = 0xe5
	PUSH_PSW = 0xf5
)

var mcPUSH_B = []types.OctupleWord{
	SP_accen | ALU_dcr, // T4
	SP_we | MAR_re,     // T5
	BREG_we | MDR_re,   // T6
	SP_accen | ALU_dcr, // T7
	SP_we | MAR_re,     // T8
	CREG_we | MDR_re,   // T9
	Noop,               // T10
}

var mcPUSH_D = []types.OctupleWord{
	SP_accen | ALU_dcr, // T4
	SP_we | MAR_re,     // T5
	DREG_we | MDR_re,   // T6
	SP_accen | ALU_dcr, // T7
	SP_we | MAR_re,     // T8
	EREG_we | MDR_re,   // T9
	Noop,               // T10
}

var mcPUSH_H = []types.OctupleWord{
	SP_accen | ALU_dcr, // T4
	SP_we | MAR_re,     // T5
	HREG_we | MDR_re,   // T6
	SP_accen | ALU_dcr, // T7
	SP_we | MAR_re,     // T8
	LREG_we | MDR_re,   // T9
	Noop,               // T10
}

var mcPUSH_PSW = []types.OctupleWord{
	SP_accen | ALU_dcr, // T4
	SP_we | MAR_re,     // T5
	ACC_we | MDR_re,    // T6
	SP_accen | ALU_dcr, // T7
	SP_we | MAR_re,     // T8
	FREG_we | MDR_re,   // T9
	Noop,               // T10
}
