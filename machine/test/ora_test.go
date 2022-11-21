package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestORA_A_0(t *testing.T) {
	source := `
	MVI A,0xe1
	MVI A,0x60
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0x60), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_A_1(t *testing.T) {
	source := `
	MVI A,0xe4
	MVI A,0xb7
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0xb7), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_A_2(t *testing.T) {
	source := `
	MVI A,0xe8
	MVI A,0xf0
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0xf0), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_A_3(t *testing.T) {
	source := `
	MVI A,0x9
	MVI A,0x6f
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0x6f), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_A_4(t *testing.T) {
	source := `
	MVI A,0x3a
	MVI A,0xee
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0xee), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_A_5(t *testing.T) {
	source := `
	MVI A,0x7f
	MVI A,0xc9
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0xc9), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_A_6(t *testing.T) {
	source := `
	MVI A,0x5f
	MVI A,0x5c
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0x5c), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_A_7(t *testing.T) {
	source := `
	MVI A,0x1f
	MVI A,0x91
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0x91), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_A_8(t *testing.T) {
	source := `
	MVI A,0xe0
	MVI A,0xf7
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0xf7), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_A_9(t *testing.T) {
	source := `
	MVI A,0xa4
	MVI A,0x22
	ORA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA A", types.Word(0x22), system.Accumulator.Value)
	check(t, "ORA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_0(t *testing.T) {
	source := `
	MVI A,0x65
	MVI B,0xa7
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0xe7), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_1(t *testing.T) {
	source := `
	MVI A,0xb6
	MVI B,0x81
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0xb7), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_2(t *testing.T) {
	source := `
	MVI A,0x7
	MVI B,0x36
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0x37), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_3(t *testing.T) {
	source := `
	MVI A,0xa4
	MVI B,0xe
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0xae), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_4(t *testing.T) {
	source := `
	MVI A,0xa1
	MVI B,0x22
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0xa3), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_5(t *testing.T) {
	source := `
	MVI A,0xd4
	MVI B,0xc1
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0xd5), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_6(t *testing.T) {
	source := `
	MVI A,0x7f
	MVI B,0x2d
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0x7f), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_7(t *testing.T) {
	source := `
	MVI A,0x68
	MVI B,0xef
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0xef), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_8(t *testing.T) {
	source := `
	MVI A,0x7b
	MVI B,0x27
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0x7f), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_B_9(t *testing.T) {
	source := `
	MVI A,0xe7
	MVI B,0x50
	ORA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA B", types.Word(0xf7), system.Accumulator.Value)
	check(t, "ORA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_0(t *testing.T) {
	source := `
	MVI A,0x75
	MVI C,0xb9
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0xfd), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_1(t *testing.T) {
	source := `
	MVI A,0xdd
	MVI C,0xf2
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0xff), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_2(t *testing.T) {
	source := `
	MVI A,0x50
	MVI C,0xce
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0xde), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_3(t *testing.T) {
	source := `
	MVI A,0x19
	MVI C,0x76
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0x7f), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_4(t *testing.T) {
	source := `
	MVI A,0x83
	MVI C,0x13
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0x93), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_5(t *testing.T) {
	source := `
	MVI A,0x18
	MVI C,0xf7
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0xff), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_6(t *testing.T) {
	source := `
	MVI A,0xa9
	MVI C,0xc1
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0xe9), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_7(t *testing.T) {
	source := `
	MVI A,0x80
	MVI C,0xc2
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0xc2), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_8(t *testing.T) {
	source := `
	MVI A,0x99
	MVI C,0x46
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0xdf), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_C_9(t *testing.T) {
	source := `
	MVI A,0x93
	MVI C,0x3f
	ORA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA C", types.Word(0xbf), system.Accumulator.Value)
	check(t, "ORA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_0(t *testing.T) {
	source := `
	MVI A,0x9b
	MVI D,0xa5
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xbf), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_1(t *testing.T) {
	source := `
	MVI A,0xce
	MVI D,0xf2
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xfe), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_2(t *testing.T) {
	source := `
	MVI A,0x93
	MVI D,0xdd
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xdf), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_3(t *testing.T) {
	source := `
	MVI A,0xa9
	MVI D,0x83
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xab), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_4(t *testing.T) {
	source := `
	MVI A,0x6d
	MVI D,0x93
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xff), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_5(t *testing.T) {
	source := `
	MVI A,0x92
	MVI D,0xd4
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xd6), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_6(t *testing.T) {
	source := `
	MVI A,0x46
	MVI D,0xcd
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xcf), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_7(t *testing.T) {
	source := `
	MVI A,0x80
	MVI D,0xa1
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xa1), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_8(t *testing.T) {
	source := `
	MVI A,0xf0
	MVI D,0x1b
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xfb), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_D_9(t *testing.T) {
	source := `
	MVI A,0x92
	MVI D,0x75
	ORA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA D", types.Word(0xf7), system.Accumulator.Value)
	check(t, "ORA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_0(t *testing.T) {
	source := `
	MVI A,0xea
	MVI E,0x2e
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0xee), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_1(t *testing.T) {
	source := `
	MVI A,0xf5
	MVI E,0xd8
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0xfd), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_2(t *testing.T) {
	source := `
	MVI A,0xd2
	MVI E,0x4d
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0xdf), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_3(t *testing.T) {
	source := `
	MVI A,0xe2
	MVI E,0xb
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0xeb), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_4(t *testing.T) {
	source := `
	MVI A,0xf6
	MVI E,0x45
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0xf7), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_5(t *testing.T) {
	source := `
	MVI A,0x2c
	MVI E,0x1e
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0x3e), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_6(t *testing.T) {
	source := `
	MVI A,0xd8
	MVI E,0x28
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0xf8), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_7(t *testing.T) {
	source := `
	MVI A,0xc2
	MVI E,0x1e
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0xde), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_8(t *testing.T) {
	source := `
	MVI A,0x62
	MVI E,0xe4
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0xe6), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_E_9(t *testing.T) {
	source := `
	MVI A,0x56
	MVI E,0x72
	ORA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA E", types.Word(0x76), system.Accumulator.Value)
	check(t, "ORA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_0(t *testing.T) {
	source := `
	MVI A,0xd7
	MVI H,0x5e
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0xdf), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_1(t *testing.T) {
	source := `
	MVI A,0x1e
	MVI H,0x18
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0x1e), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_2(t *testing.T) {
	source := `
	MVI A,0x83
	MVI H,0x60
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0xe3), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_3(t *testing.T) {
	source := `
	MVI A,0x63
	MVI H,0xaf
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0xef), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_4(t *testing.T) {
	source := `
	MVI A,0x11
	MVI H,0xf
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0x1f), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_5(t *testing.T) {
	source := `
	MVI A,0xa0
	MVI H,0x83
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0xa3), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_6(t *testing.T) {
	source := `
	MVI A,0xea
	MVI H,0xbc
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0xfe), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_7(t *testing.T) {
	source := `
	MVI A,0x1f
	MVI H,0x2b
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0x3f), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_8(t *testing.T) {
	source := `
	MVI A,0x4f
	MVI H,0x4f
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0x4f), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_H_9(t *testing.T) {
	source := `
	MVI A,0xb3
	MVI H,0xb4
	ORA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA H", types.Word(0xb7), system.Accumulator.Value)
	check(t, "ORA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_0(t *testing.T) {
	source := `
	MVI A,0x23
	MVI L,0xf1
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0xf3), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_1(t *testing.T) {
	source := `
	MVI A,0x72
	MVI L,0xf
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0x7f), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_2(t *testing.T) {
	source := `
	MVI A,0x93
	MVI L,0x61
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0xf3), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_3(t *testing.T) {
	source := `
	MVI A,0xf8
	MVI L,0xc4
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0xfc), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_4(t *testing.T) {
	source := `
	MVI A,0x9c
	MVI L,0x8c
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0x9c), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_5(t *testing.T) {
	source := `
	MVI A,0x3d
	MVI L,0xd
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0x3d), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_6(t *testing.T) {
	source := `
	MVI A,0x7d
	MVI L,0xe
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0x7f), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_7(t *testing.T) {
	source := `
	MVI A,0x1b
	MVI L,0x7d
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0x7f), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_8(t *testing.T) {
	source := `
	MVI A,0x2a
	MVI L,0xe6
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0xee), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_L_9(t *testing.T) {
	source := `
	MVI A,0x3f
	MVI L,0x3a
	ORA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORA L", types.Word(0x3f), system.Accumulator.Value)
	check(t, "ORA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestORA_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0x3f
	ORA M
	HLT
	`

	value := types.Word(0x3a)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "ORA M", types.Word(0x3f), system.Accumulator.Value)
	check(t, "ORA M (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}