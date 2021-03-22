package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestMOV_A_A(t *testing.T) {
	source := `
	MVI A,0x88
	MOV A,A
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV A,A", types.Word(136), system.Accumulator.Value)
	check(t, "MOV A,A", types.Word(136), system.Accumulator.Value)
}

func TestMOV_A_B(t *testing.T) {
	source := `
	MVI B,0x7f
	MOV A,B
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV A,B", types.Word(127), system.Accumulator.Value)
	check(t, "MOV A,B", types.Word(127), system.BRegister.Value)
}

func TestMOV_A_C(t *testing.T) {
	source := `
	MVI C,0xd1
	MOV A,C
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV A,C", types.Word(209), system.Accumulator.Value)
	check(t, "MOV A,C", types.Word(209), system.CRegister.Value)
}

func TestMOV_A_D(t *testing.T) {
	source := `
	MVI D,0x48
	MOV A,D
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV A,D", types.Word(72), system.Accumulator.Value)
	check(t, "MOV A,D", types.Word(72), system.DRegister.Value)
}

func TestMOV_A_E(t *testing.T) {
	source := `
	MVI E,0x48
	MOV A,E
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV A,E", types.Word(72), system.Accumulator.Value)
	check(t, "MOV A,E", types.Word(72), system.ERegister.Value)
}

func TestMOV_A_H(t *testing.T) {
	source := `
	MVI H,0xd0
	MOV A,H
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV A,H", types.Word(208), system.Accumulator.Value)
	check(t, "MOV A,H", types.Word(208), system.HRegister.Value)
}

func TestMOV_A_L(t *testing.T) {
	source := `
	MVI L,0x79
	MOV A,L
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV A,L", types.Word(121), system.Accumulator.Value)
	check(t, "MOV A,L", types.Word(121), system.LRegister.Value)
}

func TestMOV_A_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MOV A,M
	HLT
	`
	value := types.Word(0xeb)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "MOV A,M", value, system.Accumulator.Value)
}

func TestMOV_M_A(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0x1b
	MOV M,A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "MOV M,A", types.Word(0x1b), system.RandomAccessMemory.Values[0x1234])
}

func TestMOV_B_A(t *testing.T) {
	source := `
	MVI A,0xf
	MOV B,A
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV B,A", types.Word(15), system.BRegister.Value)
	check(t, "MOV B,A", types.Word(15), system.Accumulator.Value)
}

func TestMOV_B_B(t *testing.T) {
	source := `
	MVI B,0x93
	MOV B,B
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV B,B", types.Word(147), system.BRegister.Value)
	check(t, "MOV B,B", types.Word(147), system.BRegister.Value)
}

func TestMOV_B_C(t *testing.T) {
	source := `
	MVI C,0x62
	MOV B,C
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV B,C", types.Word(98), system.BRegister.Value)
	check(t, "MOV B,C", types.Word(98), system.CRegister.Value)
}

func TestMOV_B_D(t *testing.T) {
	source := `
	MVI D,0x19
	MOV B,D
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV B,D", types.Word(25), system.BRegister.Value)
	check(t, "MOV B,D", types.Word(25), system.DRegister.Value)
}

func TestMOV_B_E(t *testing.T) {
	source := `
	MVI E,0xd6
	MOV B,E
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV B,E", types.Word(214), system.BRegister.Value)
	check(t, "MOV B,E", types.Word(214), system.ERegister.Value)
}

func TestMOV_B_H(t *testing.T) {
	source := `
	MVI H,0xe4
	MOV B,H
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV B,H", types.Word(228), system.BRegister.Value)
	check(t, "MOV B,H", types.Word(228), system.HRegister.Value)
}

func TestMOV_B_L(t *testing.T) {
	source := `
	MVI L,0x8d
	MOV B,L
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV B,L", types.Word(141), system.BRegister.Value)
	check(t, "MOV B,L", types.Word(141), system.LRegister.Value)
}

func TestMOV_B_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MOV B,M
	HLT
	`
	value := types.Word(0xd7)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "MOV B,M", value, system.BRegister.Value)
}

func TestMOV_M_B(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI B,0xe8
	MOV M,B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "MOV M,B", types.Word(0xe8), system.RandomAccessMemory.Values[0x1234])
}

func TestMOV_C_A(t *testing.T) {
	source := `
	MVI A,0x71
	MOV C,A
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV C,A", types.Word(113), system.CRegister.Value)
	check(t, "MOV C,A", types.Word(113), system.Accumulator.Value)
}

func TestMOV_C_B(t *testing.T) {
	source := `
	MVI B,0x5e
	MOV C,B
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV C,B", types.Word(94), system.CRegister.Value)
	check(t, "MOV C,B", types.Word(94), system.BRegister.Value)
}

