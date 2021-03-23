// cmd
// See page 195 Digital Computer Electronics, Third Edition by Malvino & Brown
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
#dir ADDR 0x2000    ; The starting address for instructions.
#dir PC 0x2000      ; The initial value of Program Counter after loading program.

; --------------------------------------------------------------------------
; Fibonacci
; https://www.tutorialspoint.com/8085-program-to-generate-fibonacci-sequence
; --------------------------------------------------------------------------

START:  LXI H,0x8050    ; Pointer to the out buffer
        XRA A           ; Clear accumulator and reg. B
        MOV B,A
        MOV M,A         ; Copying content to target location
        INR A           ; Increment A
        INX H           ; Go the the next destination address.
        MOV M,A         ; Moving the content
        MVI C,0x08      ; Initialize counter
LOOP:   ADD B           ; Getting next term
        MOV B,M         ; Initializing term, e.g. F1 = F2
        INX H           ; Go to next destination address.
        MOV M,A         ; Writing to the out buffer
        DCR C           ; Decrement count until 0 is reached F3= F1 + F2 (A) = (A) + (B)
                        ; This is done with instruction ADDB.
        JNZ LOOP
        HLT             ; Terminate program
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
	showRAM(system, 0x8050, 0x8059)
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
