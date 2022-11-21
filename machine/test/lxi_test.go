package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestLXI_B(t *testing.T) {
	source := `
	LXI B,0x1234
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "LXI B (Register B)", types.Word(0x12), system.BRegister.Value)
	check(t, "LXI B (Register C)", types.Word(0x34), system.CRegister.Value)
}

func TestLXI_C(t *testing.T) {
	source := `
	LXI D,0x5678
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "LXI D (Register D)", types.Word(0x56), system.DRegister.Value)
	check(t, "LXI E (Register E)", types.Word(0x78), system.ERegister.Value)
}

func TestLXI_H(t *testing.T) {
	source := `
	LXI H,0xFFEE
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "LXI D (Register H)", types.Word(0xFF), system.HRegister.Value)
	check(t, "LXI E (Register L)", types.Word(0xEE), system.LRegister.Value)
}

func TestLXI_SP(t *testing.T) {
	source := `
	LXI SP,0x5678
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "LXI SP", types.DoubleWord(0x5678), system.StackPointer.Address)
}