func TestMOV_C_C(t *testing.T) {
	source := `
	MVI C,0x78
	MOV C,C
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV C,C", types.Word(120), system.CRegister.Value)
	check(t, "MOV C,C", types.Word(120), system.CRegister.Value)
}

func TestMOV_C_D(t *testing.T) {
	source := `
	MVI D,0x8a
	MOV C,D
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV C,D", types.Word(138), system.CRegister.Value)
	check(t, "MOV C,D", types.Word(138), system.DRegister.Value)
}

func TestMOV_C_E(t *testing.T) {
	source := `
	MVI E,0xb5
	MOV C,E
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV C,E", types.Word(181), system.CRegister.Value)
	check(t, "MOV C,E", types.Word(181), system.ERegister.Value)
}

func TestMOV_C_H(t *testing.T) {
	source := `
	MVI H,0x8c
	MOV C,H
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV C,H", types.Word(140), system.CRegister.Value)
	check(t, "MOV C,H", types.Word(140), system.HRegister.Value)
}

func TestMOV_C_L(t *testing.T) {
	source := `
	MVI L,0x96
	MOV C,L
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV C,L", types.Word(150), system.CRegister.Value)
	check(t, "MOV C,L", types.Word(150), system.LRegister.Value)
}

func TestMOV_C_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MOV C,M
	HLT
	`
	value := types.Word(0xf)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "MOV C,M", value, system.CRegister.Value)
}

func TestMOV_M_C(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI C,0xf0
	MOV M,C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "MOV M,C", types.Word(0xf0), system.RandomAccessMemory.Values[0x1234])
}

func TestMOV_D_A(t *testing.T) {
	source := `
	MVI A,0xac
	MOV D,A
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV D,A", types.Word(172), system.DRegister.Value)
	check(t, "MOV D,A", types.Word(172), system.Accumulator.Value)
}

func TestMOV_D_B(t *testing.T) {
	source := `
	MVI B,0x21
	MOV D,B
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV D,B", types.Word(33), system.DRegister.Value)
	check(t, "MOV D,B", types.Word(33), system.BRegister.Value)
}

func TestMOV_D_C(t *testing.T) {
	source := `
	MVI C,0x18
	MOV D,C
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV D,C", types.Word(24), system.DRegister.Value)
	check(t, "MOV D,C", types.Word(24), system.CRegister.Value)
}

func TestMOV_D_D(t *testing.T) {
	source := `
	MVI D,0x65
	MOV D,D
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV D,D", types.Word(101), system.DRegister.Value)
	check(t, "MOV D,D", types.Word(101), system.DRegister.Value)
}

func TestMOV_D_E(t *testing.T) {
	source := `
	MVI E,0xfe
	MOV D,E
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV D,E", types.Word(254), system.DRegister.Value)
	check(t, "MOV D,E", types.Word(254), system.ERegister.Value)
}

func TestMOV_D_H(t *testing.T) {
	source := `
	MVI H,0x54
	MOV D,H
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV D,H", types.Word(84), system.DRegister.Value)
	check(t, "MOV D,H", types.Word(84), system.HRegister.Value)
}

func TestMOV_D_L(t *testing.T) {
	source := `
	MVI L,0xc5
	MOV D,L
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV D,L", types.Word(197), system.DRegister.Value)
	check(t, "MOV D,L", types.Word(197), system.LRegister.Value)
}

func TestMOV_D_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MOV D,M
	HLT
	`
	value := types.Word(0x47)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "MOV D,M", value, system.DRegister.Value)
}

func TestMOV_M_D(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI D,0xcf
	MOV M,D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "MOV M,D", types.Word(0xcf), system.RandomAccessMemory.Values[0x1234])
}

func TestMOV_E_A(t *testing.T) {
	source := `
	MVI A,0x63
	MOV E,A
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV E,A", types.Word(99), system.ERegister.Value)
	check(t, "MOV E,A", types.Word(99), system.Accumulator.Value)
}

func TestMOV_E_B(t *testing.T) {
	source := `
	MVI B,0xfa
	MOV E,B
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV E,B", types.Word(250), system.ERegister.Value)
	check(t, "MOV E,B", types.Word(250), system.BRegister.Value)
}

func TestMOV_E_C(t *testing.T) {
	source := `
	MVI C,0x1d
	MOV E,C
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV E,C", types.Word(29), system.ERegister.Value)
	check(t, "MOV E,C", types.Word(29), system.CRegister.Value)
}

func TestMOV_E_D(t *testing.T) {
	source := `
	MVI D,0x4c
	MOV E,D
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV E,D", types.Word(76), system.ERegister.Value)
	check(t, "MOV E,D", types.Word(76), system.DRegister.Value)
}

func TestMOV_E_E(t *testing.T) {
	source := `
	MVI E,0xe9
	MOV E,E
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV E,E", types.Word(233), system.ERegister.Value)
	check(t, "MOV E,E", types.Word(233), system.ERegister.Value)
}

func TestMOV_E_H(t *testing.T) {
	source := `
	MVI H,0x16
	MOV E,H
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV E,H", types.Word(22), system.ERegister.Value)
	check(t, "MOV E,H", types.Word(22), system.HRegister.Value)
}

