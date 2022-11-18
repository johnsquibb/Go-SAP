package test

import (
	"Go-SAP3/machine"
	"Go-SAP3/machine/asm"
	"log"
	"testing"
)

func getSystem(source string) machine.System {
	instructions, directives, err := asm.Assemble(source)
	if err != nil {
		log.Fatal(err)
	}

	//log.Fatal(instructions)

	// Apply assembled instructions to RAM.
	// RAM contents will be read during system.Start()
	machine.ApplyInstructionsToRamContents(instructions, directives.Address)

	// Run the computer.
	system := machine.New()
	system.InitRom()
	system.Start()

	return system
}

func runSystem(system machine.System) machine.System {
	for {
		system.Update()
		system.StepClock()
		if system.Halt {
			break
		}
	}

	return system
}

func check(t *testing.T, name string, want interface{}, got interface{}) {
	if got != want {
		t.Errorf("%s: Expected %v got %v", name, want, got)
	}
}
