package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestDCX_B(t *testing.T) {
	source := `
	LXI B,0x1234
	DCX B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCX B (B Reg)", types.Word(0x12), system.BRegister.Value)
	check(t, "DCX B (C Reg)", types.Word(0x33), system.CRegister.Value)
	check(t, "DCX B (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCX_D(t *testing.T) {
	source := `
	LXI D,0x1234
	DCX D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCX D (D Reg)", types.Word(0x12), system.DRegister.Value)
	check(t, "DCX D (E Reg)", types.Word(0x33), system.ERegister.Value)
	check(t, "DCX D (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCX_H(t *testing.T) {
	source := `
	LXI H,0x1234
	DCX H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCX H (H Reg)", types.Word(0x12), system.HRegister.Value)
	check(t, "DCX H (L Reg)", types.Word(0x33), system.LRegister.Value)
	check(t, "DCX H (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCX_BMSB(t *testing.T) {
	source := `
	LXI B,0x1200
	DCX B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCX B MSB (B Reg)", types.Word(0x11), system.BRegister.Value)
	check(t, "DCX B MSB(C Reg)", types.Word(0xFF), system.CRegister.Value)
	check(t, "DCX B MSB (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCX_DMSB(t *testing.T) {
	source := `
	LXI D,0x1200
	DCX D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCX D MSB(D Reg)", types.Word(0x11), system.DRegister.Value)
	check(t, "DCX D MSB (E Reg)", types.Word(0xFF), system.ERegister.Value)
	check(t, "DCX D MSB (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCX_HMSB(t *testing.T) {
	source := `
	LXI H,0x1200
	DCX H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCX H MSB (H Reg)", types.Word(0x11), system.HRegister.Value)
	check(t, "DCX H MSB (L Reg)", types.Word(0xFF), system.LRegister.Value)
	check(t, "DCX H MSB (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCX_BRollover(t *testing.T) {
	source := `
	LXI B,0x0000
	DCX B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCX B Rollover (B Reg)", types.Word(0xFF), system.BRegister.Value)
	check(t, "DCX B Rollover(C Reg)", types.Word(0xFF), system.CRegister.Value)
	check(t, "DCX B Rollover (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCX_DRollover(t *testing.T) {
	source := `
	LXI D,0x0000
	DCX D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCX D Rollover(D Reg)", types.Word(0xFF), system.DRegister.Value)
	check(t, "DCX D Rollover (E Reg)", types.Word(0xFF), system.ERegister.Value)
	check(t, "DCX D Rollover (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDCX_HRollover(t *testing.T) {
	source := `
	LXI H,0x0000
	DCX H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DCX H Rollover (H Reg)", types.Word(0xFF), system.HRegister.Value)
	check(t, "DCX H Rollover (L Reg)", types.Word(0xFF), system.LRegister.Value)
	check(t, "DCX H Rollover (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}