package machine

import "SystemOnePoc/sap3/machine/types"

// MemoryAddressRegister controls the memory I/O.
// During execution, the address in the ProgramCounter is latched into the MAR. Later MAR reads the value from RAM at
// that address.
type MemoryAddressRegister struct {
	Address    types.DoubleWord
	ReadEnable types.State
}

// Update reads Address from Bus when ReadEnable is High.
func (m *MemoryAddressRegister) Update(b Bus, r *RandomAccessMemory) {
	if m.ReadEnable == types.High {
		m.Address = b.Value

		// MemoryAddressRegister Address is hardwired to RandomAccessMemory.
		r.Address = m.Address
	}
}
