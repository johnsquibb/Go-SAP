package machine

import (
	"Go-SAP3/machine/op"
	"Go-SAP3/machine/types"
)

// ControllerSequencer provides instructions to each logic unit.
type ControllerSequencer struct {
	ControlWord ControlWord
}

// ControlWord maintains the state of the flags of the various logic units.
// It is the heart of the control system.
// In hardware, it represents the state of the wires that connect to each chip on the board.
// cmd is not emulating specific integrated circuits, so all states are High when on, Low when off.
// Note: The struct members are defined in ascending bit order, i.e. ProgramCounterClockEnable is LSB.
type ControlWord struct {
	ProgramCounterReadEnable           types.State
	ProgramCounterWriteEnable          types.State
	ProgramCounterClockEnable          types.State
	MemoryAddressRegisterReadEnable    types.State
	MemoryDataRegisterReadEnable       types.State
	MemoryDataRegisterWriteEnable      types.State
	MemoryAddressBufferLSBReadEnable   types.State
	MemoryAddressBufferLSBWriteEnable  types.State
	MemoryAddressBufferMSBReadEnable   types.State
	MemoryAddressBufferMSBWriteEnable  types.State
	GeneralPurposeBufferLSBReadEnable  types.State
	GeneralPurposeBufferLSBWriteEnable types.State
	GeneralPurposeBufferMSBReadEnable  types.State
	GeneralPurposeBufferMSBWriteEnable types.State
	InstructionRegisterReadEnable      types.State
	AccumulatorReadEnable              types.State
	AccumulatorWriteEnable             types.State
	TemporaryRegisterReadEnable        types.State
	TemporaryRegisterWriteEnable       types.State
	BRegisterReadEnable                types.State
	BRegisterWriteEnable               types.State
	BRegisterAccumulatorEnable         types.State
	CRegisterReadEnable                types.State
	CRegisterWriteEnable               types.State
	CRegisterAccumulatorEnable         types.State
	DRegisterReadEnable                types.State
	DRegisterWriteEnable               types.State
	DRegisterAccumulatorEnable         types.State
	ERegisterReadEnable                types.State
	ERegisterWriteEnable               types.State
	ERegisterAccumulatorEnable         types.State
	FRegisterReadEnable                types.State
	FRegisterWriteEnable               types.State
	HRegisterReadEnable                types.State
	HRegisterWriteEnable               types.State
	HRegisterAccumulatorEnable         types.State
	LRegisterReadEnable                types.State
	LRegisterWriteEnable               types.State
	LRegisterAccumulatorEnable         types.State
	StackPointerReadEnable             types.State
	StackPointerWriteEnable            types.State
	StackPointerAccumulatorEnable      types.State
	IORegisterPortSet                  types.State
	IORegisterReadEnable               types.State
	IORegisterWriteEnable              types.State
	ArithmeticLogicUnitWriteEnable     types.State
	ArithmeticLogicUnitModeBit0        types.State
	ArithmeticLogicUnitModeBit1        types.State
	ArithmeticLogicUnitModeBit2        types.State
	ArithmeticLogicUnitModeBit3        types.State
	ArithmeticLogicUnitModeBit4        types.State
	ZeroFlagEnable                     types.State
	SignFlagEnable                     types.State
	CarryFlagEnable                    types.State
	ParityFlagEnable                   types.State
	FlagsCheckCarryFlagIsEnabled       types.State
	FlagsCheckCarryFlagIsNotEnabled    types.State
	FlagsCheckSignFlagIsEnabled        types.State
	FlagsCheckSignFlagIsNotEnabled     types.State
	FlagsCheckZeroFlagIsEnabled        types.State
	FlagsCheckZeroFlagNotEnabled       types.State
	FlagsCheckParityFlagIsEnabled      types.State
	FlagsCheckParityFlagIsNotEnabled   types.State
}

