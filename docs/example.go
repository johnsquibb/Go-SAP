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
; -------------------------
; CALL example with labels.
; -------------------------

CALL INIT
CALL INC4
CALL DEC4
CALL INC4

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

; --------------------------------
; Listing of Assembler directives.
;---------------------------------

#dir ADDR 0x8000    ; The starting address for instructions.
#dir PC 0x8000      ; The initial value of Program Counter after loading program.
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
	system.InitRom()

	// Initialize ROM data previously exported to .dat files in this directory
	// system.LoadRom(loadAddressRom(), loadControlRom())

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
	// writeRamFile(system)

	// Display the total number of machine cycles used.
	fmt.Println()
	fmt.Println(fmt.Sprintf("Machine Cycles:\t %d", machine.UpdateCtr))
	fmt.Println()

	// Display the contents of the various registers.
	fmt.Println(fmt.Sprintf("REG A:\t %02X\t", system.Accumulator.Value))
	fmt.Println(fmt.Sprintf("REG B:\t %02X\t", system.BRegister.Value))
	fmt.Println(fmt.Sprintf("REG C:\t %02X\t", system.CRegister.Value))
	fmt.Println(fmt.Sprintf("REG D:\t %02X\t", system.DRegister.Value))
	fmt.Println(fmt.Sprintf("REG E:\t %02X\t", system.ERegister.Value))
	fmt.Println(fmt.Sprintf("REG F:\t %02X\t", system.FRegister.Value))
	fmt.Println(fmt.Sprintf("REG H:\t %02X\t", system.HRegister.Value))
	fmt.Println(fmt.Sprintf("REG L:\t %02X\t", system.LRegister.Value))
	fmt.Println(fmt.Sprintf("SP:\t %02X\t", system.StackPointer.Address))
	fmt.Println(fmt.Sprintf("IO (3):\t %02X\t", system.OutputRegister3.Value))
	fmt.Println(fmt.Sprintf("IO (4):\t %02X\t", system.OutputRegister4.Value))
	fmt.Println()

	// Display the state of ALU flags.
	fmt.Println("Sign:\t", system.ArithmeticLogicUnit.Flags.Sign)
	fmt.Println("Zero:\t", system.ArithmeticLogicUnit.Flags.Zero)
	fmt.Println("Carry:\t", system.ArithmeticLogicUnit.Flags.Carry)
	fmt.Println("Parity:\t", system.ArithmeticLogicUnit.Flags.Parity)
	fmt.Println()
	fmt.Println()

	showRamRange(system, 0x8000, 16)
	fmt.Println()
	fmt.Println()

	// Restart system to ready another run.
	system.Restart()
	system.ProgramCounter.Value = types.DoubleWord(directives.PC)
}

func showRamRange(system machine.System, start int, lines int) {
	for i := start; i < start+lines; i++ {
		fmt.Println(fmt.Sprintf("RAM (%02X-%02X):  %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X\t",
			i,
			i+16,
			system.RandomAccessMemory.Values[i],
			system.RandomAccessMemory.Values[i+1],
			system.RandomAccessMemory.Values[i+2],
			system.RandomAccessMemory.Values[i+3],
			system.RandomAccessMemory.Values[i+4],
			system.RandomAccessMemory.Values[i+5],
			system.RandomAccessMemory.Values[i+6],
			system.RandomAccessMemory.Values[i+7],
			system.RandomAccessMemory.Values[i+8],
			system.RandomAccessMemory.Values[i+9],
			system.RandomAccessMemory.Values[i+10],
			system.RandomAccessMemory.Values[i+11],
			system.RandomAccessMemory.Values[i+12],
			system.RandomAccessMemory.Values[i+13],
			system.RandomAccessMemory.Values[i+14],
			system.RandomAccessMemory.Values[i+15],
		))
	}
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
