package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestDCR_A(t *testing.T) {
	source := `
	MVI A,0x10
	DCR A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCR A", types.Word(0xf), system.Accumulator.Value)
	check(t, "DCR A", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCR_B(t *testing.T) {
	source := `
	MVI B,0x16
	DCR B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCR B", types.Word(0x15), system.BRegister.Value)
	check(t, "DCR B", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCR_C(t *testing.T) {
	source := `
	MVI C,0x17
	DCR C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCR C", types.Word(0x16), system.CRegister.Value)
	check(t, "DCR C", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCR_D(t *testing.T) {
	source := `
	MVI D,0x3
	DCR D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCR D", types.Word(0x2), system.DRegister.Value)
	check(t, "DCR D", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCR_E(t *testing.T) {
	source := `
	MVI E,0x17
	DCR E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCR E", types.Word(0x16), system.ERegister.Value)
	check(t, "DCR E", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCR_H(t *testing.T) {
	source := `
	MVI H,0x1f
	DCR H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCR H", types.Word(0x1e), system.HRegister.Value)
	check(t, "DCR H", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCR_L(t *testing.T) {
	source := `
	MVI L,0x17
	DCR L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCR L", types.Word(0x16), system.LRegister.Value)
	check(t, "DCR L", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCR_M(t *testing.T) {
	source := `
	MVI A,0xFF
	MVI H,0x12
	MVI L,0x34
	DCR M
	HLT
	`

	value := types.Word(0x08)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "DCR M", types.Word(0xFF), system.Accumulator.Value)
	check(t, "DCR M", value-1, system.RandomAccessMemory.Values[0x1234])
	check(t, "DCR M (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}