// NewControlWord returns a ControlWord with all the flags set from supplied instruction QuadrupleWord.
func NewControlWord(mi types.OctupleWord) ControlWord {
	cw := ControlWord{}

	cw.ProgramCounterReadEnable = mi&op.PC_re != 0
	cw.ProgramCounterWriteEnable = mi&op.PC_we != 0
	cw.ProgramCounterClockEnable = mi&op.PC_ce != 0
	cw.MemoryAddressRegisterReadEnable = mi&op.MAR_re != 0
	cw.MemoryDataRegisterReadEnable = mi&op.MDR_re != 0
	cw.MemoryDataRegisterWriteEnable = mi&op.MDR_we != 0
	cw.MemoryAddressBufferLSBReadEnable = mi&op.MAB_lsbre != 0
	cw.MemoryAddressBufferLSBWriteEnable = mi&op.MAB_lsbwe != 0
	cw.MemoryAddressBufferMSBReadEnable = mi&op.MAB_msbre != 0
	cw.MemoryAddressBufferMSBWriteEnable = mi&op.MAB_msbwe != 0
	cw.GeneralPurposeBufferLSBReadEnable = mi&op.GPB_lsbre != 0
	cw.GeneralPurposeBufferLSBWriteEnable = mi&op.GPB_lsbwe != 0
	cw.GeneralPurposeBufferMSBReadEnable = mi&op.GPB_msbre != 0
	cw.GeneralPurposeBufferMSBWriteEnable = mi&op.GPB_msbwe != 0
	cw.InstructionRegisterReadEnable = mi&op.IR_re != 0
	cw.AccumulatorReadEnable = mi&op.ACC_re != 0
	cw.AccumulatorWriteEnable = mi&op.ACC_we != 0
	cw.TemporaryRegisterReadEnable = mi&op.TMP_re != 0
	cw.TemporaryRegisterWriteEnable = mi&op.TMP_we != 0
	cw.BRegisterReadEnable = mi&op.BREG_re != 0
	cw.BRegisterWriteEnable = mi&op.BREG_we != 0
	cw.BRegisterAccumulatorEnable = mi&op.BREG_accen != 0
	cw.CRegisterReadEnable = mi&op.CREG_re != 0
	cw.CRegisterWriteEnable = mi&op.CREG_we != 0
	cw.CRegisterAccumulatorEnable = mi&op.CREG_accen != 0
	cw.DRegisterReadEnable = mi&op.DREG_re != 0
	cw.DRegisterWriteEnable = mi&op.DREG_we != 0
	cw.DRegisterAccumulatorEnable = mi&op.DREG_accen != 0
	cw.ERegisterReadEnable = mi&op.EREG_re != 0
	cw.ERegisterWriteEnable = mi&op.EREG_we != 0
	cw.ERegisterAccumulatorEnable = mi&op.EREG_accen != 0
	cw.FRegisterReadEnable = mi&op.FREG_re != 0
	cw.FRegisterWriteEnable = mi&op.FREG_we != 0
	cw.HRegisterReadEnable = mi&op.HREG_re != 0
	cw.HRegisterWriteEnable = mi&op.HREG_we != 0
	cw.HRegisterAccumulatorEnable = mi&op.HREG_accen != 0
	cw.LRegisterReadEnable = mi&op.LREG_re != 0
	cw.LRegisterWriteEnable = mi&op.LREG_we != 0
	cw.LRegisterAccumulatorEnable = mi&op.LREG_accen != 0
	cw.StackPointerReadEnable = mi&op.SP_re != 0
	cw.StackPointerWriteEnable = mi&op.SP_we != 0
	cw.StackPointerAccumulatorEnable = mi&op.SP_accen != 0
	cw.ArithmeticLogicUnitWriteEnable = mi&op.ALU_we != 0
	cw.ArithmeticLogicUnitModeBit0 = mi&op.ALU_b0 != 0
	cw.ArithmeticLogicUnitModeBit1 = mi&op.ALU_b1 != 0
	cw.ArithmeticLogicUnitModeBit2 = mi&op.ALU_b2 != 0
	cw.ArithmeticLogicUnitModeBit3 = mi&op.ALU_b3 != 0
	cw.ArithmeticLogicUnitModeBit4 = mi&op.ALU_b4 != 0
	cw.SignFlagEnable = mi&op.FLG_sen != 0
	cw.CarryFlagEnable = mi&op.FLG_cen != 0
	cw.ZeroFlagEnable = mi&op.FLG_zen != 0
	cw.ParityFlagEnable = mi&op.FLG_pen != 0
	cw.FlagsCheckCarryFlagIsEnabled = mi&op.FLG_ccfie != 0
	cw.FlagsCheckCarryFlagIsNotEnabled = mi&op.FLG_ccfne != 0
	cw.FlagsCheckSignFlagIsEnabled = mi&op.FLG_csfie != 0
	cw.FlagsCheckSignFlagIsNotEnabled = mi&op.FLG_csfne != 0
	cw.FlagsCheckZeroFlagIsEnabled = mi&op.FLG_czfie != 0
	cw.FlagsCheckZeroFlagNotEnabled = mi&op.FLG_czfne != 0
	cw.FlagsCheckParityFlagIsEnabled = mi&op.FLG_cpfie != 0
	cw.FlagsCheckParityFlagIsNotEnabled = mi&op.FLG_cpfne != 0
	cw.IORegisterPortSet = mi&op.IO_ps != 0
	cw.IORegisterWriteEnable = mi&op.IO_we != 0
	cw.IORegisterReadEnable = mi&op.IO_re != 0

	return cw
}

