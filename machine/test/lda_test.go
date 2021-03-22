package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestLDA(t *testing.T) {
	source := `
	LDA 0x2001
	HLT
	`

	value := types.Word(255)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x2001] = value
	system = runSystem(system)

	check(t, "LDA", value, system.Accumulator.Value)
}

func TestLDAMultipleSTAMultiple(t *testing.T) {
	source := `
	LDA 0x2001
	STA	0x2002

	LDA 0x2003
	STA 0x2004

	LDA 0x2005
	STA 0x2006

	HLT
	`

	system := getSystem(source)

	v1 := types.Word(10)
	v2 := types.Word(20)
	v3 := types.Word(30)

	system.RandomAccessMemory.Values[0x2001] = v1
	system.RandomAccessMemory.Values[0x2003] = v2
	system.RandomAccessMemory.Values[0x2005] = v3

	system = runSystem(system)

	check(t, "LDA, STA", v3, system.Accumulator.Value)
	check(t, "LDA, STA", v1, system.RandomAccessMemory.Values[0x2002])
	check(t, "LDA, STA", v2, system.RandomAccessMemory.Values[0x2004])
	check(t, "LDA, STA", v3, system.RandomAccessMemory.Values[0x2006])
}
