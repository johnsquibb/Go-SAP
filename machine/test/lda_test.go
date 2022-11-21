package test

import (
	"Go-SAP/machine/types"
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

func TestLHLD(t *testing.T) {
	source := `
	LHLD 0x1234
	HLT
	`

	v1 := types.Word(0xFF)
	v2 := types.Word(0xEE)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = v1
	system.RandomAccessMemory.Values[0x1235] = v2
	system = runSystem(system)

	check(t, "LHLD (REG L)", v1, system.LRegister.Value)
	check(t, "LHLD (REG H)", v2, system.HRegister.Value)
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

func TestLDAX_B(t *testing.T) {
	source := `
	LXI B,0x1234
	LDAX B
	HLT
	`

	value := types.Word(255)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "LDAX B", value, system.Accumulator.Value)
}

func TestLDAX_D(t *testing.T) {
	source := `
	LXI D,0x1234
	LDAX D
	HLT
	`

	value := types.Word(255)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x1234] = value
	system = runSystem(system)

	check(t, "LDAX D", value, system.Accumulator.Value)
}
