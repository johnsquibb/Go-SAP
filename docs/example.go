package main

import (
	"Go-SAP3/machine"
	"Go-SAP3/machine/asm"
	"Go-SAP3/machine/types"
	"fmt"
	"log"
)

// Sample Assembly source code.
// Assembly code is converted to machine code by Assemble method.
var source = `
; -------------------------
; CALL example with labels.
; -------------------------

CALL INIT
CALL INC4
CALL DEC4
CALL INC4

OUT 0x3

CALL INC4
OUT 0x4

HLT

INIT:
	MVI A,0x20
RET

INC4:
	INR A
	INR A
	INR A
	INR A
RET

DEC4:
	DCR A
	DCR A
	DCR A
	DCR A
RET
`

func main() {
	// Assemble the sample source into machine instructions.
	instructions, directives, err := asm.Assemble(source)
	if err != nil {
		log.Fatal(err)
	}

	// Apply assembled instructions to RAM.
	// RAM contents will be read during system.Start()
	machine.ApplyInstructionsToRamContents(instructions, directives.Address)

	// Create new system, set the program counter per directive.
	system := machine.New()
	system.Start()
	system.ProgramCounter.Value = types.DoubleWord(directives.PC)

	for {
		system.Update()

		system.StepClock()

		if system.Halt {
			break
		}
	}

	// Display the total number of machine cycles used.
	fmt.Println()
	fmt.Println(fmt.Sprintf("Machine Cycles:\t %d", machine.UpdateCtr))
	fmt.Println()

	// Display the contents of the various registers.
	fmt.Println(fmt.Sprintf("REG A:\t 0x%X\t (%d)", system.Accumulator.Value, system.Accumulator.Value))
	fmt.Println(fmt.Sprintf("REG B:\t 0x%X\t (%d)", system.BRegister.Value, system.BRegister.Value))
	fmt.Println(fmt.Sprintf("REG C:\t 0x%X\t (%d)", system.CRegister.Value, system.CRegister.Value))
	fmt.Println(fmt.Sprintf("REG D:\t 0x%X\t (%d)", system.DRegister.Value, system.DRegister.Value))
	fmt.Println(fmt.Sprintf("REG E:\t 0x%X\t (%d)", system.ERegister.Value, system.ERegister.Value))
	fmt.Println(fmt.Sprintf("REG F:\t 0x%X\t (%d)", system.FRegister.Value, system.FRegister.Value))
	fmt.Println(fmt.Sprintf("REG H:\t 0x%X\t (%d)", system.HRegister.Value, system.HRegister.Value))
	fmt.Println(fmt.Sprintf("REG L:\t 0x%X\t (%d)", system.LRegister.Value, system.LRegister.Value))
	fmt.Println(fmt.Sprintf("OUTPUT (3):\t 0x%X\t (%d)", system.OutputRegister3.Value, system.OutputRegister3.Value))
	fmt.Println(fmt.Sprintf("OUTPUT (4):\t 0x%X\t (%d)", system.OutputRegister4.Value, system.OutputRegister4.Value))
	fmt.Println()

	// Display the state of ALU flags.
	fmt.Println("Sign:\t", system.ArithmeticLogicUnit.Flags.Sign)
	fmt.Println("Zero:\t", system.ArithmeticLogicUnit.Flags.Zero)
	fmt.Println("Carry:\t", system.ArithmeticLogicUnit.Flags.Carry)
	fmt.Println("Parity:\t", system.ArithmeticLogicUnit.Flags.Parity)
	fmt.Println()
	fmt.Println()

	// Restart system to ready another run.
	system.Restart()
	system.ProgramCounter.Value = types.DoubleWord(directives.PC)
}
