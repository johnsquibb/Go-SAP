package test

import (
	"Go-SAP3/machine/types"
	"testing"
)

func TestRALWhenCarrySet(t *testing.T) {
	source := `
	MVI A,0b0111_0100
	STC
	RAL
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RAL when carry set", types.Word(0b1110_1001), system.Accumulator.Value)
	check(t, "RAL when carry set", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRALWhenMSBHighAndCarrySet(t *testing.T) {
	source := `
	MVI A,0b1111_0100
	STC
	RAL
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RAL when MSB high, carry set", types.Word(0b1110_1001), system.Accumulator.Value)
	check(t, "RAL when MSB high, carry set", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRALWhenCarryNotSet(t *testing.T) {
	source := `
	MVI A,0b0101_0101
	STC
	CMC
	RAL
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RAL when carry not set", types.Word(0b1010_1010), system.Accumulator.Value)
	check(t, "RAL when carry not set", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRARWhenCarrySet(t *testing.T) {
	source := `
	MVI A,0b0111_0100
	STC
	RAR
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RAR when carry set", types.Word(0b1011_1010), system.Accumulator.Value)
	check(t, "RAR when carry set", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRARWhenMSBHighAndCarrySet(t *testing.T) {
	source := `
	MVI A,0b1111_0100
	STC
	RAR
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RAR when MSB high, carry set", types.Word(0b1111_1010), system.Accumulator.Value)
	check(t, "RAR when MSB high, carry set", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRARWhenCarryNotSet(t *testing.T) {
	source := `
	MVI A,0b0101_0101
	STC
	CMC
	RAR
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RAR when carry not set", types.Word(0b0010_1010), system.Accumulator.Value)
	check(t, "RAR when carry not set", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRLCWhenCarrySet(t *testing.T) {
	source := `
	MVI A,0b0111_0100
	STC
	RLC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RLC when carry set", types.Word(0b1110_1000), system.Accumulator.Value)
	check(t, "RLC when carry set", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRLCWhenMSBHighAndCarrySet(t *testing.T) {
	source := `
	MVI A,0b1111_0100
	STC
	RLC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RLC when MSB high, carry set", types.Word(0b1110_1000), system.Accumulator.Value)
	check(t, "RLC when MSB high, carry set", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRLCWhenCarryNotSet(t *testing.T) {
	source := `
	MVI A,0b0101_0101
	STC
	CMC
	RLC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RLC when carry not set", types.Word(0b1010_1010), system.Accumulator.Value)
	check(t, "RLC when carry not set", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRRCWhenCarrySet(t *testing.T) {
	source := `
	MVI A,0b0111_0100
	STC
	RRC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RRC when carry set", types.Word(0b0011_1010), system.Accumulator.Value)
	check(t, "RRC when carry set", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRRCWhenMSBHighAndCarrySet(t *testing.T) {
	source := `
	MVI A,0b1111_0100
	STC
	RRC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RRC when MSB high, carry set", types.Word(0b0111_1010), system.Accumulator.Value)
	check(t, "RRC when MSB high, carry set", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestRRCWhenLSBHighAndCarrySet(t *testing.T) {
	source := `
	MVI A,0b1111_0101
	STC
	RRC
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RRC when MSB high, carry set", types.Word(0b0111_1010), system.Accumulator.Value)
	check(t, "RRC when MSB high, carry set", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}