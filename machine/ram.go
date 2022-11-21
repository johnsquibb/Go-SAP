package machine

import (
	"Go-SAP/machine/types"
	"log"
)

// RamContents represents the RAM values input by assembler.
var RamContents = types.Ram{}

// ApplyInstructionsToRamContents sets instructions to RAM locations in sequential order.
func ApplyInstructionsToRamContents(instructions []types.Word, address int) {
	if len(instructions) > len(RamContents) {
		log.Fatal("Size of instructions exceeds available RAM")
	}

	for offset, instruction := range instructions {
		RamContents[address+offset] = instruction
	}
}

// RandomAccessMemory represents 64KB 8-bit-wide dynamic RAM.
// RandomAccessMemory is connected to MemoryAddressRegister for addressing, and MemoryDataRegister for value access.
type RandomAccessMemory struct {
	Address types.DoubleWord
	Values  types.Ram
}

// Get returns the Value at Address previously set by MemoryAddressRegister.
func (r RandomAccessMemory) Get() types.Word {
	return r.Values[r.Address]
}

// Set applies the Value at Address previously set by MemoryAddressRegister.
func (r *RandomAccessMemory) Set(w types.Word) {
	r.Values[r.Address] = w
}
