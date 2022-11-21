package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestINR_A(t *testing.T) {
	source := `
	MVI A,0xd
	INR A
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INR A", types.Word(0xe), system.Accumulator.Value)
	check(t, "INR A", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINR_B(t *testing.T) {
	source := `
	MVI B,0x4
	INR B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INR B", types.Word(0x5), system.BRegister.Value)
	check(t, "INR B", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINR_C(t *testing.T) {
	source := `
	MVI C,0x13
	INR C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INR C", types.Word(0x14), system.CRegister.Value)
	check(t, "INR C", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINR_D(t *testing.T) {
	source := `
	MVI D,0xe
	INR D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INR D", types.Word(0xf), system.DRegister.Value)
	check(t, "INR D", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINR_E(t *testing.T) {
	source := `
	MVI E,0x1c
	INR E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INR E", types.Word(0x1d), system.ERegister.Value)
	check(t, "INR E", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINR_H(t *testing.T) {
	source := `
	MVI H,0x4
	INR H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INR H", types.Word(0x5), system.HRegister.Value)
	check(t, "INR H", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINR_L(t *testing.T) {
	source := `
	MVI L,0x12
	INR L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INR L", types.Word(0x13), system.LRegister.Value)
	check(t, "INR L", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINR_M(t *testing.T) {
	source := `
	MVI A,0xFF
	MVI H,0x12
	MVI L,0x34
	INR M
	HLT
	`

	value := types.Word(0x08)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "INR M", types.Word(0xFF), system.Accumulator.Value)
	check(t, "INR M", value+1, system.RandomAccessMemory.Values[0x1234])
	check(t, "INR M (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}
