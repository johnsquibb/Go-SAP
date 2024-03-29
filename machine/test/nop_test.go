package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestNOP(t *testing.T) {
	source := `
	NOP
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "NOP", types.Word(0), system.Accumulator.Value)
}
