package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestSBB_A(t *testing.T) {
	source := `
	MVI A,0x15
	MVI A,0x15
	STC 
	SBB A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SBB A", types.Word(255), system.Accumulator.Value)
	check(t, "SBB A (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "SBB A (zero)", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "SBB A (sign)", types.High, system.ArithmeticLogicUnit.Flags.Sign)
}

func TestSBB_B(t *testing.T) {
	source := `
	MVI A,0x7d
	MVI B,0x15
	STC 
	SBB B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SBB B", types.Word(103), system.Accumulator.Value)
	check(t, "SBB B (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "SBB B (zero)", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "SBB B (sign)", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
}

func TestSBB_C(t *testing.T) {
	source := `
	MVI A,0x4f
	MVI C,0x17
	STC 
	SBB C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SBB C", types.Word(55), system.Accumulator.Value)
	check(t, "SBB C (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "SBB C (zero)", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "SBB C (sign)", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
}

func TestSBB_D(t *testing.T) {
	source := `
	MVI A,0x6d
	MVI D,0x17
	STC 
	SBB D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SBB D", types.Word(85), system.Accumulator.Value)
	check(t, "SBB D (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "SBB D (zero)", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "SBB D (sign)", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
}

func TestSBB_E(t *testing.T) {
	source := `
	MVI A,0x6a
	MVI E,0x20
	STC 
	SBB E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SBB E", types.Word(73), system.Accumulator.Value)
	check(t, "SBB E (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "SBB E (zero)", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "SBB E (sign)", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
}

func TestSBB_H(t *testing.T) {
	source := `
	MVI A,0x49
	MVI H,0x18
	STC 
	SBB H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SBB H", types.Word(48), system.Accumulator.Value)
	check(t, "SBB H (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "SBB H (zero)", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "SBB H (sign)", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
}

func TestSBB_L(t *testing.T) {
	source := `
	MVI A,0x43
	MVI L,0x1c
	STC 
	SBB L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SBB L", types.Word(38), system.Accumulator.Value)
	check(t, "SBB L (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "SBB L (zero)", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "SBB L (sign)", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
}

func TestSBB_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0x32
	STC
	SBB M
	HLT
	`

	value := types.Word(0x08)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "SBB M", types.Word(0x2a) - 1, system.Accumulator.Value)
	check(t, "SBB L (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "SBB L (zero)", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "SBB L (sign)", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
}