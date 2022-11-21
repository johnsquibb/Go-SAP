package op

import "Go-SAP/machine/types"

const (
	ADC_A = 0x8f
	ADC_B = 0x88
	ADC_C = 0x89
	ADC_D = 0x8a
	ADC_E = 0x8b
	ADC_H = 0x8c
	ADC_L = 0x8d
	ADC_M = 0x8e
)

var mcADC_A = []types.OctupleWord{
	ACC_we | TMP_re,                       // T4
	ALU_we | ALU_adc | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADC_B = []types.OctupleWord{
	BREG_we | TMP_re,                      // T4
	ALU_we | ALU_adc | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADC_C = []types.OctupleWord{
	CREG_we | TMP_re,                      // T4
	ALU_we | ALU_adc | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADC_D = []types.OctupleWord{
	DREG_we | TMP_re,                      // T4
	ALU_we | ALU_adc | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADC_E = []types.OctupleWord{
	EREG_we | TMP_re,                      // T4
	ALU_we | ALU_adc | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADC_H = []types.OctupleWord{
	HREG_we | TMP_re,                      // T4
	ALU_we | ALU_adc | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADC_L = []types.OctupleWord{
	LREG_we | TMP_re,                      // T4
	ALU_we | ALU_adc | FLG_enall | ACC_re, // T5
	Noop,                                  // T6
}

var mcADC_M = []types.OctupleWord{
	HREG_we | MAB_msbre,                   // T4
	LREG_we | MAB_lsbre,                   // T5
	MAB_lsbwe | MAB_msbwe | MAR_re,        // T6
	MDR_we | TMP_re,                       // T7
	ALU_we | ALU_adc | FLG_enall | ACC_re, // T8
	Noop,                                  // T9
}
