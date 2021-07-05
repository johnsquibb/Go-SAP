// Demo of SAP-3.
// Inspired by Digital Computer Electronics, Third Edition by Malvino & Brown
// Emulated by John Squibb, 2021
package main

import (
	"Go-SAP3/machine"
	"Go-SAP3/machine/asm"
	"Go-SAP3/machine/types"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"time"
)

var ticks uint64

// MhzDelay is the amount of time to delay per step to approximate 1MHz clock speed.
var MhzDelay uint64

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
`

func main() {
	RunInteractiveBenchmark()
}

func RunInteractiveBenchmark() {
	instructions, directives, err := asm.Assemble(source)
	if err != nil {
		log.Fatal(err)
	}

	// Apply assembled instructions to RAM.
	// RAM contents will be read during system.Start()
	machine.ApplyInstructionsToRamContents(instructions, directives.Address)

	// Benchmark printer.
	p := message.NewPrinter(language.English)

	// Enable debug to see state at each step as JSON.
	// RAM is intentionally omitted due to size, but may be included by using system.DebugFull() instead.
	machine.EnableDebug(false)

	// Create new system, set the program counter per directive.
	system := machine.New()
	system.Start()
	system.ProgramCounter.Value = types.DoubleWord(directives.PC)

	// Manually set a value in RAM for our prime number example.
	// After execution, 0xF101 will contain the number of unique factors for the number. 2 factors means prime.
	// I kind of want to anonymously submit this for the code challenge answer at work.
	system.RandomAccessMemory.Values[0xF100] = 97

	// Start ticking the clock, try to ascertain delay for desired clock speed.
	startClock()
	MhzDelay = calculateTicksPerMicrosecond()

	// Benchmark some runs.
	start := time.Now()
	startTicks := ticks

	for {
		system.Update()

		system.StepClock()

		// Approximately 1MHz clock speed.
		ctr := MhzDelay * 1
		for ctr > 0 {
			ctr--
		}

		if system.Halt {
			break
		}
	}

	end := time.Since(start)
	endTicks := ticks - startTicks

	// Benchmarks.
	fmt.Println()
	fmt.Println()
	fmt.Println("Time:")
	fmt.Println(p.Sprintf("\t %d ms", end.Milliseconds()))
	fmt.Println(p.Sprintf("\t %d µs", end.Microseconds()))
	fmt.Println(p.Sprintf("\t %d ns", end.Nanoseconds()))
	fmt.Println(p.Sprintf("\t %d ticks", endTicks))
	fmt.Println(p.Sprintf("\t using %d ticks per µs", MhzDelay))
	fmt.Println()
	fmt.Println(p.Sprintf("Machine Cycles:\t %d", machine.UpdateCtr))
	fmt.Println()

	fmt.Println(fmt.Sprintf("REG A:\t 0x%X\t (%d)", system.Accumulator.Value, system.Accumulator.Value))
	fmt.Println(fmt.Sprintf("REG B:\t 0x%X\t (%d)", system.BRegister.Value, system.BRegister.Value))
	fmt.Println(fmt.Sprintf("REG C:\t 0x%X\t (%d)", system.CRegister.Value, system.CRegister.Value))
	fmt.Println(fmt.Sprintf("REG D:\t 0x%X\t (%d)", system.DRegister.Value, system.DRegister.Value))
	fmt.Println(fmt.Sprintf("REG E:\t 0x%X\t (%d)", system.ERegister.Value, system.ERegister.Value))
	fmt.Println(fmt.Sprintf("REG F:\t 0x%X\t (%d)", system.FRegister.Value, system.FRegister.Value))
	fmt.Println(fmt.Sprintf("REG H:\t 0x%X\t (%d)", system.HRegister.Value, system.HRegister.Value))
	fmt.Println(fmt.Sprintf("REG L:\t 0x%X\t (%d)", system.LRegister.Value, system.LRegister.Value))
	fmt.Println()

	fmt.Println("Sign:\t", system.ArithmeticLogicUnit.Flags.Sign)
	fmt.Println("Zero:\t", system.ArithmeticLogicUnit.Flags.Zero)
	fmt.Println("Carry:\t", system.ArithmeticLogicUnit.Flags.Carry)
	fmt.Println("Parity:\t", system.ArithmeticLogicUnit.Flags.Parity)
	fmt.Println()
	fmt.Println()

	// Examine a range of RAM.
	showRAM(system, 0xF100, 0xF101)
	fmt.Println()
	fmt.Println()

	// Show JSON dump of system state after run.
	//machine.EnableDebug(true)
	//system.DumpFormattedSystemDebugObject()

	// Restart system.
	system.Restart()
	system.ProgramCounter.Value = types.DoubleWord(directives.PC)
}

func showRAM(s machine.System, start, end int) {
	for i := start; i <= end; i++ {
		v := s.RandomAccessMemory.Values[i]
		fmt.Println(fmt.Sprintf("RAM (%X):\t 0x%X\t (%d)", i, v, v))
	}
}

func startClock() {
	go func(t *uint64) {
		for {
			*t++
			if *t > 1<<63 {
				*t = 0
			}
		}
	}(&ticks)
}

// calculateTicksPerMicrosecond calculates the number of ticks per microsecond.
func calculateTicksPerMicrosecond() uint64 {
	var its = uint64(1_000_000_000)
	s := time.Now()
	for i := uint64(0); i < its; i++ {
		// no-op
	}
	e := time.Now().Sub(s).Microseconds()

	microseconds := float64(e)
	ticksPerMicrosecond := float64(ticks) / microseconds

	// ticksPerMicrosecond should represent the approximate number of ticks to approximate 1MHz clock speed.
	return uint64(ticksPerMicrosecond)
}
