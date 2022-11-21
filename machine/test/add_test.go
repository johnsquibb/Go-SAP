package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestADD_A(t *testing.T) {
	source := `
	MVI A,0x0
	MVI A,0x26
	ADD A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD A", types.Word(76), system.Accumulator.Value)
	check(t, "ADD A (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_B(t *testing.T) {
	source := `
	MVI A,0x0
	MVI B,0x24
	ADD B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD B", types.Word(36), system.Accumulator.Value)
	check(t, "ADD B (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_C(t *testing.T) {
	source := `
	MVI A,0x0
	MVI C,0x6f
	ADD C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD C", types.Word(111), system.Accumulator.Value)
	check(t, "ADD C (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_D(t *testing.T) {
	source := `
	MVI A,0x0
	MVI D,0x7a
	ADD D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD D", types.Word(122), system.Accumulator.Value)
	check(t, "ADD D (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_E(t *testing.T) {
	source := `
	MVI A,0x0
	MVI E,0xdb
	ADD E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD E", types.Word(219), system.Accumulator.Value)
	check(t, "ADD E (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_H(t *testing.T) {
	source := `
	MVI A,0x0
	MVI H,0xf0
	ADD H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD H", types.Word(240), system.Accumulator.Value)
	check(t, "ADD H (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_L(t *testing.T) {
	source := `
	MVI A,0x0
	MVI L,0x70
	ADD L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD L", types.Word(112), system.Accumulator.Value)
	check(t, "ADD L (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0x0
	ADD M
	HLT
	`

	value := types.Word(127)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "ADD M", value, system.Accumulator.Value)
	check(t, "ADD M (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

// Rollover

func TestADD_A_Rollover(t *testing.T) {
	source := `
    MVI A,0xFF
    ADD A
    HLT
    `

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD A Rollover", types.Word(254), system.Accumulator.Value)
	check(t, "ADD A Rollover (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_B_Rollover(t *testing.T) {
	source := `
	MVI A,0xFF
	MVI B,0xa
	ADD B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD B Rollover", types.Word(9), system.Accumulator.Value)
	check(t, "ADD B Rollover (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_C_Rollover(t *testing.T) {
	source := `
	MVI A,0xFF
	MVI C,0x1e
	ADD C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD C Rollover", types.Word(29), system.Accumulator.Value)
	check(t, "ADD C Rollover (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_D_Rollover(t *testing.T) {
	source := `
	MVI A,0xFF
	MVI D,0x3
	ADD D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD D Rollover", types.Word(2), system.Accumulator.Value)
	check(t, "ADD D Rollover (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_E_Rollover(t *testing.T) {
	source := `
	MVI A,0xFF
	MVI E,0x17
	ADD E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD E Rollover", types.Word(22), system.Accumulator.Value)
	check(t, "ADD E Rollover (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_H_Rollover(t *testing.T) {
	source := `
	MVI A,0xFF
	MVI H,0x7
	ADD H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD H Rollover", types.Word(6), system.Accumulator.Value)
	check(t, "ADD H Rollover (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_L_Rollover(t *testing.T) {
	source := `
	MVI A,0xFF
	MVI L,0x17
	ADD L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADD L Rollover", types.Word(22), system.Accumulator.Value)
	check(t, "ADD L Rollover (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADD_M_Rollover(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0x1
	ADD M
	HLT
	`

	value := types.Word(0xFF)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "ADD M", types.Word(0x0), system.Accumulator.Value)
	check(t, "ADD M (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}