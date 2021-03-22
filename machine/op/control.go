package op

// Control bit flags.
const (
	Noop       = 0x0                // 0
	PC_re      = 0x1                // 1
	PC_we      = 0x2                // 2
	PC_ce      = 0x4                // 3
	MAR_re     = 0x8                // 4
	MDR_re     = 0x10               // 5
	MDR_we     = 0x20               // 6
	MAB_lsbre  = 0x40               // 7
	MAB_lsbwe  = 0x80               // 8
	MAB_msbre  = 0x100              // 9
	MAB_msbwe  = 0x200              // 10
	GPB_lsbre  = 0x400              // 11
	GPB_lsbwe  = 0x800              // 12
	GPB_msbre  = 0x1000             // 13
	GPB_msbwe  = 0x2000             // 14
	IR_re      = 0x4000             // 15
	ACC_re     = 0x8000             // 16
	ACC_we     = 0x10000            // 17
	TMP_re     = 0x20000            // 18
	TMP_we     = 0x40000            // 19
	BREG_re    = 0x80000            // 20
	BREG_we    = 0x100000           // 21
	BREG_accen = 0x200000           // 22
	CREG_re    = 0x400000           // 23
	CREG_we    = 0x800000           // 24
	CREG_accen = 0x1000000          // 25
	DREG_re    = 0x2000000          // 26
	DREG_we    = 0x4000000          // 27
	DREG_accen = 0x8000000          // 28
	EREG_re    = 0x10000000         // 29
	EREG_we    = 0x20000000         // 30
	EREG_accen = 0x40000000         // 31
	FREG_re    = 0x80000000         // 32
	FREG_we    = 0x100000000        // 33
	HREG_re    = 0x200000000        // 34
	HREG_we    = 0x400000000        // 35
	HREG_accen = 0x800000000        // 36
	LREG_re    = 0x1000000000       // 37
	LREG_we    = 0x2000000000       // 38
	LREG_accen = 0x4000000000       // 39
	SP_re      = 0x8000000000       // 40
	SP_we      = 0x10000000000      // 41
	SP_accen   = 0x20000000000      // 42
	ALU_we     = 0x40000000000      // 43
	ALU_b0     = 0x80000000000      // 44
	ALU_b1     = 0x100000000000     // 45
	ALU_b2     = 0x200000000000     // 46
	ALU_b3     = 0x400000000000     // 47
	ALU_b4     = 0x800000000000     // 48
	FLG_csfie  = 0x1000000000000    // 49
	FLG_csfne  = 0x2000000000000    // 50
	FLG_czfie  = 0x4000000000000    // 51
	FLG_czfne  = 0x8000000000000    // 52
	FLG_ccfie  = 0x10000000000000   // 53
	FLG_ccfne  = 0x20000000000000   // 54
	FLG_cpfie  = 0x40000000000000   // 55
	FLG_cpfne  = 0x80000000000000   // 56
	FLG_cen    = 0x100000000000000  // 57
	FLG_zen    = 0x200000000000000  // 58
	FLG_sen    = 0x400000000000000  // 59
	FLG_pen    = 0x800000000000000  // 60
	IN_ps      = 0x1000000000000000 // 61
	IN_we      = 0x2000000000000000 // 62
	OUT_ps     = 0x4000000000000000 // 63
	OUT_re     = 0x8000000000000000 // 64
)

// Control bit flags for AluMode readability.
const (
	// Operations.
	ALU_add = 0 // ADD is default operation.
	ALU_sub = ALU_b0
	ALU_inr = ALU_b1
	ALU_dcr = ALU_b1 | ALU_b0
	ALU_cma = ALU_b2
	ALU_ana = ALU_b2 | ALU_b0
	ALU_ora = ALU_b2 | ALU_b1
	ALU_xra = ALU_b2 | ALU_b1 | ALU_b0
	ALU_ral = ALU_b3
	ALU_rar = ALU_b3 | ALU_b0
	ALU_adc = ALU_b3 | ALU_b1
	ALU_sbb = ALU_b3 | ALU_b1 | ALU_b0
	ALU_stc = ALU_b3 | ALU_b2
	ALU_cmc = ALU_b3 | ALU_b2 | ALU_b0
	ALU_rlc = ALU_b3 | ALU_b2 | ALU_b1
	ALU_rrc = ALU_b3 | ALU_b2 | ALU_b1 | ALU_b0
	ALU_cmp = ALU_b4
	ALU_inx = ALU_b4 | ALU_b0
	ALU_dcx = ALU_b4 | ALU_b1
)

// Control bit flags for flags readability.
const (
	FLG_enall = FLG_cen | FLG_sen | FLG_zen | FLG_pen
	FLG_nc    = FLG_sen | FLG_zen | FLG_pen
)
