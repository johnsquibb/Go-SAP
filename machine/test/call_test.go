package test

import (
	"Go-SAP3/machine/types"
	"testing"
)

func TestCALLDirect(t *testing.T) {
	source := `
	LXI SP,0x2100
	CALL 0x07
	HLT
	MVI A,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CALL", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CALL", types.Word(0x20), system.Accumulator.Value)
}

func TestCALLWithLabels(t *testing.T) {
	source := `
	LXI SP,0x2100
	CALL INIT
	CALL INC
	CALL DEC
	CALL INC
	
	HLT
	
	INIT:
		MVI A,0x20
	RET
	
	INC:
		INR A
	RET
	
	DEC:
		DCR A
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CALL", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CALL with labels", types.Word(33), system.Accumulator.Value)
}

func TestCALLWithLabelsAndComments(t *testing.T) {
	source := `
	LXI SP,0x8000
	CALL INIT	; Initialize ACC
	CALL INC	; Increment ACC
	CALL DEC	; Decrement ACC
	CALL INC	; Increment ACC
	
	; END PROGRAM
	HLT
	
	; Initialize ACC
	INIT:
		MVI A,0x20
	RET
	
	; Increment ACC
	INC:
		INR A
	RET
	
	; Decrement ACC	
	DEC:
		DCR A
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CALL", types.DoubleWord(0x8000), system.StackPointer.Address)
	check(t, "CALL with labels and comments", types.Word(33), system.Accumulator.Value)
}

func TestNestedCalls(t *testing.T) {
	source := `
	LXI SP,0x2100
	CALL INIT
	CALL INCDEC
	
	HLT
	
	INIT:
		MVI A,0x20
	RET
	
	INC:
		INR A
	RET
	
	DEC:
		DCR A
	RET

	INCDEC:
		CALL INC
		CALL DEC
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CALL", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CALL with labels", types.Word(32), system.Accumulator.Value)
}

func TestNestedCallsMoreNesting(t *testing.T) {
	source := `
	LXI SP,0x2100
	CALL INIT
	CALL INCDEC
	
	HLT
	
	INIT:
		MVI A,0x20
	RET
	
	INC:
		CALL DOINC
	RET

	DOINC:
		INR A
	RET
	
	DEC:
		DCR A
	RET

	INCDEC:
		CALL INC
		CALL DEC
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CALL", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CALL with labels", types.Word(32), system.Accumulator.Value)
}

// Zero Flag
func TestCNZFlagNotSet(t *testing.T) {
	source := `
	LXI SP,0x2100
	CNZ 0x07
	HLT
	MVI A,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CNZ", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CNZ", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "CNZ", types.Word(0x20), system.Accumulator.Value)
}

func TestCNZFlagSet(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	ADD A
	CNZ 0x0A
	HLT
	MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CNZ", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CNZ", types.High, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "CNZ", types.Word(0x00), system.BRegister.Value)
}

func TestCZFlagNotSet(t *testing.T) {
	source := `
	LXI SP,0x2100
	CZ 0x07
	HLT
	MVI A,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CZ", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CZ", types.Low, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "CZ", types.Word(0x0), system.Accumulator.Value)
}

func TestCZFlagSet(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	ADD A
	CZ 0x0A
	HLT
	MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CZ", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CZ", types.High, system.ArithmeticLogicUnit.Flags.Zero)
	check(t, "CZ", types.Word(0x20), system.BRegister.Value)
}

// Carry Flag
func TestCNCFlagNotSet(t *testing.T) {
	source := `
	LXI SP,0x2100
	CNC 0x07
	HLT
	MVI A,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CNC", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CNC", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "CNC", types.Word(0x20), system.Accumulator.Value)
}

func TestCNCFlagSet(t *testing.T) {
	source := `
	LXI SP,0x2100
	STC
	CNC 0x08
	HLT
	MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CNC", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CNC t", types.High, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "CNC", types.Word(0x00), system.BRegister.Value)
}

func TestCCFlagNotSet(t *testing.T) {
	source := `
	LXI SP,0x2100
	STC
	CMC
	CC 0x09
	HLT
	MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CC", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CC", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "CC", types.Word(0x0), system.Accumulator.Value)
}

func TestCCFlagSet(t *testing.T) {
	source := `
	LXI SP,0x2100
	STC
	CC 0x08
	HLT
	MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CC", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CC", types.High, system.ArithmeticLogicUnit.Flags.Carry)
	check(t, "CC", types.Word(0x20), system.BRegister.Value)
}

// Parity flag
func TestCPOIsOdd(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	CPO DO
	HLT
	DO: MVI A,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CPO", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CPO", types.Low, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "CPO", types.Word(0x20), system.Accumulator.Value)
}

// Parity flag
func TestCPOIsEven(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	INR A
	INR A
	CPO DO
	HLT
	DO: MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CPO", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CPO", types.High, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "CPO", types.Word(0x00), system.BRegister.Value)
}

func TestCPEIsOdd(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	CPE DO
	HLT
	DO: MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CPE", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CPE", types.Low, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "CPE", types.Word(0x0), system.BRegister.Value)
}

// Parity flag
func TestCPEIsEven(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	INR A
	INR A
	CPE DO
	HLT
	DO: MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CPE", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CPE", types.High, system.ArithmeticLogicUnit.Flags.Parity)
	check(t, "CPE", types.Word(0x20), system.BRegister.Value)
}

// Sign flag
func TestCPIsPositive(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	CP DO
	HLT
	DO: MVI A,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CP", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CP", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
	check(t, "CP", types.Word(0x20), system.Accumulator.Value)
}

func TestCPIsNegative(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	DCR A
	CP DO
	HLT
	DO: MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CP", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CP", types.High, system.ArithmeticLogicUnit.Flags.Sign)
	check(t, "CP", types.Word(0x00), system.BRegister.Value)
}

func TestCMIsPositive(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	INR A
	CM DO
	HLT
	DO: MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CM", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CM", types.Low, system.ArithmeticLogicUnit.Flags.Sign)
	check(t, "CM", types.Word(0x00), system.BRegister.Value)
}

func TestCMIsNegative(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0x0
	DCR A
	CM DO
	HLT
	DO: MVI B,0x20
	RET
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CM", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "CM", types.High, system.ArithmeticLogicUnit.Flags.Sign)
	check(t, "CM", types.Word(0x20), system.BRegister.Value)
}
