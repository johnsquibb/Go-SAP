package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestANI(t *testing.T) {
	source := `
	MVI A,0xC
	ANI 0x8
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ANI", types.Word(0x8), system.Accumulator.Value)
}

func TestORI(t *testing.T) {
	source := `
	MVI A,0xC
	ORI 0x8
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ORI", types.Word(0xC), system.Accumulator.Value)
}

func TestXRI(t *testing.T) {
	source := `
	MVI A,0x8
	XRI 0xC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XRI", types.Word(0x4), system.Accumulator.Value)
}

func TestADI(t *testing.T) {
	source := `
	MVI A,0x04
	ADI 0x08
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ADI", types.Word(0xC), system.Accumulator.Value)
}

func TestACI(t *testing.T) {
	source := `
	MVI A,0x04
	STC
	ACI 0x08
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "ACI", types.Word(0xD), system.Accumulator.Value)
}

func TestSUI(t *testing.T) {
	source := `
	MVI A,0x08
	SUI 0x04
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SUI", types.Word(0x4), system.Accumulator.Value)
}

func TestSBI(t *testing.T) {
	source := `
	MVI A,0x08
	STC
	SBI 0x04
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "SBI", types.Word(0x3), system.Accumulator.Value)
}

func TestCPI(t *testing.T) {
	source := `
	MVI A,0x04
	CPI 0x04
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CPI", types.Word(0x4), system.Accumulator.Value)
	check(t, "CPI", types.High, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCPINotZero(t *testing.T) {
	source := `
	MVI A,0x05
	CPI 0x04
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CPI", types.Word(0x5), system.Accumulator.Value)
	check(t, "CPI", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
}

func TestCPISign(t *testing.T) {
	source := `
	MVI A,0x03
	CPI 0x04
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CPI", types.Word(0x3), system.Accumulator.Value)
	check(t, "CPI", types.High, system.ArithmeticLogicUnit.Flags.Sign)
}

func TestCPICarry(t *testing.T) {
	source := `
	MVI A,0x0
	CPI 0x04
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CPI", types.Word(0x0), system.Accumulator.Value)
	check(t, "CPI", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}
