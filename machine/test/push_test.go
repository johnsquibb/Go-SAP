package test

import (
	"Go-SAP3/machine/types"
	"testing"
)

func TestPUSH_B(t *testing.T) {
	source := `
	LXI SP,0x2100
	LXI B,0x1234
	PUSH B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "PUSH B", types.DoubleWord(0x20FE), system.StackPointer.Address)
	check(t, "PUSH B (RAM MSB)", system.BRegister.Value, system.RandomAccessMemory.Values[0x20FF])
	check(t, "PUSH B (RAM LSB)", system.CRegister.Value, system.RandomAccessMemory.Values[0x20FE])
}

func TestPUSH_D(t *testing.T) {
	source := `
	LXI SP,0x2100
	LXI D,0x1234
	PUSH D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "PUSH D", types.DoubleWord(0x20FE), system.StackPointer.Address)
	check(t, "PUSH D (RAM MSB)", system.DRegister.Value, system.RandomAccessMemory.Values[0x20FF])
	check(t, "PUSH D (RAM LSB)", system.ERegister.Value, system.RandomAccessMemory.Values[0x20FE])
}

func TestPUSH_H(t *testing.T) {
	source := `
	LXI SP,0x2100
	LXI H,0x1234
	PUSH H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "PUSH H", types.DoubleWord(0x20FE), system.StackPointer.Address)
	check(t, "PUSH H (RAM MSB)", system.HRegister.Value, system.RandomAccessMemory.Values[0x20FF])
	check(t, "PUSH H (RAM LSB)", system.LRegister.Value, system.RandomAccessMemory.Values[0x20FE])
}

func TestPUSH_PSW(t *testing.T) {
	source := `
	LXI SP,0x2100
	MVI A,0xFF
	STC
	PUSH PSW
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "PUSH PSW SP", types.DoubleWord(0x20FE), system.StackPointer.Address)
	check(t, "PUSH PSW SP (RAM MSB)", system.Accumulator.Value, system.RandomAccessMemory.Values[0x20FF])
	check(t, "PUSH PSW SP (RAM LSB)", system.FRegister.Value, system.RandomAccessMemory.Values[0x20FE])
}