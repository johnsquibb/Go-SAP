package machine

// StepClock increments the various clocks that make the System "tick".
// This can be used to manually advance the clock in STEP mode, otherwise Clock() should be used.
// A HLT Instruction in InstructionRegister causes the system to stop the clock.
func (s *System) StepClock() {
	if s.RingCounter.Value == END {
		s.Reset()
		return
	}

	// Update counters states.
	s.RingCounter.Update(&s.PresettableCounter)
	s.PresettableCounter.Update(s.AddressReadOnlyMemory)

	if !s.Halt {
		s.RingCounter.Clock()
		s.PresettableCounter.Clock()
	}
}
