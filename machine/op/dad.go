package op

import "Go-SAP/machine/types"

// Double-add BC Register pair to the HL Register pair.
// H and L both act as accumulator during this instruction.
// Byte order is Little-Endian.

const (
	DAD_B  = 0x09
	DAD_D  = 0x19
	DAD_H  = 0x29
	DAD_SP = 0x39
)

var mcDAD_B = []types.OctupleWord{
	CREG_we | TMP_re, // T4
	LREG_accen | ALU_we | ALU_add | FLG_cen | LREG_re, // T5
	BREG_we | TMP_re, // T6
	HREG_accen | ALU_we | ALU_adc | FLG_cen | HREG_re, // T7
	Noop, // T8
}

var mcDAD_D = []types.OctupleWord{
	EREG_we | TMP_re, // T4
	LREG_accen | ALU_we | ALU_add | FLG_cen | LREG_re, // T5
	DREG_we | TMP_re, // T6
	HREG_accen | ALU_we | ALU_adc | FLG_cen | HREG_re, // T7
	Noop, // T8
}

var mcDAD_H = []types.OctupleWord{
	LREG_we | TMP_re, // T4
	LREG_accen | ALU_we | ALU_add | FLG_cen | LREG_re, // T5
	HREG_we | TMP_re, // T6
	HREG_accen | ALU_we | ALU_adc | FLG_cen | HREG_re, // T7
	Noop, // T8
}

var mcDAD_SP = []types.OctupleWord{
	SP_we | GPB_lsbre | GPB_msbre, // T1

	// Add LSB of SP to Reg L
	GPB_lsbwe | TMP_re, // T2
	LREG_accen | ALU_we | ALU_add | FLG_cen | LREG_re, // T3

	// Add MSB of SP to Reg H
	GPB_msbwe | TMP_re, // T4
	HREG_accen | ALU_we | ALU_adc | FLG_cen | HREG_re, // T5
	Noop, // T6
}
