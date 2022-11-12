package op

// Control bit flags.
const (
	// 16KB Control ROM A
	Noop      = 0x0
	PC_re     = 0x1
	PC_we     = 0x2
	PC_ce     = 0x4
	MAR_re    = 0x8
	MDR_re    = 0x10
	MDR_we    = 0x20
	MAB_lsbre = 0x40
	MAB_lsbwe = 0x80
	MAB_msbre = 0x100
	MAB_msbwe = 0x200
	GPB_lsbre = 0x400
	GPB_lsbwe = 0x800
	GPB_msbre = 0x1000
	GPB_msbwe = 0x2000
	IR_re     = 0x4000

	// 16KB Control ROM B
	ACC_re     = 0x8000
	ACC_we     = 0x10000
	TMP_re     = 0x20000
	TMP_we     = 0x40000
	BREG_re    = 0x80000
	BREG_we    = 0x100000
	BREG_accen = 0x200000
	CREG_re    = 0x400000
	CREG_we    = 0x800000
	CREG_accen = 0x1000000
	DREG_re    = 0x2000000
	DREG_we    = 0x4000000
	DREG_accen = 0x8000000
	EREG_re    = 0x10000000
	EREG_we    = 0x20000000
	EREG_accen = 0x40000000

	// 16KB Control ROM C
	FREG_re    = 0x80000000
	FREG_we    = 0x100000000
	HREG_re    = 0x200000000
	HREG_we    = 0x400000000
	HREG_accen = 0x800000000
	LREG_re    = 0x1000000000
	LREG_we    = 0x2000000000
	LREG_accen = 0x4000000000
	SP_re      = 0x8000000000
	SP_we      = 0x10000000000
	SP_accen   = 0x20000000000
	ALU_we     = 0x40000000000
	ALU_b0     = 0x80000000000
	ALU_b1     = 0x100000000000
	ALU_b2     = 0x200000000000
	ALU_b3     = 0x400000000000

	// 16KB Control ROM D
	ALU_b4    = 0x800000000000
	FLG_csfie = 0x1000000000000
	FLG_csfne = 0x2000000000000
	FLG_czfie = 0x4000000000000
	FLG_czfne = 0x8000000000000
	FLG_ccfie = 0x10000000000000
	FLG_ccfne = 0x20000000000000
	FLG_cpfie = 0x40000000000000
	FLG_cpfne = 0x80000000000000
	FLG_cen   = 0x100000000000000
	FLG_zen   = 0x200000000000000
	FLG_sen   = 0x400000000000000
	FLG_pen   = 0x800000000000000
	IO_ps     = 0x1000000000000000
	IO_we     = 0x2000000000000000
	IO_re     = 0x4000000000000000
)

// Control bit flags for AluMode readability.
const (
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
