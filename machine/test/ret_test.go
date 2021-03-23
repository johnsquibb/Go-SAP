package test

import (
	"Go-SAP3/machine/types"
	"testing"
)

func TestRET(t *testing.T) {
	source := `
	LXI SP,0x0080
	CALL 0x7
	HLT 
	INR A
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RET", types.DoubleWord(0x0080), system.StackPointer.Address)
	check(t, "RET", types.Word(1), system.Accumulator.Value)
}

// Zero flag.
func TestRNZNotZero(t *testing.T) {
	source := `
	LXI SP,0x2100
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RNZ 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RNZ", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RNZ", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "RNZ", types.Word(0x20), system.BRegister.Value)
}

func TestRNZIsZero(t *testing.T) {
	source := `
	LXI SP,0x2100
	INR A
	DCR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RNZ 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RNZ", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RNZ", types.High, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "RNZ", types.Word(0x40), system.BRegister.Value)
}

func TestRZNotZero(t *testing.T) {
	source := `
	LXI SP,0x2100
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RZ 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RZ", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RZ", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "RZ", types.Word(0x40), system.BRegister.Value)
}

func TestRZIsZero(t *testing.T) {
	source := `
	LXI SP,0x2100
	INR A
	DCR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RZ 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RZ", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RZ", types.High, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "RZ", types.Word(0x20), system.BRegister.Value)
}

// Carry flag
func TestRNCNotCarry(t *testing.T) {
	source := `
	LXI SP,0x2100
	STC 
	CMC
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RNC 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RNC", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RNC", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "RNC", types.Word(0x20), system.BRegister.Value)
}

func TestRNCCarry(t *testing.T) {
	source := `
	LXI SP,0x2100
	STC 
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RNC 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RNC", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RNC", types.High, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "RNC", types.Word(0x40), system.BRegister.Value)
}

func TestRCNotCarry(t *testing.T) {
	source := `
	LXI SP,0x2100
	STC 
	CMC
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RC 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RC", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RC", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "RC", types.Word(0x40), system.BRegister.Value)
}

func TestRCCarry(t *testing.T) {
	source := `
	LXI SP,0x2100
	STC 
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RC 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RC", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RC", types.High, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "RC", types.Word(0x20), system.BRegister.Value)
}

// Parity flag
func TestRPOIsOdd(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RPO 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RPO", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RPO", types.Low, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "RPO", types.Word(0x20), system.BRegister.Value)
}

func TestRPOIsEven(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	INR A
	INR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RPO 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RPO", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RPO", types.High, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "RPO", types.Word(0x40), system.BRegister.Value)
}

func TestRPEIsOdd(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RPE 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RPE", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RPE", types.Low, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "RPE", types.Word(0x40), system.BRegister.Value)
}

func TestRPEIsEven(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	INR A
	INR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RPE 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RPE", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RPE", types.High, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "RPE", types.Word(0x20), system.BRegister.Value)
}

// Sign
func TestRPIsPositive(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RP 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RP", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RP", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
	check(t, "RP", types.Word(0x20), system.BRegister.Value)
}

func TestRPIsNegative(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	DCR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RP 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RP", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RP", types.High, system.ArithmeticLogicUnit.Flags.Sign)
	check(t, "RP", types.Word(0x40), system.BRegister.Value)
}

func TestRMIsPositive(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RM 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RM", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RM", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
	check(t, "RM", types.Word(0x40), system.BRegister.Value)
}

func TestRMIsNegative(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	DCR A
	CALL DO
	HLT 

	DO:
	MVI B,0x20
	RM 
	MVI B,0x40
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "RM", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "RM", types.High, system.ArithmeticLogicUnit.Flags.Sign)
	check(t, "RM", types.Word(0x20), system.BRegister.Value)
}