// UpdateCtr tracks the number of complete executions of Update.
// This can be used to track machine cycle count.
var UpdateCtr = 0

// Update sets ControlWord flags to System units and executes the next T state.
func (c *ControllerSequencer) Update(s *System) {

	// ReadEnable control word.
	s.ControlReadOnlyMemory.Update(s.PresettableCounter)
	mi := s.ControlReadOnlyMemory.Get()
	c.ControlWord = NewControlWord(mi)

	// Set flags for each unit.
	s.ProgramCounter.ReadEnable = c.ControlWord.ProgramCounterReadEnable
	s.ProgramCounter.WriteEnable = c.ControlWord.ProgramCounterWriteEnable
	s.ProgramCounter.ClockEnable = c.ControlWord.ProgramCounterClockEnable

	s.MemoryAddressRegister.ReadEnable = c.ControlWord.MemoryAddressRegisterReadEnable

	s.MemoryDataRegister.ReadEnable = c.ControlWord.MemoryDataRegisterReadEnable
	s.MemoryDataRegister.WriteEnable = c.ControlWord.MemoryDataRegisterWriteEnable

	s.MemoryAddressBuffer.LSBReadEnable = c.ControlWord.MemoryAddressBufferLSBReadEnable
	s.MemoryAddressBuffer.LSBWriteEnable = c.ControlWord.MemoryAddressBufferLSBWriteEnable
	s.MemoryAddressBuffer.MSBReadEnable = c.ControlWord.MemoryAddressBufferMSBReadEnable
	s.MemoryAddressBuffer.MSBWriteEnable = c.ControlWord.MemoryAddressBufferMSBWriteEnable

	s.GeneralPurposeBuffer.LSBReadEnable = c.ControlWord.GeneralPurposeBufferLSBReadEnable
	s.GeneralPurposeBuffer.LSBWriteEnable = c.ControlWord.GeneralPurposeBufferLSBWriteEnable
	s.GeneralPurposeBuffer.MSBReadEnable = c.ControlWord.GeneralPurposeBufferMSBReadEnable
	s.GeneralPurposeBuffer.MSBWriteEnable = c.ControlWord.GeneralPurposeBufferMSBWriteEnable

	s.InstructionRegister.ReadEnable = c.ControlWord.InstructionRegisterReadEnable

	s.Accumulator.ReadEnable = c.ControlWord.AccumulatorReadEnable
	s.Accumulator.WriteEnable = c.ControlWord.AccumulatorWriteEnable

	s.TemporaryRegister.ReadEnable = c.ControlWord.TemporaryRegisterReadEnable
	s.TemporaryRegister.WriteEnable = c.ControlWord.TemporaryRegisterWriteEnable

	s.BRegister.ReadEnable = c.ControlWord.BRegisterReadEnable
	s.BRegister.WriteEnable = c.ControlWord.BRegisterWriteEnable

	s.CRegister.ReadEnable = c.ControlWord.CRegisterReadEnable
	s.CRegister.WriteEnable = c.ControlWord.CRegisterWriteEnable

	s.DRegister.ReadEnable = c.ControlWord.DRegisterReadEnable
	s.DRegister.WriteEnable = c.ControlWord.DRegisterWriteEnable

	s.ERegister.ReadEnable = c.ControlWord.ERegisterReadEnable
	s.ERegister.WriteEnable = c.ControlWord.ERegisterWriteEnable

	s.FRegister.ReadEnable = c.ControlWord.FRegisterReadEnable
	s.FRegister.WriteEnable = c.ControlWord.FRegisterWriteEnable

	s.StackPointer.ReadEnable = c.ControlWord.StackPointerReadEnable
	s.StackPointer.WriteEnable = c.ControlWord.StackPointerWriteEnable

	s.HRegister.ReadEnable = c.ControlWord.HRegisterReadEnable
	s.HRegister.WriteEnable = c.ControlWord.HRegisterWriteEnable

	s.LRegister.ReadEnable = c.ControlWord.LRegisterReadEnable
	s.LRegister.WriteEnable = c.ControlWord.LRegisterWriteEnable

	s.ArithmeticLogicUnit.WriteEnable = c.ControlWord.ArithmeticLogicUnitWriteEnable
	s.ArithmeticLogicUnit.AluMode.Bit0 = c.ControlWord.ArithmeticLogicUnitModeBit0
	s.ArithmeticLogicUnit.AluMode.Bit1 = c.ControlWord.ArithmeticLogicUnitModeBit1
	s.ArithmeticLogicUnit.AluMode.Bit2 = c.ControlWord.ArithmeticLogicUnitModeBit2
	s.ArithmeticLogicUnit.AluMode.Bit3 = c.ControlWord.ArithmeticLogicUnitModeBit3
	s.ArithmeticLogicUnit.AluMode.Bit4 = c.ControlWord.ArithmeticLogicUnitModeBit4
	s.ArithmeticLogicUnit.Flags.CarryEnable = c.ControlWord.CarryFlagEnable
	s.ArithmeticLogicUnit.Flags.SignEnable = c.ControlWord.SignFlagEnable
	s.ArithmeticLogicUnit.Flags.ZeroEnable = c.ControlWord.ZeroFlagEnable
	s.ArithmeticLogicUnit.Flags.ParityEnable = c.ControlWord.ParityFlagEnable

	// SAP can skip execution when Noop is received.
	// Since all control word flags are off by default, there is nothing to do.
	// Reset the ring counter and go to next instruction.
	if mi == op.Noop {
		s.DumpFormattedSystemDebugObject()
		s.RingCounter.Value = END
		return
	}

	switch s.RingCounter.Value {
	// Fetch Cycle
	case T1:
		c.AddressState(s)
	case T2:
		c.IncrementState(s)
	case T3:
		c.MemoryState(s)
	// Execution Cycle
	default:
		c.ExecutionState(s)
	}

	UpdateCtr++
}

