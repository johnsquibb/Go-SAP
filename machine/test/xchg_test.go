package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestXCHG(t *testing.T) {
	source := `
	LXI H,0x1234
	LXI D,0x5678
	XCHG
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "XCHG", types.Word(0x12), system.DRegister.Value)
	check(t, "XCHG", types.Word(0x34), system.ERegister.Value)

	check(t, "XCHG", types.Word(0x56), system.HRegister.Value)
	check(t, "XCHG", types.Word(0x78), system.LRegister.Value)
}
