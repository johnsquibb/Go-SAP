package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestCMAWithHex(t *testing.T) {
	source := `
	MVI A,0x8
	CMA
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMA with hex", types.Word(0b1111_0111), system.Accumulator.Value)
}

func TestCMAWithBinary(t *testing.T) {
	source := `
	MVI A,0b0000_1000
	CMA
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "CMA with binary", types.Word(0b1111_0111), system.Accumulator.Value)
}