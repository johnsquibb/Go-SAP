package test

import (
	"SystemOnePoc/sap3/machine/types"
	"testing"
)

func TestMVI_A(t *testing.T) {
	source := `
	MVI A,0x5
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MVI A", types.Word(5), system.Accumulator.Value)
}

func TestMVI_B(t *testing.T) {
	source := `
	MVI B,0x74
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MVI B", types.Word(116), system.BRegister.Value)
}

func TestMVI_C(t *testing.T) {
	source := `
	MVI C,0x67
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MVI C", types.Word(103), system.CRegister.Value)
}

func TestMVI_D(t *testing.T) {
	source := `
	MVI D,0x34
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MVI D", types.Word(52), system.DRegister.Value)
}

func TestMVI_E(t *testing.T) {
	source := `
	MVI E,0x45
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MVI E", types.Word(69), system.ERegister.Value)
}

func TestMVI_H(t *testing.T) {
	source := `
	MVI H,0x2e
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MVI H", types.Word(46), system.HRegister.Value)
}

func TestMVI_L(t *testing.T) {
	source := `
	MVI L,0xfb
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MVI L", types.Word(251), system.LRegister.Value)
}

func TestMVI_M(t *testing.T) {
	source := `
	MVI L,0x88
	MVI H,0x88
	MVI M,0x32
	HLT
	`
	system := getSystem(source)
	system = runSystem(system)
	check(t, "MVI M", types.Word(0x32), system.RandomAccessMemory.Values[0x8888])
}