// AddressState updates system state at T1.
func (c ControllerSequencer) AddressState(s *System) {
	// Set address into memory.
	s.ProgramCounter.Update(&s.Bus)
	s.MemoryAddressRegister.Update(s.Bus, &s.RandomAccessMemory)
}

// IncrementState updates system state at T2.
func (c ControllerSequencer) IncrementState(s *System) {
	// Increment program counter.
	s.ProgramCounter.Update(&s.Bus)
}

// MemoryState updates system state at T3.
func (c ControllerSequencer) MemoryState(s *System) {
	// Transfer instruction from memory to instruction register.
	s.MemoryDataRegister.Update(&s.Bus, &s.RandomAccessMemory)
	s.InstructionRegister.Update(&s.Bus)

	// If there is a Halt instruction, this is the point at which honor it.
	s.InstructionRegister.Halt(s)

	// Set starting address for ControllerSequencer from Address ROM.
	s.AddressReadOnlyMemory.Update(s.InstructionRegister)
}

// ExecutionState updates system state at T4 - END
func (c ControllerSequencer) ExecutionState(s *System) {
	c.UpdateALU(s)
	c.CheckConditionFlags(s)

	WriteRegisters(s)
	ReadRegisters(s)

	c.UpdateInputs(s)
	c.UpdateOutputs(s)
}

