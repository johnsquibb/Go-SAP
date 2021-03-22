package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestJMP(t *testing.T) {
	source := `
	JMP SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JMP", types.Word(0), system.Accumulator.Value)
	check(t, "JMP", types.Word(0x16), system.BRegister.Value)
}

func TestJM(t *testing.T) {
	source := `
	DCR A
	JM SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JM", types.Word(255), system.Accumulator.Value)
	check(t, "JM", types.Word(0x16), system.BRegister.Value)
}

func TestJP(t *testing.T) {
	source := `
	INR A
	JP SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JM", types.Word(1), system.Accumulator.Value)
	check(t, "JM", types.Word(0x16), system.BRegister.Value)
}

func TestJZ(t *testing.T) {
	source := `
	DCR A
	INR A
	JZ SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JZ", types.Word(0), system.Accumulator.Value)
	check(t, "JZ", types.Word(0x16), system.BRegister.Value)
}

func TestJNZ(t *testing.T) {
	source := `
	DCR A
	JNZ SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JNZ", types.Word(255), system.Accumulator.Value)
	check(t, "JNZ", types.Word(0x16), system.BRegister.Value)
}

func TestJC(t *testing.T) {
	source := `
	INR A
	STC
	JC SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JC", types.Word(1), system.Accumulator.Value)
	check(t, "JC", types.Word(0x16), system.BRegister.Value)
}

func TestJNC(t *testing.T) {
	source := `
	INR A
	STC
	CMC
	JNC SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JNC", types.Word(1), system.Accumulator.Value)
	check(t, "JNC", types.Word(0x16), system.BRegister.Value)
}

func TestJPENonZero(t *testing.T) {
	source := `
	MVI A,0x02
	INR A
	JPE SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JPE non-zero", types.High, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "JPE non-zero", types.Word(3), system.Accumulator.Value)
	check(t, "JPE non-zero", types.Word(0x16), system.BRegister.Value)
}

func TestJPEZero(t *testing.T) {
	source := `
	MVI A,0x0
	INR A
	DCR A
	JPE SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JPE zero", types.High, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "JPE zero", types.Word(0), system.Accumulator.Value)
	check(t, "JPE zero", types.Word(0x16), system.BRegister.Value)
}

func TestJPO(t *testing.T) {
	source := `
	MVI A,0x0
	INR A
	JPO SETB
	SETA: MVI A,0x08
	SETB: MVI B,0x16
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "JPO", types.Low, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "JPO", types.Word(1), system.Accumulator.Value)
	check(t, "JPO", types.Word(0x16), system.BRegister.Value)
}
