package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestPOP_B(t *testing.T) {
	source := `
	LXI SP,0x2100
	LXI D,0x1234
	PUSH D
	POP B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "POP B", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "POP B (MSB)", system.BRegister.Value, system.DRegister.Value)
	check(t, "POP B (LSB)", system.CRegister.Value, system.ERegister.Value)
}

func TestPOP_D(t *testing.T) {
	source := `
	LXI SP,0x2100
	LXI B,0x1234
	PUSH B
	POP D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "POP D", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "POP D (MSB)", system.DRegister.Value, system.BRegister.Value)
	check(t, "POP D (LSB)", system.ERegister.Value, system.CRegister.Value)
}

func TestPOP_H(t *testing.T) {
	source := `
	LXI SP,0x2100
	LXI B,0x1234
	PUSH B
	POP H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "POP H", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "POP H (MSB)", system.HRegister.Value, system.BRegister.Value)
	check(t, "POP H (LSB)", system.LRegister.Value, system.CRegister.Value)
}

func TestPOP_PSW(t *testing.T) {
	source := `
	LXI SP,0x2100
	LXI B,0x1234
	PUSH B
	POP PSW
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "POP PSW", types.DoubleWord(0x2100), system.StackPointer.Address)
	check(t, "POP PSW (MSB)", system.Accumulator.Value, system.BRegister.Value)
	check(t, "POP PSW (LSB)", system.FRegister.Value, system.CRegister.Value)
}