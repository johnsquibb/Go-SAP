package test

import (
	"Go-SAP3/machine/types"
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

func TestSHLD(t *testing.T) {
	source := `
	LXI H,0x1234
	SHLD 0x2000
	HLT
	`

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x2000] = types.Word(0)
	system.RandomAccessMemory.Values[0x2001] = types.Word(0)
	system = runSystem(system)

	check(t, "SHLD (LSB)", types.Word(0x34), system.RandomAccessMemory.Values[0x2000])
	check(t, "SHLD (MSB)", types.Word(0x12), system.RandomAccessMemory.Values[0x2001])
}

func TestSTAX_B(t *testing.T) {
	source := `
	LDA 0x2000
	LXI B,0x1234
	STAX B
	HLT
	`

	value := types.Word(0xFF)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x2000] = value
	system = runSystem(system)

	check(t, "STAX B", value, system.Accumulator.Value)
	check(t, "STAX B", value, system.RandomAccessMemory.Values[0x1234])
}

func TestSTAX_D(t *testing.T) {
	source := `
	LDA 0x2000
	LXI D,0x1234
	STAX D
	HLT
	`

	value := types.Word(0xFE)

	system := getSystem(source)
	system.RandomAccessMemory.Values[0x2000] = value
	system = runSystem(system)

	check(t, "STAX D", value, system.Accumulator.Value)
	check(t, "STAX D", value, system.RandomAccessMemory.Values[0x1234])
}
