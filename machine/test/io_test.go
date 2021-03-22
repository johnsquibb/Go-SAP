package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestOutPort3(t *testing.T) {
	source := `
	MVI A,0x01
	OUT 0x3
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "OUT 3", types.Word(0x01), system.OutputRegister3.Value)
}

func TestOutPort4(t *testing.T) {
	source := `
	MVI A,0x01
	OUT 0x4
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "OUT 4", types.Word(0x01), system.OutputRegister4.Value)
}

func TestInPort1(t *testing.T) {
	source := `
	IN 0x1
	HLT
	`

	system := getSystem(source)
	system.InputRegister1.Value = 0x32
	system = runSystem(system)

	check(t, "IN 1", types.Word(0x32), system.Accumulator.Value)
}

func TestInPort2(t *testing.T) {
	source := `
	IN 0x2
	HLT
	`

	system := getSystem(source)
	system.InputRegister2.Value = 0x32
	system = runSystem(system)

	check(t, "IN 2", types.Word(0x32), system.Accumulator.Value)
}
