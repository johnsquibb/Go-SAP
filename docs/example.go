package main

import (
	"Go-SAP3/machine"
	"Go-SAP3/machine/asm"
	"Go-SAP3/machine/types"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Sample Assembly source code.
// Assembly code is converted to machine code by Assemble method.
var source = `
ACI 0x11
ACI 0x11
ACI 0x11
ADC A
ADC A
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

	// Initialize ROM data from defaults established in op/microcode.go
	// system.InitRom()

	// Initialize ROM data previously exported to .dat files in this directory
	system.LoadRom(loadAddressRom(), loadControlRom())

	system.Start()
	system.ProgramCounter.Value = types.DoubleWord(directives.PC)

	for {
		system.Update()

		system.StepClock()

		if system.Halt {
			break
		}
	}

	// Export ROM content to data files in local directory.
	// Can be used to boostrap the system in subsequent runs (see above)
	// writeControlRomFile(system)
	// writeAddressRomFile(system)

	// Write RAM to file in local directory for debugging purposes.
	writeRamFile(system)

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

func writeRamFile(system machine.System) {
	f, err := os.Create("RAM.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for addr, val := range system.RandomAccessMemory.Values {
		_, err := f.WriteString(fmt.Sprintf("%X %X\n", addr, val))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func writeControlRomFile(system machine.System) {
	f, err := os.Create("controlROM.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for addr, val := range system.ControlReadOnlyMemory.Values {
		_, err := f.WriteString(fmt.Sprintf("%X %X\n", addr, val))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func writeAddressRomFile(system machine.System) {
	f, err := os.Create("addressROM.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for addr, val := range system.AddressReadOnlyMemory.Values {
		_, err := f.WriteString(fmt.Sprintf("%X %X\n", addr, val))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loadAddressRom() types.AddressRom {
	var addressRom = types.AddressRom{}

	for addr, value := range readRomFile("addressROM.dat") {
		addressRom[addr] = types.DoubleWord(value)
	}

	return addressRom
}

func loadControlRom() types.ControlRom {
	var controlRom = types.ControlRom{}

	for addr, value := range readRomFile("controlROM.dat") {
		controlRom[addr] = types.OctupleWord(value)
	}

	return controlRom
}

func readRomFile(filename string) []int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Each line format is: Address Value
		vals := strings.Split(scanner.Text(), " ")

		value, _ := strconv.ParseInt(vals[1], 16, 64)

		lines = append(lines, value)
	}

	return lines
}
