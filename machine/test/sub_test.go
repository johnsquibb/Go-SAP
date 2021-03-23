package test

import (
	"Go-SAP3/machine/types"
	"testing"
)

func TestSUB_A(t *testing.T) {
	source := `
	MVI A,0x6
	MVI A,0x6
	SUB A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SUB A", types.Word(0), system.Accumulator.Value)
	check(t, "SUB A (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestSUB_B(t *testing.T) {
	source := `
	MVI A,0x23
	MVI B,0x13
	SUB B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SUB B", types.Word(16), system.Accumulator.Value)
	check(t, "SUB B (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestSUB_C(t *testing.T) {
	source := `
	MVI A,0x78
	MVI C,0x8
	SUB C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SUB C", types.Word(112), system.Accumulator.Value)
	check(t, "SUB C (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestSUB_D(t *testing.T) {
	source := `
	MVI A,0x5e
	MVI D,0x16
	SUB D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SUB D", types.Word(72), system.Accumulator.Value)
	check(t, "SUB D (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestSUB_E(t *testing.T) {
	source := `
	MVI A,0x70
	MVI E,0x1
	SUB E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SUB E", types.Word(111), system.Accumulator.Value)
	check(t, "SUB E (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestSUB_H(t *testing.T) {
	source := `
	MVI A,0x6b
	MVI H,0x1
	SUB H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SUB H", types.Word(106), system.Accumulator.Value)
	check(t, "SUB H (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestSUB_L(t *testing.T) {
	source := `
	MVI A,0x5c
	MVI L,0x14
	SUB L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SUB L", types.Word(72), system.Accumulator.Value)
	check(t, "SUB L (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestSUB_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0x32
	SUB M
	HLT
	`

	value := types.Word(0x8)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "SUB M", types.Word(0x2A), system.Accumulator.Value)
	check(t, "SUB M (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

