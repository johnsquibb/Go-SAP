package machine

import "Go-SAP3/machine/types"

// T1 thru END states are controlled by RingCounter clock.
// T1 - T3 represent the periods of the Fetch Cycle.
// T4+ represent the periods of the variable Execution Cycle.
const (
	T1 = iota // Address State
	T2        // Increment State
	T3        // Memory State
	T4        // Execution Cycle 1
	T5
	T6
	T7
	T8
	T9
	T10
	T11
	T12
	T13
	T14
	T15
	T16
	T17
	END = T17
)

// RingCounter handles the timing states for each program cycle.
// T1 - T3 are the fetch cycle, T4 - END are the execution cycles.
type RingCounter struct {
	Value types.QuadrupleWord
}

// System clears PresettableCounter during T1.
// Update sets PresettableCounter ReadEnable during T3.
func (r *RingCounter) Update(p *PresettableCounter) {
	p.SetEnable = r.Value == T3
}

// Clear resets the RingCounter state to the first Time state.
func (r *RingCounter) Clear() {
	r.Value = T1
}

// Clock increments the RingCounter State.
func (r *RingCounter) Clock() {
	r.Value++
	// Hard-wired to 18 states.
	if r.Value > END {
		r.Clear()
	}
}

// PresettableCounter starts counting from a set position. When Clear() is called, this value is 0.
// ControllerSequencer sets the starting address from AddressReadOnlyMemory when T3 is High. During the other T states,
// the counter counts, providing the address inputs to ControlReadOnlyMemory.
// Setting of the counter is only possible when SetEnable is High.
type PresettableCounter struct {
	Value     types.DoubleWord
	SetEnable types.State
}

// Update sets PresettableCounter Value to value from AddressReadOnlyMemory when SetEnable is High.
func (p *PresettableCounter) Update(a AddressReadOnlyMemory) {
	if p.SetEnable == types.High {
		p.Value = a.Get()
	}
}

// Clock increments PresettableCounter Value only when SetEnable is Low.
func (p *PresettableCounter) Clock() {
	if p.SetEnable == types.Low {
		p.Value++
	}
}

// Clear resets PresettableCounter Value.
func (p *PresettableCounter) Clear() {
	p.Value = 0
}

// ProgramCounter is the pointer to instructions in the MemoryAddressRegister.
// The counter can only be incremented when ClockEnable is High.
// The Value is set to output when WriteEnable is High.
type ProgramCounter struct {
	Value       types.DoubleWord
	ClockEnable types.State
	ReadEnable  types.State
	WriteEnable types.State
}

// Get returns the ProgramCounter value as Word.
func (p ProgramCounter) Get() types.DoubleWord {
	return p.Value
}

// Update reads ProgramCounter Value from Bus when ReadEnable is High.
// Update writes ProgramCounter Value to Bus when WriteEnable is High.
// Update increments ProgramCounter Value when ClockEnable is High.
func (p *ProgramCounter) Update(b *Bus) {
	if p.ReadEnable == types.High {
		p.Value = b.Value
	}

	if p.WriteEnable == types.High {
		b.Value = p.Value
	}

	if p.ClockEnable == types.High {
		p.Increment()
	}
}

// Increment increments ProgramCounter Value.
func (p *ProgramCounter) Increment() {
	p.Value++
}

// Clear resets the clock Value to 0.
func (p *ProgramCounter) Clear() {
	p.Value = 0
}
