package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestXRA_A_0(t *testing.T) {
	source := `
	MVI A,0x6e
	MVI A,0x7b
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0x7b), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_A_1(t *testing.T) {
	source := `
	MVI A,0x9f
	MVI A,0x52
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0x52), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_A_2(t *testing.T) {
	source := `
	MVI A,0xc
	MVI A,0xb2
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0xb2), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_A_3(t *testing.T) {
	source := `
	MVI A,0x3d
	MVI A,0x1e
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0x1e), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_A_4(t *testing.T) {
	source := `
	MVI A,0xa5
	MVI A,0xad
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0xad), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_A_5(t *testing.T) {
	source := `
	MVI A,0xa1
	MVI A,0x7f
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0x7f), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_A_6(t *testing.T) {
	source := `
	MVI A,0xfc
	MVI A,0x6a
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0x6a), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_A_7(t *testing.T) {
	source := `
	MVI A,0xd4
	MVI A,0x55
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0x55), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_A_8(t *testing.T) {
	source := `
	MVI A,0x85
	MVI A,0xe6
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0xe6), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_A_9(t *testing.T) {
	source := `
	MVI A,0xfa
	MVI A,0x4e
	XRA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA A", types.Word(0x4e), system.Accumulator.Value)
	check(t, "XRA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_0(t *testing.T) {
	source := `
	MVI A,0xc6
	MVI B,0xba
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0x7c), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_1(t *testing.T) {
	source := `
	MVI A,0x10
	MVI B,0xcc
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0xdc), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_2(t *testing.T) {
	source := `
	MVI A,0xf8
	MVI B,0x49
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0xb1), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_3(t *testing.T) {
	source := `
	MVI A,0x7f
	MVI B,0x60
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0x1f), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_4(t *testing.T) {
	source := `
	MVI A,0x39
	MVI B,0x8a
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0xb3), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_5(t *testing.T) {
	source := `
	MVI A,0x37
	MVI B,0xce
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0xf9), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_6(t *testing.T) {
	source := `
	MVI A,0xa7
	MVI B,0xb6
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0x11), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_7(t *testing.T) {
	source := `
	MVI A,0x40
	MVI B,0xe8
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0xa8), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_8(t *testing.T) {
	source := `
	MVI A,0xab
	MVI B,0x60
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0xcb), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_B_9(t *testing.T) {
	source := `
	MVI A,0x54
	MVI B,0x4b
	XRA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA B", types.Word(0x1f), system.Accumulator.Value)
	check(t, "XRA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_0(t *testing.T) {
	source := `
	MVI A,0x31
	MVI C,0x56
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0x67), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_1(t *testing.T) {
	source := `
	MVI A,0x99
	MVI C,0x70
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0xe9), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_2(t *testing.T) {
	source := `
	MVI A,0x92
	MVI C,0xca
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0x58), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_3(t *testing.T) {
	source := `
	MVI A,0x6a
	MVI C,0xbd
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0xd7), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_4(t *testing.T) {
	source := `
	MVI A,0x60
	MVI C,0xa0
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0xc0), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_5(t *testing.T) {
	source := `
	MVI A,0xe7
	MVI C,0xf5
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0x12), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_6(t *testing.T) {
	source := `
	MVI A,0xad
	MVI C,0xce
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0x63), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_7(t *testing.T) {
	source := `
	MVI A,0x5d
	MVI C,0x51
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0xc), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_8(t *testing.T) {
	source := `
	MVI A,0x12
	MVI C,0x78
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0x6a), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_C_9(t *testing.T) {
	source := `
	MVI A,0x83
	MVI C,0x6b
	XRA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA C", types.Word(0xe8), system.Accumulator.Value)
	check(t, "XRA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_0(t *testing.T) {
	source := `
	MVI A,0x21
	MVI D,0x27
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0x6), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_1(t *testing.T) {
	source := `
	MVI A,0x50
	MVI D,0xd8
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0x88), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_2(t *testing.T) {
	source := `
	MVI A,0x4
	MVI D,0xb
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0xf), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_3(t *testing.T) {
	source := `
	MVI A,0xb1
	MVI D,0x76
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0xc7), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_4(t *testing.T) {
	source := `
	MVI A,0x70
	MVI D,0x9e
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0xee), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_5(t *testing.T) {
	source := `
	MVI A,0x18
	MVI D,0x93
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0x8b), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_6(t *testing.T) {
	source := `
	MVI A,0x37
	MVI D,0xb
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0x3c), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_7(t *testing.T) {
	source := `
	MVI A,0x5f
	MVI D,0x3b
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0x64), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_8(t *testing.T) {
	source := `
	MVI A,0x9d
	MVI D,0x58
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0xc5), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_D_9(t *testing.T) {
	source := `
	MVI A,0xc2
	MVI D,0xe0
	XRA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA D", types.Word(0x22), system.Accumulator.Value)
	check(t, "XRA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_0(t *testing.T) {
	source := `
	MVI A,0xec
	MVI E,0xce
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0x22), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_1(t *testing.T) {
	source := `
	MVI A,0x9b
	MVI E,0x12
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0x89), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_2(t *testing.T) {
	source := `
	MVI A,0x36
	MVI E,0xa
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0x3c), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_3(t *testing.T) {
	source := `
	MVI A,0x45
	MVI E,0xa8
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0xed), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_4(t *testing.T) {
	source := `
	MVI A,0x65
	MVI E,0xf
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0x6a), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_5(t *testing.T) {
	source := `
	MVI A,0xf
	MVI E,0x19
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0x16), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_6(t *testing.T) {
	source := `
	MVI A,0x12
	MVI E,0xdd
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0xcf), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_7(t *testing.T) {
	source := `
	MVI A,0xa
	MVI E,0xe
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0x4), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_8(t *testing.T) {
	source := `
	MVI A,0x1b
	MVI E,0x72
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0x69), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_E_9(t *testing.T) {
	source := `
	MVI A,0xc8
	MVI E,0x52
	XRA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA E", types.Word(0x9a), system.Accumulator.Value)
	check(t, "XRA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_0(t *testing.T) {
	source := `
	MVI A,0xad
	MVI H,0x3b
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0x96), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_1(t *testing.T) {
	source := `
	MVI A,0xbb
	MVI H,0xc9
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0x72), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_2(t *testing.T) {
	source := `
	MVI A,0x6f
	MVI H,0xef
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0x80), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_3(t *testing.T) {
	source := `
	MVI A,0x29
	MVI H,0x8b
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0xa2), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_4(t *testing.T) {
	source := `
	MVI A,0x7b
	MVI H,0xcf
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0xb4), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_5(t *testing.T) {
	source := `
	MVI A,0x39
	MVI H,0xe1
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0xd8), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_6(t *testing.T) {
	source := `
	MVI A,0x24
	MVI H,0x36
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0x12), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_7(t *testing.T) {
	source := `
	MVI A,0x99
	MVI H,0x25
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0xbc), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_8(t *testing.T) {
	source := `
	MVI A,0xfb
	MVI H,0xc3
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0x38), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_H_9(t *testing.T) {
	source := `
	MVI A,0x65
	MVI H,0x53
	XRA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA H", types.Word(0x36), system.Accumulator.Value)
	check(t, "XRA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_0(t *testing.T) {
	source := `
	MVI A,0x51
	MVI L,0xd0
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0x81), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_1(t *testing.T) {
	source := `
	MVI A,0xe1
	MVI L,0x2
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0xe3), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_2(t *testing.T) {
	source := `
	MVI A,0x1a
	MVI L,0x75
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0x6f), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_3(t *testing.T) {
	source := `
	MVI A,0x28
	MVI L,0xf1
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0xd9), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_4(t *testing.T) {
	source := `
	MVI A,0xb8
	MVI L,0xd1
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0x69), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_5(t *testing.T) {
	source := `
	MVI A,0x59
	MVI L,0x47
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0x1e), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_6(t *testing.T) {
	source := `
	MVI A,0x23
	MVI L,0x74
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0x57), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_7(t *testing.T) {
	source := `
	MVI A,0x30
	MVI L,0xa7
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0x97), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_8(t *testing.T) {
	source := `
	MVI A,0x32
	MVI L,0x57
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0x65), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_L_9(t *testing.T) {
	source := `
	MVI A,0xc9
	MVI L,0xf9
	XRA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRA L", types.Word(0x30), system.Accumulator.Value)
	check(t, "XRA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestXRA_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0xc9
	XRA M
	HLT
	`

	value := types.Word(0xf9)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "XRA M", types.Word(0x30), system.Accumulator.Value)
	check(t, "XRA M (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}