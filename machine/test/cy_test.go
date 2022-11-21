package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestCarryDefault(t *testing.T) {
	source := `
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "Carry default", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestSTC(t *testing.T) {
	source := `
	STC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "STC", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestSTCTwice(t *testing.T) {
	source := `
	STC
	STC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "STC (twice)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestCMCDefault(t *testing.T) {
	source := `
	CMC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMC", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestCMCAfterSTC(t *testing.T) {
	source := `
	STC
	CMC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMC", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}