// CheckConditionFlags checks for state of ArithmeticLogicUnit flags during conditional jump/call/return executions.
// When conditions are met, the controller will skip to the END to ensure the destination is not addressed.
func (c ControllerSequencer) CheckConditionFlags(s *System) {

	// JC/CC/RC
	if c.ControlWord.FlagsCheckCarryFlagIsEnabled == types.High && s.ArithmeticLogicUnit.Flags.Carry == types.Low {
		s.RingCounter.Value = END
		return
	}

	// JNC/CNC/RNC
	if c.ControlWord.FlagsCheckCarryFlagIsNotEnabled == types.High && s.ArithmeticLogicUnit.Flags.Carry == types.High {
		s.RingCounter.Value = END
		return
	}

	// JM/CM/RM
	if c.ControlWord.FlagsCheckSignFlagIsEnabled == types.High && s.ArithmeticLogicUnit.Flags.Sign == types.Low {
		s.RingCounter.Value = END
		return
	}

	// JP/CP/RP
	if c.ControlWord.FlagsCheckSignFlagIsNotEnabled == types.High && s.ArithmeticLogicUnit.Flags.Sign == types.High {
		s.RingCounter.Value = END
		return
	}

	// JZ/CZ/RZ
	if c.ControlWord.FlagsCheckZeroFlagIsEnabled == types.High && s.ArithmeticLogicUnit.Flags.Zero == types.Low {
		s.RingCounter.Value = END
		return
	}

	// JNZ/CNZ/RNZ
	if c.ControlWord.FlagsCheckZeroFlagNotEnabled == types.High && s.ArithmeticLogicUnit.Flags.Zero == types.High {
		s.RingCounter.Value = END
		return
	}

	// JPE/CPE/RPE
	if c.ControlWord.FlagsCheckParityFlagIsEnabled == types.High && s.ArithmeticLogicUnit.Flags.Parity == types.Low {
		s.RingCounter.Value = END
		return
	}

	// JPO/CPO/RPO
	if c.ControlWord.FlagsCheckParityFlagIsNotEnabled == types.High && s.ArithmeticLogicUnit.Flags.Parity == types.High {
		s.RingCounter.Value = END
		return
	}
}

// UpdateOutputs updates all output registers.
func (c ControllerSequencer) UpdateOutputs(s *System) {
	s.OutputRegister3.ReadEnable = types.Low
	s.OutputRegister4.ReadEnable = types.Low

	if c.ControlWord.IORegisterPortSet {
		switch types.Word(s.Bus.Value) {
		case 0b0011:
			s.OutputRegister3.ReadEnable = types.High
		case 0b0100:
			s.OutputRegister4.ReadEnable = types.High
		}
	}

	if s.OutputRegister3.ReadEnable == types.High {
		s.OutputRegister3.Update(&s.Bus)
	}

	if s.OutputRegister4.ReadEnable == types.High {
		s.OutputRegister4.Update(&s.Bus)
	}
}

// UpdateInputs updates all input registers.
func (c ControllerSequencer) UpdateInputs(s *System) {
	s.InputRegister1.WriteEnable = types.Low
	s.InputRegister2.WriteEnable = types.Low

	if c.ControlWord.IORegisterPortSet {
		switch types.Word(s.Bus.Value) {
		case 0b0001:
			s.InputRegister1.WriteEnable = types.High
		case 0b0010:
			s.InputRegister2.WriteEnable = types.High
		}
	}

	if s.InputRegister1.WriteEnable == types.High {
		s.InputRegister1.Update(&s.Bus)
	}

	if s.InputRegister2.WriteEnable == types.High {
		s.InputRegister2.Update(&s.Bus)
	}
}

