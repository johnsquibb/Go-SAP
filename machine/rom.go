package machine

import (
	"Go-SAP/machine/op"
	"Go-SAP/machine/types"
)

// AddressRomContents emulates a physical ROM integrated circuit.
// The keys correspond to the OpCodes, and the values correspond to the starting addresses
// sent to the PresettableCounter. Starting addresses should correspond to keys of ControlRomContents below.
// TODO This will eventually be read from file to allow for modification.
var AddressRomContents = types.AddressRom{}

// ControlRomContents emulates a physical ROM integrated circuit.
// The bit flags correspond to the states of the ControlWord flags, as indicated by "Active" comments below.
// TODO This will eventually be read from file to allow for modification.
var ControlRomContents = types.ControlRom{}

// FillRomContents initializes ControlRomContents and AddressRomContents with Microcode.
func FillRomContents() {
	ctr := 0

	// Apply fetch instruction to start of control ROM.
	// The system assumes 0 is the start of the fetch instruction.
	for _, instr := range op.FetchInstructions {
		ControlRomContents[ctr] = instr
		ctr++
	}

	// Apply all other instructions.
	for opcode, mc := range op.Microcodes {
		if mc != nil {
			// Set address ROM to point to starting location in control ROM.
			AddressRomContents[opcode] = types.DoubleWord(ctr)
			for _, instr := range mc {
				ControlRomContents[ctr] = instr
				ctr++
			}
		}
	}
}

// AddressReadOnlyMemory accepts input from InstructionRegister and outputs 8-bit addresses for use by the
// PresettableCounter. The PresettableCounter will tell the ControlReadOnlyMemory to start reading at an address, which
// is where the command ControlWord bits are found for the desired command.
// In hardware, the InstructionRegister.Address is hard-wired to this ROM.
type AddressReadOnlyMemory struct {
	Address types.Word
	Values  types.AddressRom
}

// Update sets the AddressReadOnlyMemory Address from InstructionRegister Instruction.
func (a *AddressReadOnlyMemory) Update(i InstructionRegister) {
	a.Address = i.Instruction
}

// Get returns the Value at current Address pointer for AddressReadOnlyMemory as Word.
func (a AddressReadOnlyMemory) Get() types.DoubleWord {
	return a.Values[a.Address]
}

// ControlReadOnlyMemory stores the microinstructions. The Values are translated to ControlWord values.
type ControlReadOnlyMemory struct {
	Address types.DoubleWord
	Values  types.ControlRom
}

// Update sets the ControlReadOnlyMemory Address to the PresettableCounter Value.
func (c *ControlReadOnlyMemory) Update(p PresettableCounter) {
	c.Address = p.Value
}

// Get returns the ControlReadOnlyMemory value at current address.
func (c ControlReadOnlyMemory) Get() types.OctupleWord {
	return c.Values[c.Address]
}
