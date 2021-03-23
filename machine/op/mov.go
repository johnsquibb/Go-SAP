package op

import "Go-SAP3/machine/types"

const (
	MOV_A_A = 0x7f
	MOV_A_B = 0x78
	MOV_A_C = 0x79
	MOV_A_D = 0x7a
	MOV_A_E = 0x7b
	MOV_A_H = 0x7c
	MOV_A_L = 0x7d
	MOV_A_M = 0x7e

	MOV_B_A = 0x47
	MOV_B_B = 0x40
	MOV_B_C = 0x41
	MOV_B_D = 0x42
	MOV_B_E = 0x43
	MOV_B_H = 0x44
	MOV_B_L = 0x45
	MOV_B_M = 0x46

	MOV_C_A = 0x4f
	MOV_C_B = 0x48
	MOV_C_C = 0x49
	MOV_C_D = 0x4a
	MOV_C_E = 0x4b
	MOV_C_H = 0x4c
	MOV_C_L = 0x4d
	MOV_C_M = 0x4e

	MOV_D_A = 0x57
	MOV_D_B = 0x50
	MOV_D_C = 0x51
	MOV_D_D = 0x52
	MOV_D_E = 0x53
	MOV_D_H = 0x54
	MOV_D_L = 0x55
	MOV_D_M = 0x56

	MOV_E_A = 0x5f
	MOV_E_B = 0x58
	MOV_E_C = 0x59
	MOV_E_D = 0x5a
	MOV_E_E = 0x5b
	MOV_E_H = 0x5c
	MOV_E_L = 0x5d
	MOV_E_M = 0x5e

	MOV_H_A = 0x67
	MOV_H_B = 0x60
	MOV_H_C = 0x61
	MOV_H_D = 0x62
	MOV_H_E = 0x63
	MOV_H_H = 0x64
	MOV_H_L = 0x65
	MOV_H_M = 0x66

	MOV_L_A = 0x6f
	MOV_L_B = 0x68
	MOV_L_C = 0x69
	MOV_L_D = 0x6a
	MOV_L_E = 0x6b
	MOV_L_H = 0x6c
	MOV_L_L = 0x6d
	MOV_L_M = 0x6e

	MOV_M_A = 0x77
	MOV_M_B = 0x70
	MOV_M_C = 0x71
	MOV_M_D = 0x72
	MOV_M_E = 0x73
	MOV_M_H = 0x74
	MOV_M_L = 0x75
)

// Accumulator

var mcMOV_A_A = []types.OctupleWord{
	Noop, // T5
}

var mcMOV_A_B = []types.OctupleWord{
	BREG_we | ACC_re, // T4
	Noop,             // T5
}

var mcMOV_A_C = []types.OctupleWord{
	CREG_we | ACC_re, // T4
	Noop,             // T5
}

var mcMOV_A_D = []types.OctupleWord{
	DREG_we | ACC_re, // T4
	Noop,             // T5
}

var mcMOV_A_E = []types.OctupleWord{
	EREG_we | ACC_re, // T4
	Noop,             // T5
}

var mcMOV_A_H = []types.OctupleWord{
	HREG_we | ACC_re, // T4
	Noop,             // T5
}

var mcMOV_A_L = []types.OctupleWord{
	LREG_we | ACC_re, // T4
	Noop,             // T5
}

var mcMOV_A_M = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | ACC_re,                // T7
	Noop,                           // T8
}

// Register B

var mcMOV_B_A = []types.OctupleWord{
	ACC_we | BREG_re, // T4
	Noop,             // T5
}

var mcMOV_B_B = []types.OctupleWord{
	Noop, // T5
}

var mcMOV_B_C = []types.OctupleWord{
	CREG_we | BREG_re, // T4
	Noop,              // T5
}

var mcMOV_B_D = []types.OctupleWord{
	DREG_we | BREG_re, // T4
	Noop,              // T5
}

var mcMOV_B_E = []types.OctupleWord{
	EREG_we | BREG_re, // T4
	Noop,              // T5
}

var mcMOV_B_H = []types.OctupleWord{
	HREG_we | BREG_re, // T4
	Noop,              // T5
}

var mcMOV_B_L = []types.OctupleWord{
	LREG_we | BREG_re, // T4
	Noop,              // T5
}

var mcMOV_B_M = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | BREG_re,               // T7
	Noop,                           // T8
}

// Register C

var mcMOV_C_A = []types.OctupleWord{
	ACC_we | CREG_re, // T4
	Noop,             // T5
}

var mcMOV_C_B = []types.OctupleWord{
	BREG_we | CREG_re, // T4
	Noop,              // T5
}

var mcMOV_C_C = []types.OctupleWord{
	Noop, // T5
}

var mcMOV_C_D = []types.OctupleWord{
	DREG_we | CREG_re, // T4
	Noop,              // T5
}

var mcMOV_C_E = []types.OctupleWord{
	EREG_we | CREG_re, // T4
	Noop,              // T5
}

var mcMOV_C_H = []types.OctupleWord{
	HREG_we | CREG_re, // T4
	Noop,              // T5
}

