package main

import (
	"Go-SAP3/machine"
	"Go-SAP3/machine/asm"
	"Go-SAP3/machine/types"
	"fmt"
	"log"
	"os"
)

// Sample Assembly source code.
// Assembly code is converted to machine code by Assemble method.
var source = `
LXI H,0x1234
LXI D,0x5678
XCHG
HLT
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

	dumpRam(system)

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
	fmt.Println(fmt.Sprintf("SP:\t 0x%X\t (%d)", system.StackPointer.Address, system.StackPointer.Address))
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

	fmt.Println(fmt.Sprintf("RAM ADDR (8000):\t 0x%X\t", system.RandomAccessMemory.Values[0x8000]))
	fmt.Println(fmt.Sprintf("RAM ADDR (8001):\t 0x%X\t", system.RandomAccessMemory.Values[0x8001]))
	fmt.Println(fmt.Sprintf("RAM ADDR (8002):\t 0x%X\t", system.RandomAccessMemory.Values[0x8002]))
	fmt.Println(fmt.Sprintf("RAM ADDR (8003):\t 0x%X\t", system.RandomAccessMemory.Values[0x8003]))
	fmt.Println(fmt.Sprintf("RAM ADDR (8050):\t 0x%X\t", system.RandomAccessMemory.Values[0x8050]))
	fmt.Println(fmt.Sprintf("RAM ADDR (8051):\t 0x%X\t", system.RandomAccessMemory.Values[0x8051]))

	// Restart system to ready another run.
	system.Restart()
	system.ProgramCounter.Value = types.DoubleWord(directives.PC)

}

func dumpRam(system machine.System) {
	// create file
	f, err := os.Create("RAM.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	// Write memory to file for debugging purposes.
	for addr, val := range system.RandomAccessMemory.Values {
		_, err := f.WriteString(fmt.Sprintf("%X %X\n", addr, val))
		if err != nil {
			log.Fatal(err)
		}
	}
}
