package test

import (
	"Go-SAP3/machine/types"
	"testing"
)

func TestCMP_A_0(t *testing.T) {
	source := `
	MVI A,0xf9
	MVI A,0xf9
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0xf9), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0xf9), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_A_1(t *testing.T) {
	source := `
	MVI A,0xce
	MVI A,0xce
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0xce), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0xce), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_A_2(t *testing.T) {
	source := `
	MVI A,0xcc
	MVI A,0xcc
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0xcc), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0xcc), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_A_3(t *testing.T) {
	source := `
	MVI A,0xee
	MVI A,0xee
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0xee), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0xee), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_A_4(t *testing.T) {
	source := `
	MVI A,0x9e
	MVI A,0x9e
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0x9e), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0x9e), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_A_5(t *testing.T) {
	source := `
	MVI A,0xa7
	MVI A,0xa7
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0xa7), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0xa7), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_A_6(t *testing.T) {
	source := `
	MVI A,0xe0
	MVI A,0xe0
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0xe0), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0xe0), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_A_7(t *testing.T) {
	source := `
	MVI A,0x88
	MVI A,0x88
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0x88), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0x88), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_A_8(t *testing.T) {
	source := `
	MVI A,0x38
	MVI A,0x38
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0x38), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0x38), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_A_9(t *testing.T) {
	source := `
	MVI A,0x6a
	MVI A,0x6a
	CMP A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP A", types.Word(0x6a), system.Accumulator.Value)
	check(t, "CMP A", types.Word(0x6a), system.Accumulator.Value)
	check(t, "CMP A", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_0(t *testing.T) {
	source := `
	MVI A,0xb3
	MVI B,0xa0
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0xb3), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0xa0), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_1(t *testing.T) {
	source := `
	MVI A,0x47
	MVI B,0x98
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0x47), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0x98), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_2(t *testing.T) {
	source := `
	MVI A,0x37
	MVI B,0xac
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0x37), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0xac), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_3(t *testing.T) {
	source := `
	MVI A,0x1a
	MVI B,0x34
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0x1a), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0x34), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_4(t *testing.T) {
	source := `
	MVI A,0xfc
	MVI B,0x29
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0xfc), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0x29), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_5(t *testing.T) {
	source := `
	MVI A,0xbc
	MVI B,0xc3
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0xbc), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0xc3), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_6(t *testing.T) {
	source := `
	MVI A,0xf
	MVI B,0xb2
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0xf), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0xb2), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_7(t *testing.T) {
	source := `
	MVI A,0xa
	MVI B,0x5a
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0xa), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0x5a), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_8(t *testing.T) {
	source := `
	MVI A,0x4a
	MVI B,0xec
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0x4a), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0xec), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_B_9(t *testing.T) {
	source := `
	MVI A,0x1f
	MVI B,0x9d
	CMP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP B", types.Word(0x1f), system.Accumulator.Value)
	check(t, "CMP B", types.Word(0x9d), system.BRegister.Value)
	check(t, "CMP B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_0(t *testing.T) {
	source := `
	MVI A,0xf
	MVI C,0xe5
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0xf), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0xe5), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_1(t *testing.T) {
	source := `
	MVI A,0x2f
	MVI C,0x56
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0x2f), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0x56), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_2(t *testing.T) {
	source := `
	MVI A,0x3f
	MVI C,0x7
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0x3f), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0x7), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_3(t *testing.T) {
	source := `
	MVI A,0x6b
	MVI C,0xbd
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0x6b), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0xbd), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_4(t *testing.T) {
	source := `
	MVI A,0xf0
	MVI C,0x90
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0xf0), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0x90), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_5(t *testing.T) {
	source := `
	MVI A,0x62
	MVI C,0x5
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0x62), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0x5), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_6(t *testing.T) {
	source := `
	MVI A,0xed
	MVI C,0xbc
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0xed), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0xbc), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_7(t *testing.T) {
	source := `
	MVI A,0x26
	MVI C,0xf6
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0x26), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0xf6), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_8(t *testing.T) {
	source := `
	MVI A,0xc0
	MVI C,0xc
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0xc0), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0xc), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_C_9(t *testing.T) {
	source := `
	MVI A,0xfc
	MVI C,0xcc
	CMP C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP C", types.Word(0xfc), system.Accumulator.Value)
	check(t, "CMP C", types.Word(0xcc), system.CRegister.Value)
	check(t, "CMP C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_0(t *testing.T) {
	source := `
	MVI A,0xc
	MVI D,0xa4
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0xc), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0xa4), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_1(t *testing.T) {
	source := `
	MVI A,0xfd
	MVI D,0x1a
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0xfd), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0x1a), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_2(t *testing.T) {
	source := `
	MVI A,0xb0
	MVI D,0x79
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0xb0), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0x79), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_3(t *testing.T) {
	source := `
	MVI A,0x38
	MVI D,0x4
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0x38), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0x4), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_4(t *testing.T) {
	source := `
	MVI A,0x18
	MVI D,0xc
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0x18), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0xc), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_5(t *testing.T) {
	source := `
	MVI A,0x76
	MVI D,0x1
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0x76), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0x1), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_6(t *testing.T) {
	source := `
	MVI A,0xd2
	MVI D,0x7d
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0xd2), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0x7d), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_7(t *testing.T) {
	source := `
	MVI A,0xf5
	MVI D,0x4a
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0xf5), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0x4a), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_8(t *testing.T) {
	source := `
	MVI A,0x95
	MVI D,0xc7
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0x95), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0xc7), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_D_9(t *testing.T) {
	source := `
	MVI A,0xe6
	MVI D,0xd0
	CMP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP D", types.Word(0xe6), system.Accumulator.Value)
	check(t, "CMP D", types.Word(0xd0), system.DRegister.Value)
	check(t, "CMP D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_0(t *testing.T) {
	source := `
	MVI A,0xfc
	MVI E,0x6c
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0xfc), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0x6c), system.ERegister.Value)
	check(t, "CMP E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_1(t *testing.T) {
	source := `
	MVI A,0x9
	MVI E,0x9
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0x9), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0x9), system.ERegister.Value)
	check(t, "CMP E", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_2(t *testing.T) {
	source := `
	MVI A,0xde
	MVI E,0x4c
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0xde), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0x4c), system.ERegister.Value)
	check(t, "CMP E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_3(t *testing.T) {
	source := `
	MVI A,0x91
	MVI E,0x36
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0x91), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0x36), system.ERegister.Value)
	check(t, "CMP E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_4(t *testing.T) {
	source := `
	MVI A,0xa0
	MVI E,0x4a
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0xa0), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0x4a), system.ERegister.Value)
	check(t, "CMP E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_5(t *testing.T) {
	source := `
	MVI A,0xeb
	MVI E,0x1c
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0xeb), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0x1c), system.ERegister.Value)
	check(t, "CMP E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_6(t *testing.T) {
	source := `
	MVI A,0x70
	MVI E,0x6a
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0x70), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0x6a), system.ERegister.Value)
	check(t, "CMP E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_7(t *testing.T) {
	source := `
	MVI A,0x1a
	MVI E,0x30
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0x1a), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0x30), system.ERegister.Value)
	check(t, "CMP E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_8(t *testing.T) {
	source := `
	MVI A,0xa1
	MVI E,0xbc
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0xa1), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0xbc), system.ERegister.Value)
	check(t, "CMP E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_E_9(t *testing.T) {
	source := `
	MVI A,0x80
	MVI E,0x2f
	CMP E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP E", types.Word(0x80), system.Accumulator.Value)
	check(t, "CMP E", types.Word(0x2f), system.ERegister.Value)
	check(t, "CMP E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_0(t *testing.T) {
	source := `
	MVI A,0x2d
	MVI H,0x83
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0x2d), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0x83), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_1(t *testing.T) {
	source := `
	MVI A,0x39
	MVI H,0xcf
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0x39), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0xcf), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_2(t *testing.T) {
	source := `
	MVI A,0xd9
	MVI H,0xb9
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0xd9), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0xb9), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_3(t *testing.T) {
	source := `
	MVI A,0xea
	MVI H,0x40
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0xea), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0x40), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_4(t *testing.T) {
	source := `
	MVI A,0x8e
	MVI H,0x34
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0x8e), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0x34), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_5(t *testing.T) {
	source := `
	MVI A,0x24
	MVI H,0xdd
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0x24), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0xdd), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_6(t *testing.T) {
	source := `
	MVI A,0xec
	MVI H,0x2f
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0xec), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0x2f), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_7(t *testing.T) {
	source := `
	MVI A,0x21
	MVI H,0x8e
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0x21), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0x8e), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_8(t *testing.T) {
	source := `
	MVI A,0x94
	MVI H,0x4d
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0x94), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0x4d), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_H_9(t *testing.T) {
	source := `
	MVI A,0xfe
	MVI H,0x39
	CMP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP H", types.Word(0xfe), system.Accumulator.Value)
	check(t, "CMP H", types.Word(0x39), system.HRegister.Value)
	check(t, "CMP H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_0(t *testing.T) {
	source := `
	MVI A,0xfe
	MVI L,0xc8
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0xfe), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0xc8), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_1(t *testing.T) {
	source := `
	MVI A,0x74
	MVI L,0x1a
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0x74), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0x1a), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_2(t *testing.T) {
	source := `
	MVI A,0x9e
	MVI L,0x16
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0x9e), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0x16), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_3(t *testing.T) {
	source := `
	MVI A,0xb3
	MVI L,0x4f
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0xb3), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0x4f), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_4(t *testing.T) {
	source := `
	MVI A,0x68
	MVI L,0x19
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0x68), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0x19), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_5(t *testing.T) {
	source := `
	MVI A,0xf6
	MVI L,0xdb
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0xf6), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0xdb), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_6(t *testing.T) {
	source := `
	MVI A,0xae
	MVI L,0xd9
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0xae), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0xd9), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_7(t *testing.T) {
	source := `
	MVI A,0xd1
	MVI L,0xf2
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0xd1), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0xf2), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_8(t *testing.T) {
	source := `
	MVI A,0x62
	MVI L,0x3d
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0x62), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0x3d), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_L_9(t *testing.T) {
	source := `
	MVI A,0x40
	MVI L,0x54
	CMP L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMP L", types.Word(0x40), system.Accumulator.Value)
	check(t, "CMP L", types.Word(0x54), system.LRegister.Value)
	check(t, "CMP L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCMP_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0x40
	CMP M
	HLT
	`

	value := types.Word(0x54)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "CMP M", types.Word(0x40), system.Accumulator.Value)
	check(t, "CMP M (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}