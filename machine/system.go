package machine

import "Go-SAP/machine/types"

// New returns an initialized System structure.
func New() System {
	return System{
		Bus:                   Bus{},
		RingCounter:           RingCounter{},
		PresettableCounter:    PresettableCounter{},
		ProgramCounter:        ProgramCounter{},
		StackPointer:          StackPointer{},
		MemoryAddressBuffer:   DoubleWordBuffer{},
		GeneralPurposeBuffer:  DoubleWordBuffer{},
		MemoryAddressRegister: MemoryAddressRegister{},
		MemoryDataRegister:    MemoryDataRegister{},
		RandomAccessMemory:    RandomAccessMemory{},
		AddressReadOnlyMemory: AddressReadOnlyMemory{},
		ControlReadOnlyMemory: ControlReadOnlyMemory{},
		InstructionRegister:   InstructionRegister{},
		ArithmeticLogicUnit:   ArithmeticLogicUnit{},
		FRegister:             Register{},
		Accumulator:           Register{},
		TemporaryRegister:     Register{},
		BRegister:             Register{},
		CRegister:             Register{},
		DRegister:             Register{},
		ERegister:             Register{},
		HRegister:             Register{},
		LRegister:             Register{},
		InputRegister1:        Register{},
		InputRegister2:        Register{},
		OutputRegister3:       Register{},
		OutputRegister4:       Register{},
		ControllerSequencer:   ControllerSequencer{},
	}
}

// System represents the various logical units.
// In hardware, each unit of System would comprise a variety of TTL and other integrated circuits.
type System struct {
	Bus                   Bus
	RingCounter           RingCounter
	PresettableCounter    PresettableCounter
	ProgramCounter        ProgramCounter
	StackPointer          StackPointer
	MemoryAddressBuffer   DoubleWordBuffer
	GeneralPurposeBuffer  DoubleWordBuffer
	MemoryAddressRegister MemoryAddressRegister
	MemoryDataRegister    MemoryDataRegister
	RandomAccessMemory    RandomAccessMemory
	AddressReadOnlyMemory AddressReadOnlyMemory
	ControlReadOnlyMemory ControlReadOnlyMemory
	InstructionRegister   InstructionRegister
	ArithmeticLogicUnit   ArithmeticLogicUnit
	Accumulator           Register
	TemporaryRegister     Register
	BRegister             Register
	CRegister             Register
	DRegister             Register
	ERegister             Register
	FRegister             Register
	HRegister             Register
	LRegister             Register
	InputRegister1        Register
	InputRegister2        Register
	OutputRegister3       Register
	OutputRegister4       Register
	ControllerSequencer   ControllerSequencer
	Halt                  bool
}

// Start initializes the system, and is analogous to power on.
func (s *System) Start() {
	s.InitRam()
	s.Restart()
}

// Restart is a soft init, the current state of AddressReadOnlyMemory and RandomAccessMemory are preserved.
func (s *System) Restart() {
	s.InitCounters()
	s.ClearRegisters()
	s.ClearFlags()
	s.Halt = false
}

// InitRom sets the initial state of AddressReadOnlyMemory and ControlReadOnlyMemory from the preconfigured values.
func (s *System) InitRom() {
	FillRomContents()
	s.AddressReadOnlyMemory.Values = AddressRomContents
	s.ControlReadOnlyMemory.Values = ControlRomContents
}

// LoadRom sets the initial state of AddressReadOnlyMemory and ControlReadOnlyMemory from the supplied values.
func (s *System) LoadRom(addressRomContents types.AddressRom, controlRomContents types.ControlRom) {
	s.AddressReadOnlyMemory.Values = addressRomContents
	s.ControlReadOnlyMemory.Values = controlRomContents
}

// InitRam sets the initial state of RandomAccessMemory from the preconfigured values set in RamContents.
func (s *System) InitRam() {
	s.RandomAccessMemory.Values = RamContents
}

// InitCounters sets the initial state of all the system counters.
func (s *System) InitCounters() {
	s.RingCounter.Clear()
	s.ProgramCounter.Clear()
	s.InstructionRegister.Clear()
	s.PresettableCounter.Clear()
	UpdateCtr = 0
}

// ClearRegisters empties all the registers that System writes to.
func (s *System) ClearRegisters() {
	s.MemoryAddressRegister.Address = 0
	s.MemoryDataRegister.Value = 0
	s.Accumulator.Value = 0
	s.TemporaryRegister.Value = 0
	s.BRegister.Value = 0
	s.CRegister.Value = 0
	s.Bus.Value = 0
	s.InputRegister1.Value = 0
	s.InputRegister2.Value = 0
	s.OutputRegister3.Value = 0
	s.OutputRegister4.Value = 0
}

// ClearFlags sets all flags to Low and empties the flags register.
func (s *System) ClearFlags() {
	s.ArithmeticLogicUnit.Flags.Carry = types.Low
	s.ArithmeticLogicUnit.Flags.Zero = types.Low
	s.ArithmeticLogicUnit.Flags.Sign = types.Low
	s.ArithmeticLogicUnit.Flags.Parity = types.Low
	s.FRegister.Value = 0
}

// Reset resets various System clocks.
func (s *System) Reset() {
	s.RingCounter.Clear()
	s.PresettableCounter.Clear()
}

// Update handles the various system processes.
func (s *System) Update() {
	s.ControllerSequencer.Update(s)
}
