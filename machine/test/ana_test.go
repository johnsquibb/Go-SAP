package test

import (
	"Go-SAP3/machine/types"
	"testing"
)

func TestANA_A_0(t *testing.T) {
	source := `
	MVI A,0xcf
	MVI A,0x5b
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0x5b), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_A_1(t *testing.T) {
	source := `
	MVI A,0x3a
	MVI A,0x22
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0x22), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_A_2(t *testing.T) {
	source := `
	MVI A,0x68
	MVI A,0xbd
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0xbd), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_A_3(t *testing.T) {
	source := `
	MVI A,0x41
	MVI A,0xeb
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0xeb), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_A_4(t *testing.T) {
	source := `
	MVI A,0x29
	MVI A,0x83
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0x83), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_A_5(t *testing.T) {
	source := `
	MVI A,0x51
	MVI A,0x1c
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0x1c), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_A_6(t *testing.T) {
	source := `
	MVI A,0x59
	MVI A,0xb5
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0xb5), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_A_7(t *testing.T) {
	source := `
	MVI A,0xe6
	MVI A,0x48
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0x48), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_A_8(t *testing.T) {
	source := `
	MVI A,0x16
	MVI A,0x31
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0x31), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_A_9(t *testing.T) {
	source := `
	MVI A,0x5a
	MVI A,0x2a
	ANA A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA A", types.Word(0x2a), system.Accumulator.Value)
	check(t, "ANA A", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_0(t *testing.T) {
	source := `
	MVI A,0xfe
	MVI B,0xec
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0xec), system.Accumulator.Value)
	check(t, "ANA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_1(t *testing.T) {
	source := `
	MVI A,0xe8
	MVI B,0x1f
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0x8), system.Accumulator.Value)
	check(t, "ANA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_2(t *testing.T) {
	source := `
	MVI A,0xb9
	MVI B,0xe4
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0xa0), system.Accumulator.Value)
	check(t, "ANA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_3(t *testing.T) {
	source := `
	MVI A,0xce
	MVI B,0x1
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0x0), system.Accumulator.Value)
	check(t, "ANA B", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_4(t *testing.T) {
	source := `
	MVI A,0x46
	MVI B,0xf3
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0x42), system.Accumulator.Value)
	check(t, "ANA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_5(t *testing.T) {
	source := `
	MVI A,0x2
	MVI B,0xbf
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0x2), system.Accumulator.Value)
	check(t, "ANA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_6(t *testing.T) {
	source := `
	MVI A,0xc9
	MVI B,0x91
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0x81), system.Accumulator.Value)
	check(t, "ANA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_7(t *testing.T) {
	source := `
	MVI A,0xf2
	MVI B,0x1c
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0x10), system.Accumulator.Value)
	check(t, "ANA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_8(t *testing.T) {
	source := `
	MVI A,0xe8
	MVI B,0xb4
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0xa0), system.Accumulator.Value)
	check(t, "ANA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_B_9(t *testing.T) {
	source := `
	MVI A,0xfe
	MVI B,0xa4
	ANA B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA B", types.Word(0xa4), system.Accumulator.Value)
	check(t, "ANA B", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_0(t *testing.T) {
	source := `
	MVI A,0x7e
	MVI C,0xfa
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x7a), system.Accumulator.Value)
	check(t, "ANA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_1(t *testing.T) {
	source := `
	MVI A,0x2
	MVI C,0x49
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x0), system.Accumulator.Value)
	check(t, "ANA C", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_2(t *testing.T) {
	source := `
	MVI A,0x77
	MVI C,0x92
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x12), system.Accumulator.Value)
	check(t, "ANA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_3(t *testing.T) {
	source := `
	MVI A,0x46
	MVI C,0x3c
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x4), system.Accumulator.Value)
	check(t, "ANA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_4(t *testing.T) {
	source := `
	MVI A,0x25
	MVI C,0x67
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x25), system.Accumulator.Value)
	check(t, "ANA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_5(t *testing.T) {
	source := `
	MVI A,0x9
	MVI C,0x64
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x0), system.Accumulator.Value)
	check(t, "ANA C", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_6(t *testing.T) {
	source := `
	MVI A,0xf3
	MVI C,0x4d
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x41), system.Accumulator.Value)
	check(t, "ANA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_7(t *testing.T) {
	source := `
	MVI A,0xc9
	MVI C,0x7b
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x49), system.Accumulator.Value)
	check(t, "ANA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_8(t *testing.T) {
	source := `
	MVI A,0xc
	MVI C,0x6b
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x8), system.Accumulator.Value)
	check(t, "ANA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_C_9(t *testing.T) {
	source := `
	MVI A,0xd
	MVI C,0x5
	ANA C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA C", types.Word(0x5), system.Accumulator.Value)
	check(t, "ANA C", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_0(t *testing.T) {
	source := `
	MVI A,0x53
	MVI D,0x26
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0x2), system.Accumulator.Value)
	check(t, "ANA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_1(t *testing.T) {
	source := `
	MVI A,0xf4
	MVI D,0x40
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0x40), system.Accumulator.Value)
	check(t, "ANA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_2(t *testing.T) {
	source := `
	MVI A,0x7c
	MVI D,0xe1
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0x60), system.Accumulator.Value)
	check(t, "ANA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_3(t *testing.T) {
	source := `
	MVI A,0xf6
	MVI D,0x3b
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0x32), system.Accumulator.Value)
	check(t, "ANA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_4(t *testing.T) {
	source := `
	MVI A,0xeb
	MVI D,0xeb
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0xeb), system.Accumulator.Value)
	check(t, "ANA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_5(t *testing.T) {
	source := `
	MVI A,0xa0
	MVI D,0xe7
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0xa0), system.Accumulator.Value)
	check(t, "ANA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_6(t *testing.T) {
	source := `
	MVI A,0x64
	MVI D,0x16
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0x4), system.Accumulator.Value)
	check(t, "ANA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_7(t *testing.T) {
	source := `
	MVI A,0x79
	MVI D,0x5d
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0x59), system.Accumulator.Value)
	check(t, "ANA D", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_8(t *testing.T) {
	source := `
	MVI A,0x2
	MVI D,0xa9
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0x0), system.Accumulator.Value)
	check(t, "ANA D", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_D_9(t *testing.T) {
	source := `
	MVI A,0x88
	MVI D,0x66
	ANA D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA D", types.Word(0x0), system.Accumulator.Value)
	check(t, "ANA D", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_0(t *testing.T) {
	source := `
	MVI A,0xe0
	MVI E,0x62
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0x60), system.Accumulator.Value)
	check(t, "ANA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_1(t *testing.T) {
	source := `
	MVI A,0x66
	MVI E,0x8d
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0x4), system.Accumulator.Value)
	check(t, "ANA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_2(t *testing.T) {
	source := `
	MVI A,0xab
	MVI E,0x96
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0x82), system.Accumulator.Value)
	check(t, "ANA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_3(t *testing.T) {
	source := `
	MVI A,0x3b
	MVI E,0x34
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0x30), system.Accumulator.Value)
	check(t, "ANA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_4(t *testing.T) {
	source := `
	MVI A,0x97
	MVI E,0x17
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0x17), system.Accumulator.Value)
	check(t, "ANA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_5(t *testing.T) {
	source := `
	MVI A,0x2d
	MVI E,0x59
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0x9), system.Accumulator.Value)
	check(t, "ANA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_6(t *testing.T) {
	source := `
	MVI A,0xbf
	MVI E,0x64
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0x24), system.Accumulator.Value)
	check(t, "ANA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_7(t *testing.T) {
	source := `
	MVI A,0x48
	MVI E,0x6e
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0x48), system.Accumulator.Value)
	check(t, "ANA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_8(t *testing.T) {
	source := `
	MVI A,0xe7
	MVI E,0xac
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0xa4), system.Accumulator.Value)
	check(t, "ANA E", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_E_9(t *testing.T) {
	source := `
	MVI A,0x6d
	MVI E,0x92
	ANA E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA E", types.Word(0x0), system.Accumulator.Value)
	check(t, "ANA E", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_0(t *testing.T) {
	source := `
	MVI A,0xc7
	MVI H,0xd
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0x5), system.Accumulator.Value)
	check(t, "ANA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_1(t *testing.T) {
	source := `
	MVI A,0xad
	MVI H,0x6e
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0x2c), system.Accumulator.Value)
	check(t, "ANA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_2(t *testing.T) {
	source := `
	MVI A,0xb6
	MVI H,0xf9
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0xb0), system.Accumulator.Value)
	check(t, "ANA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_3(t *testing.T) {
	source := `
	MVI A,0x1e
	MVI H,0x1
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0x0), system.Accumulator.Value)
	check(t, "ANA H", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_4(t *testing.T) {
	source := `
	MVI A,0x64
	MVI H,0xb1
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0x20), system.Accumulator.Value)
	check(t, "ANA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_5(t *testing.T) {
	source := `
	MVI A,0xd2
	MVI H,0xd7
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0xd2), system.Accumulator.Value)
	check(t, "ANA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_6(t *testing.T) {
	source := `
	MVI A,0x61
	MVI H,0xf7
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0x61), system.Accumulator.Value)
	check(t, "ANA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_7(t *testing.T) {
	source := `
	MVI A,0x38
	MVI H,0x4
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0x0), system.Accumulator.Value)
	check(t, "ANA H", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_8(t *testing.T) {
	source := `
	MVI A,0x39
	MVI H,0x51
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0x11), system.Accumulator.Value)
	check(t, "ANA H", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_H_9(t *testing.T) {
	source := `
	MVI A,0x56
	MVI H,0x1
	ANA H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA H", types.Word(0x0), system.Accumulator.Value)
	check(t, "ANA H", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_0(t *testing.T) {
	source := `
	MVI A,0x51
	MVI L,0x52
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0x50), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_1(t *testing.T) {
	source := `
	MVI A,0x2e
	MVI L,0xfe
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0x2e), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_2(t *testing.T) {
	source := `
	MVI A,0x22
	MVI L,0xff
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0x22), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_3(t *testing.T) {
	source := `
	MVI A,0x94
	MVI L,0xa3
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0x80), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_4(t *testing.T) {
	source := `
	MVI A,0x7b
	MVI L,0xdc
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0x58), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_5(t *testing.T) {
	source := `
	MVI A,0x4b
	MVI L,0x42
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0x42), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_6(t *testing.T) {
	source := `
	MVI A,0x63
	MVI L,0x35
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0x21), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_7(t *testing.T) {
	source := `
	MVI A,0xcf
	MVI L,0xc9
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0xc9), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_8(t *testing.T) {
	source := `
	MVI A,0x4e
	MVI L,0xce
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0x4e), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_L_9(t *testing.T) {
	source := `
	MVI A,0x6d
	MVI L,0x7d
	ANA L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANA L", types.Word(0x6d), system.Accumulator.Value)
	check(t, "ANA L", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestANA_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0x6d
	ANA M
	HLT
	`

	value := types.Word(0x7d)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "ANA M", types.Word(0x6d), system.Accumulator.Value)
	check(t, "ANA M (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}


