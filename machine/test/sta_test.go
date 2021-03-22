package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestSTA(t *testing.T) {
	source := `
	LDA 0x2001
	STA 0x2002
	HLT
	`

	value := types.Word(255)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x2001] = value
	system = runSystem(system)

	check(t, "STA", value, system.Accumulator.Value)
	check(t, "STA", value, system.RandomAccessMemory.Values[0x2002])
}