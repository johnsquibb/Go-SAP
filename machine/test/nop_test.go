package test

import (
	"SystemOnePoc/sap3/machine/types"
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