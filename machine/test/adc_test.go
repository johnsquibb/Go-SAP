package test

import (
	"Go-SAP/machine/types"
	"testing"
)

// TODO Skipped Register A

func TestADC_B(t *testing.T) {
	source := `
	MVI B,0x4
	STC
	ADC B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADC B", types.Word(5), system.Accumulator.Value)
	check(t, "ADC B (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADC_C(t *testing.T) {
	source := `
	MVI C,0xe
	STC
	ADC C
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADC C", types.Word(15), system.Accumulator.Value)
	check(t, "ADC C (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADC_D(t *testing.T) {
	source := `
	MVI D,0x1e
	STC
	ADC D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADC D", types.Word(31), system.Accumulator.Value)
	check(t, "ADC D (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADC_E(t *testing.T) {
	source := `
	MVI E,0x10
	STC
	ADC E
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADC E", types.Word(17), system.Accumulator.Value)
	check(t, "ADC E (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADC_H(t *testing.T) {
	source := `
	MVI H,0x10
	STC
	ADC H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADC H", types.Word(17), system.Accumulator.Value)
	check(t, "ADC H (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADC_L(t *testing.T) {
	source := `
	MVI L,0x11
	STC
	ADC L
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADC L", types.Word(18), system.Accumulator.Value)
	check(t, "ADC L (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestADC_M(t *testing.T) {
	source := `
	MVI H,0x12
	MVI L,0x34
	MVI A,0x0
	STC
	ADC M
	HLT
	`

	value := types.Word(0x32)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "ADC M", value+1, system.Accumulator.Value)
	check(t, "ADC M (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}