var mcMOV_C_L = []types.OctupleWord{
	LREG_we | CREG_re, // T4
	Noop,              // T5
}

var mcMOV_C_M = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | CREG_re,               // T7
	Noop,                           // T8
}

// Register D

var mcMOV_D_A = []types.OctupleWord{
	ACC_we | DREG_re, // T4
	Noop,             // T5
}

var mcMOV_D_B = []types.OctupleWord{
	BREG_we | DREG_re, // T4
	Noop,              // T5
}

var mcMOV_D_C = []types.OctupleWord{
	CREG_we | DREG_re, // T4
	Noop,              // T5
}

var mcMOV_D_D = []types.OctupleWord{
	Noop, // T5
}

var mcMOV_D_E = []types.OctupleWord{
	EREG_we | DREG_re, // T4
	Noop,              // T5
}

var mcMOV_D_H = []types.OctupleWord{
	HREG_we | DREG_re, // T4
	Noop,              // T5
}

var mcMOV_D_L = []types.OctupleWord{
	LREG_we | DREG_re, // T4
	Noop,              // T5
}

var mcMOV_D_M = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | DREG_re,               // T7
	Noop,                           // T8
}

// Register E

var mcMOV_E_A = []types.OctupleWord{
	ACC_we | EREG_re, // T4
	Noop,             // T5
}

var mcMOV_E_B = []types.OctupleWord{
	BREG_we | EREG_re, // T4
	Noop,              // T5
}

var mcMOV_E_C = []types.OctupleWord{
	CREG_we | EREG_re, // T4
	Noop,              // T5
}

var mcMOV_E_D = []types.OctupleWord{
	DREG_we | EREG_re, // T4
	Noop,              // T5
}

var mcMOV_E_E = []types.OctupleWord{
	Noop, // T5
}

var mcMOV_E_H = []types.OctupleWord{
	HREG_we | EREG_re, // T4
	Noop,              // T5
}

var mcMOV_E_L = []types.OctupleWord{
	LREG_we | EREG_re, // T4
	Noop,              // T5
}

var mcMOV_E_M = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | EREG_re,               // T7
	Noop,                           // T8
}

// Register H

var mcMOV_H_A = []types.OctupleWord{
	ACC_we | HREG_re, // T4
	Noop,             // T5
}

var mcMOV_H_B = []types.OctupleWord{
	BREG_we | HREG_re, // T4
	Noop,              // T5
}

var mcMOV_H_C = []types.OctupleWord{
	CREG_we | HREG_re, // T4
	Noop,              // T5
}

var mcMOV_H_D = []types.OctupleWord{
	DREG_we | HREG_re, // T4
	Noop,              // T5
}

var mcMOV_H_E = []types.OctupleWord{
	EREG_we | HREG_re, // T4
	Noop,              // T5
}

var mcMOV_H_H = []types.OctupleWord{
	Noop, // T5
}

var mcMOV_H_L = []types.OctupleWord{
	LREG_we | HREG_re, // T4
	Noop,              // T5
}

var mcMOV_H_M = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | HREG_re,               // T7
	Noop,                           // T8
}

// Register L

var mcMOV_L_A = []types.OctupleWord{
	ACC_we | LREG_re, // T4
	Noop,             // T5
}

var mcMOV_L_B = []types.OctupleWord{
	BREG_we | LREG_re, // T4
	Noop,              // T5
}

var mcMOV_L_C = []types.OctupleWord{
	CREG_we | LREG_re, // T4
	Noop,              // T5
}

var mcMOV_L_D = []types.OctupleWord{
	DREG_we | LREG_re, // T4
	Noop,              // T5
}

var mcMOV_L_E = []types.OctupleWord{
	EREG_we | LREG_re, // T4
	Noop,              // T5
}

var mcMOV_L_H = []types.OctupleWord{
	HREG_we | LREG_re, // T4
	Noop,              // T5
}

var mcMOV_L_L = []types.OctupleWord{
	Noop, // T5
}

var mcMOV_L_M = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	MDR_we | LREG_re,               // T7
	Noop,                           // T8
}

// Register M
var mcMOV_M_A = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	ACC_we | MDR_re,                // T7
	Noop,                           // T8
}

var mcMOV_M_B = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	BREG_we | MDR_re,                // T7
	Noop,                           // T8
}

var mcMOV_M_C = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	CREG_we | MDR_re,                // T7
	Noop,                           // T8
}

var mcMOV_M_D = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	DREG_we | MDR_re,                // T7
	Noop,                           // T8
}

var mcMOV_M_E = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	EREG_we | MDR_re,                // T7
	Noop,                           // T8
}

var mcMOV_M_H = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	HREG_we | MDR_re,               // T7
	Noop,                           // T8
}

var mcMOV_M_L = []types.OctupleWord{
	HREG_we | MAB_msbre,            // T4
	LREG_we | MAB_lsbre,            // T5
	MAB_lsbwe | MAB_msbwe | MAR_re, // T6
	LREG_we | MDR_re,                // T7
	Noop,                           // T8
}