func TestMOV_E_L(t *testing.T) {
	source := `
	MVI L,0x56
	MOV E,L
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV E,L", types.Word(86), system.ERegister.Value)
	check(t, "MOV E,L", types.Word(86), system.LRegister.Value)
}

func TestMOV_E_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MOV E,M
	HLT
	`
	value := types.Word(0x7b)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "MOV E,M", value, system.ERegister.Value)
}

func TestMOV_M_E(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI E,0x1d
	MOV M,E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "MOV M,E", types.Word(0x1d), system.RandomAccessMemory.Values[0x1234])
}

func TestMOV_H_A(t *testing.T) {
	source := `
	MVI A,0xf7
	MOV H,A
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV H,A", types.Word(247), system.HRegister.Value)
	check(t, "MOV H,A", types.Word(247), system.Accumulator.Value)
}

func TestMOV_H_B(t *testing.T) {
	source := `
	MVI B,0x91
	MOV H,B
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV H,B", types.Word(145), system.HRegister.Value)
	check(t, "MOV H,B", types.Word(145), system.BRegister.Value)
}

func TestMOV_H_C(t *testing.T) {
	source := `
	MVI C,0x33
	MOV H,C
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV H,C", types.Word(51), system.HRegister.Value)
	check(t, "MOV H,C", types.Word(51), system.CRegister.Value)
}

func TestMOV_H_D(t *testing.T) {
	source := `
	MVI D,0x6c
	MOV H,D
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV H,D", types.Word(108), system.HRegister.Value)
	check(t, "MOV H,D", types.Word(108), system.DRegister.Value)
}

func TestMOV_H_E(t *testing.T) {
	source := `
	MVI E,0xbf
	MOV H,E
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV H,E", types.Word(191), system.HRegister.Value)
	check(t, "MOV H,E", types.Word(191), system.ERegister.Value)
}

func TestMOV_H_H(t *testing.T) {
	source := `
	MVI H,0x79
	MOV H,H
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV H,H", types.Word(121), system.HRegister.Value)
	check(t, "MOV H,H", types.Word(121), system.HRegister.Value)
}

func TestMOV_H_L(t *testing.T) {
	source := `
	MVI L,0xb
	MOV H,L
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV H,L", types.Word(11), system.HRegister.Value)
	check(t, "MOV H,L", types.Word(11), system.LRegister.Value)
}

func TestMOV_H_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MOV H,M
	HLT
	`
	value := types.Word(0x20)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "MOV H,M", value, system.HRegister.Value)
}

func TestMOV_M_H(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MOV M,H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "MOV M,H", types.Word(0x12), system.RandomAccessMemory.Values[0x1234])
}

func TestMOV_L_A(t *testing.T) {
	source := `
	MVI A,0x83
	MOV L,A
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV L,A", types.Word(131), system.LRegister.Value)
	check(t, "MOV L,A", types.Word(131), system.Accumulator.Value)
}

func TestMOV_L_B(t *testing.T) {
	source := `
	MVI B,0x94
	MOV L,B
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV L,B", types.Word(148), system.LRegister.Value)
	check(t, "MOV L,B", types.Word(148), system.BRegister.Value)
}

func TestMOV_L_C(t *testing.T) {
	source := `
	MVI C,0xd4
	MOV L,C
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV L,C", types.Word(212), system.LRegister.Value)
	check(t, "MOV L,C", types.Word(212), system.CRegister.Value)
}

func TestMOV_L_D(t *testing.T) {
	source := `
	MVI D,0x26
	MOV L,D
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV L,D", types.Word(38), system.LRegister.Value)
	check(t, "MOV L,D", types.Word(38), system.DRegister.Value)
}

func TestMOV_L_E(t *testing.T) {
	source := `
	MVI E,0x7e
	MOV L,E
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV L,E", types.Word(126), system.LRegister.Value)
	check(t, "MOV L,E", types.Word(126), system.ERegister.Value)
}

func TestMOV_L_H(t *testing.T) {
	source := `
	MVI H,0x91
	MOV L,H
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV L,H", types.Word(145), system.LRegister.Value)
	check(t, "MOV L,H", types.Word(145), system.HRegister.Value)
}

func TestMOV_L_L(t *testing.T) {
	source := `
	MVI L,0x98
	MOV L,L
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MOV L,L", types.Word(152), system.LRegister.Value)
	check(t, "MOV L,L", types.Word(152), system.LRegister.Value)
}

func TestMOV_L_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MOV L,M
	HLT
	`
	value := types.Word(0x62)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "MOV L,M", value, system.LRegister.Value)
}

func TestMOV_M_L(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MOV M,L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "MOV M,L", types.Word(0x34), system.RandomAccessMemory.Values[0x1234])
}

