package machine

import (
	"Go-SAP/machine/op"
	"Go-SAP/machine/types"
)

// InstructionRegister is part of the control unit.
// The Instruction Word is the location in AddressReadOnlyMemory that contains the start address for PresettableCounter,
// which tells the ControlReadOnlyMemory which instruction to run. In hardware, this is hardwired directly to the ROM.
type InstructionRegister struct {
	Instruction types.Word
	ReadEnable  types.State
}

// Get returns the full DoubleWord representation of the Instruction:MemoryAddress for display.
func (i InstructionRegister) Get() types.DoubleWord {
	return types.DoubleWord(i.Instruction)
}

// Update reads Instruction from Bus when ReadEnable is High.
func (i *InstructionRegister) Update(b *Bus) {
	if i.ReadEnable == types.High {
		i.Instruction = types.Word(b.Value)
	}
}

// Clear resets the Address and Bus values for InstructionRegister.
func (i *InstructionRegister) Clear() {
	i.Instruction = 0
}

// Halt sets System Halt High when Instruction is HLT.
func (i *InstructionRegister) Halt(s *System) {
	s.Halt = i.Instruction == op.HLT
}
