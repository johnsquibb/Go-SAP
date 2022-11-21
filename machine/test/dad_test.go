package test

import (
	"Go-SAP/machine/types"
	"testing"
)

func TestDAD_B(t *testing.T) {
	source := `
	LXI H,0x0000
	LXI B,0x1234
	DAD B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DAD B (H)", types.Word(18), system.HRegister.Value)
	check(t, "DAD B (L)", types.Word(52), system.LRegister.Value)
	check(t, "DAD B (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDAD_BCarry(t *testing.T) {
	source := `
	LXI H,0xFFFF
	LXI B,0x0001
	DAD B
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DAD B Carry (H)", types.Word(0), system.HRegister.Value)
	check(t, "DAD B Carry (L)", types.Word(0), system.LRegister.Value)
	check(t, "DAD B Carry (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDAD_D(t *testing.T) {
	source := `
	LXI H,0x0000
	LXI D,0x1234
	DAD D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DAD D (H)", types.Word(18), system.HRegister.Value)
	check(t, "DAD D (L)", types.Word(52), system.LRegister.Value)
	check(t, "DAD D (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDAD_DCarry(t *testing.T) {
	source := `
	LXI H,0xFFFF
	LXI D,0x0001
	DAD D
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DAD D Carry (H)", types.Word(0), system.HRegister.Value)
	check(t, "DAD D Carry (L)", types.Word(0), system.LRegister.Value)
	check(t, "DAD D Carry (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDAD_H(t *testing.T) {
	source := `
	LXI H,0x1234
	DAD H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DAD H (H)", types.Word(36), system.HRegister.Value)
	check(t, "DAD H (L)", types.Word(104), system.LRegister.Value)
	check(t, "DAD H (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDAD_SP(t *testing.T) {
	source := `
	LXI H,0x1234
	LXI SP,0x1111
	DAD SP
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DAD H (H)", types.Word(0x23), system.HRegister.Value)
	check(t, "DAD H (L)", types.Word(0x45), system.LRegister.Value)
	check(t, "DAD H (carry)", types.Low, system.ArithmeticLogicUnit.Flags.Carry)
}

func TestDAD_HCarry(t *testing.T) {
	source := `
	LXI H,0x8080
	DAD H
	HLT
	`

	system := getSystem(source)
	system = runSystem(system)

	check(t, "DAD D Carry (H)", types.Word(1), system.HRegister.Value)
	check(t, "DAD D Carry (L)", types.Word(0), system.LRegister.Value)
	check(t, "DAD D Carry (carry)", types.High, system.ArithmeticLogicUnit.Flags.Carry)
}