// UpdateALU updates the ArithmeticLogicUnit from values in registers.
func (c ControllerSequencer) UpdateALU(s *System) {
	// BC pair.
	if c.ControlWord.BRegisterAccumulatorEnable == types.High &&
		c.ControlWord.CRegisterAccumulatorEnable == types.High {
		s.ArithmeticLogicUnit.Update(s.BRegister, s.CRegister, s)
		s.BRegister.Value = s.ArithmeticLogicUnit.A
		s.CRegister.Value = s.ArithmeticLogicUnit.B
		return
	}

	// DE Pair.
	if c.ControlWord.DRegisterAccumulatorEnable == types.High &&
		c.ControlWord.ERegisterAccumulatorEnable == types.High {
		s.ArithmeticLogicUnit.Update(s.DRegister, s.ERegister, s)
		s.DRegister.Value = s.ArithmeticLogicUnit.A
		s.ERegister.Value = s.ArithmeticLogicUnit.B
		return
	}

	// HL Pair.
	if c.ControlWord.HRegisterAccumulatorEnable == types.High &&
		c.ControlWord.LRegisterAccumulatorEnable == types.High {
		s.ArithmeticLogicUnit.Update(s.HRegister, s.LRegister, s)
		s.HRegister.Value = s.ArithmeticLogicUnit.A
		s.LRegister.Value = s.ArithmeticLogicUnit.B
		return
	}

	// StackPointer has built-in increment/decrement operations.
	if c.ControlWord.StackPointerAccumulatorEnable == types.High {
		// Increment
		if c.ControlWord.ArithmeticLogicUnitModeBit4 == types.High &&
			c.ControlWord.ArithmeticLogicUnitModeBit0 == types.High {
			s.StackPointer.Increment()
			return
		}

		// Decrement
		if c.ControlWord.ArithmeticLogicUnitModeBit4 == types.High &&
			c.ControlWord.ArithmeticLogicUnitModeBit1 == types.High {
			s.StackPointer.Decrement()
			return
		}

		return
	}

	// B
	if c.ControlWord.BRegisterAccumulatorEnable == types.High {
		s.ArithmeticLogicUnit.Update(s.BRegister, s.TemporaryRegister, s)
		return
	}

	// C
	if c.ControlWord.CRegisterAccumulatorEnable == types.High {
		s.ArithmeticLogicUnit.Update(s.CRegister, s.TemporaryRegister, s)
		return
	}

	// D
	if c.ControlWord.DRegisterAccumulatorEnable == types.High {
		s.ArithmeticLogicUnit.Update(s.DRegister, s.TemporaryRegister, s)
		return
	}

	// E
	if c.ControlWord.ERegisterAccumulatorEnable == types.High {
		s.ArithmeticLogicUnit.Update(s.ERegister, s.TemporaryRegister, s)
		return
	}

	// H
	if c.ControlWord.HRegisterAccumulatorEnable == types.High {
		s.ArithmeticLogicUnit.Update(s.HRegister, s.TemporaryRegister, s)
		return
	}

	// L
	if c.ControlWord.LRegisterAccumulatorEnable == types.High {
		s.ArithmeticLogicUnit.Update(s.LRegister, s.TemporaryRegister, s)
		return
	}

	s.ArithmeticLogicUnit.Update(s.Accumulator, s.TemporaryRegister, s)
}

