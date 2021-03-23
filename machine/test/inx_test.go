package test

import (
	"Go-SAP3/machine/types"
	"testing"
)

func TestINX_B(t *testing.T) {
	source := `
	LXI B,0x1234
	INX B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INX B (B Reg)", types.Word(0x12), system.BRegister.Value)
	check(t, "INX B (C Reg)", types.Word(0x35), system.CRegister.Value)
	check(t, "INX B (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINX_D(t *testing.T) {
	source := `
	LXI D,0x1234
	INX D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INX D (D Reg)", types.Word(0x12), system.DRegister.Value)
	check(t, "INX D (E Reg)", types.Word(0x35), system.ERegister.Value)
	check(t, "INX D (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINX_H(t *testing.T) {
	source := `
	LXI H,0x1234
	INX H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INX H (H Reg)", types.Word(0x12), system.HRegister.Value)
	check(t, "INX H (L Reg)", types.Word(0x35), system.LRegister.Value)
	check(t, "INX H (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINX_BMSB(t *testing.T) {
	source := `
	LXI B,0x12FF
	INX B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INX B MSB (B Reg)", types.Word(0x13), system.BRegister.Value)
	check(t, "INX B MSB(C Reg)", types.Word(0x0), system.CRegister.Value)
	check(t, "INX B MSB (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINX_DMSB(t *testing.T) {
	source := `
	LXI D,0x12FF
	INX D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INX D MSB(D Reg)", types.Word(0x13), system.DRegister.Value)
	check(t, "INX D MSB (E Reg)", types.Word(0x0), system.ERegister.Value)
	check(t, "INX D MSB (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINX_HMSB(t *testing.T) {
	source := `
	LXI H,0x12FF
	INX H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INX H MSB (H Reg)", types.Word(0x13), system.HRegister.Value)
	check(t, "INX H MSB (L Reg)", types.Word(0x0), system.LRegister.Value)
	check(t, "INX H MSB (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINX_BRollover(t *testing.T) {
	source := `
	LXI B,0xFFFF
	INX B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INX B Rollover (B Reg)", types.Word(0x0), system.BRegister.Value)
	check(t, "INX B Rollover(C Reg)", types.Word(0x0), system.CRegister.Value)
	check(t, "INX B Rollover (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINX_DRollover(t *testing.T) {
	source := `
	LXI D,0xFFFF
	INX D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INX D Rollover(D Reg)", types.Word(0x0), system.DRegister.Value)
	check(t, "INX D Rollover (E Reg)", types.Word(0x0), system.ERegister.Value)
	check(t, "INX D Rollover (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestINX_HRollover(t *testing.T) {
	source := `
	LXI H,0xFFFF
	INX H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "INX H Rollover (H Reg)", types.Word(0x0), system.HRegister.Value)
	check(t, "INX H Rollover (L Reg)", types.Word(0x0), system.LRegister.Value)
	check(t, "INX H Rollover (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}