package machine

import "Go-SAP/machine/types"

// StackPointer keeps track of the stack pointer.
type StackPointer struct {
	Address     types.DoubleWord
	ReadEnable  types.State
	WriteEnable types.State
}

// Update reads Address from Bus when ReadEnable is High.
// Update writes Address to Bus when WriteEnable is High.
func (sp *StackPointer) Update(b *Bus) {
	if sp.ReadEnable == types.High {
		sp.Address = b.Value
	}

	if sp.WriteEnable == types.High {
		b.Value = sp.Address
	}
}

// Increment increments the StackPointer by 1.
func (sp *StackPointer) Increment() {
	sp.Address++
}

// Decrement decrements the StackPointer by 1.
func (sp *StackPointer) Decrement() {
	sp.Address--
}
