package machine

import (
	"SystemOnePoc/sap3/machine/types"
)

// MemoryDataRegister controls the memory I/O.
// During execution, the address in the ProgramCounter is latched into the MemoryAddressRegister.
// During execution, MemoryDataRegister reads the value at that address from RandomAccessMemory.
// that address.
type MemoryDataRegister struct {
	Value       types.Word
	ReadEnable  types.State
	WriteEnable types.State
}

// Update reads MemoryDataRegister Value from Bus when ReadEnable is High.
// Update writes MemoryDataRegister Value to Bus when WriteEnable is High.
func (m *MemoryDataRegister) Update(b *Bus, r *RandomAccessMemory) {
	if m.ReadEnable == types.High {
		m.Value = types.Word(b.Value)
		r.Set(m.Value)
	}

	if m.WriteEnable == types.High {
		m.Value = r.Get()
		b.Value = types.DoubleWord(m.Value)
	}
}