// WriteRegisters updates all registers in order to perform write operations.
func WriteRegisters(s *System) {

	if s.ProgramCounter.WriteEnable || s.ProgramCounter.ClockEnable == types.High {
		s.ProgramCounter.Update(&s.Bus)
	}

	if s.Accumulator.WriteEnable == types.High {
		s.Accumulator.Update(&s.Bus)
	}

	if s.BRegister.WriteEnable == types.High {
		s.BRegister.Update(&s.Bus)
	}

	if s.CRegister.WriteEnable == types.High {
		s.CRegister.Update(&s.Bus)
	}

	if s.DRegister.WriteEnable == types.High {
		s.DRegister.Update(&s.Bus)
	}

	if s.ERegister.WriteEnable == types.High {
		s.ERegister.Update(&s.Bus)
	}

	if s.FRegister.WriteEnable == types.High {
		s.FRegister.Update(&s.Bus)
	}

	if s.HRegister.WriteEnable == types.High {
		s.HRegister.Update(&s.Bus)
	}

	if s.LRegister.WriteEnable == types.High {
		s.LRegister.Update(&s.Bus)
	}

	if s.StackPointer.WriteEnable == types.High {
		s.StackPointer.Update(&s.Bus)
	}

	if s.TemporaryRegister.WriteEnable == types.High {
		s.TemporaryRegister.Update(&s.Bus)
	}

	if s.InputRegister1.WriteEnable == types.High {
		s.InputRegister1.Update(&s.Bus)
	}

	if s.InputRegister2.WriteEnable == types.High {
		s.InputRegister2.Update(&s.Bus)
	}

	if s.MemoryDataRegister.WriteEnable == types.High {
		s.MemoryDataRegister.Update(&s.Bus, &s.RandomAccessMemory)
	}

	if s.MemoryAddressBuffer.MSBWriteEnable || s.MemoryAddressBuffer.LSBWriteEnable {
		s.MemoryAddressBuffer.Update(&s.Bus)
	}

	if s.GeneralPurposeBuffer.MSBWriteEnable || s.GeneralPurposeBuffer.LSBWriteEnable {
		s.GeneralPurposeBuffer.Update(&s.Bus)
	}
}

// ReadRegisters updates all registers in order to perform read operations.
func ReadRegisters(s *System) {
	if s.ProgramCounter.ReadEnable == types.High {
		s.ProgramCounter.Update(&s.Bus)
	}

	if s.Accumulator.ReadEnable == types.High {
		s.Accumulator.Update(&s.Bus)
	}

	if s.BRegister.ReadEnable == types.High {
		s.BRegister.Update(&s.Bus)
	}

	if s.CRegister.ReadEnable == types.High {
		s.CRegister.Update(&s.Bus)
	}

	if s.DRegister.ReadEnable == types.High {
		s.DRegister.Update(&s.Bus)
	}

	if s.ERegister.ReadEnable == types.High {
		s.ERegister.Update(&s.Bus)
	}

	if s.FRegister.ReadEnable == types.High {
		s.FRegister.Update(&s.Bus)
	}

	if s.HRegister.ReadEnable == types.High {
		s.HRegister.Update(&s.Bus)
	}

	if s.LRegister.ReadEnable == types.High {
		s.LRegister.Update(&s.Bus)
	}

	if s.StackPointer.ReadEnable == types.High {
		s.StackPointer.Update(&s.Bus)
	}

	if s.TemporaryRegister.ReadEnable == types.High {
		s.TemporaryRegister.Update(&s.Bus)
	}

	if s.OutputRegister3.ReadEnable == types.High {
		s.OutputRegister3.Update(&s.Bus)
	}

	if s.OutputRegister4.ReadEnable == types.High {
		s.OutputRegister4.Update(&s.Bus)
	}

	if s.MemoryAddressRegister.ReadEnable == types.High {
		s.MemoryAddressRegister.Update(s.Bus, &s.RandomAccessMemory)
	}

	if s.MemoryDataRegister.ReadEnable == types.High {
		s.MemoryDataRegister.Update(&s.Bus, &s.RandomAccessMemory)
	}

	if s.MemoryAddressBuffer.MSBReadEnable || s.MemoryAddressBuffer.LSBReadEnable {
		s.MemoryAddressBuffer.Update(&s.Bus)
	}

	if s.GeneralPurposeBuffer.MSBReadEnable || s.GeneralPurposeBuffer.LSBReadEnable {
		s.GeneralPurposeBuffer.Update(&s.Bus)
	}
}
