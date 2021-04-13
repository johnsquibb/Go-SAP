package machine

import (
	"encoding/json"
	"log"
)

var debugEnabled = false

// EnableDebug enables the Debug output.
func EnableDebug(enabled bool) {
	debugEnabled = enabled
}

// Debug outputs JSON serialized System state excluding large objects such as RAM, ROM.
func (s System) Debug() {
	o := GetSystemDebugObject(s)
	o.RandomAccessMemory = nil
	o.ControlReadOnlyMemory = nil
	o.AddressReadOnlyMemory = nil
	b, _ := json.Marshal(o)
	Debug(string(b))
}

// DebugFull outputs complete JSON serialized System state.
func (s System) DebugFull() {
	o := GetSystemDebugObject(s)
	b, _ := json.Marshal(o)
	Debug(string(b))
}

// DumpFormattedSystemDebugObject formatted JSON serialized System state excluding large objects.
func (s System) DumpFormattedSystemDebugObject() {
	b := s.GetFormattedSystemDebugObject()
	Debug(b)
}

// GetFormattedSystemDebugObject returns formatted JSON serialized System state excluding large objects.
func (s System) GetFormattedSystemDebugObject() string {
	o := GetSystemDebugObject(s)
	o.RandomAccessMemory = nil
	o.ControlReadOnlyMemory = nil
	o.AddressReadOnlyMemory = nil
	b, _ := json.MarshalIndent(o, "  ", "  ")
	return string(b)
}

// DumpFormattedFullSystemDebugObject dumps formatted complete JSON serialized System state.
func (s System) DumpFormattedFullSystemDebugObject() {
	o := GetSystemDebugObject(s)
	b, _ := json.MarshalIndent(o, "  ", "  ")
	Debug(string(b))
}

// Debug writes messages to standard out.
func Debug(message ...interface{}) {
	if debugEnabled {
		log.Println(message...)
	}
}

// DebugJson supplies formatted JSON messages to Debug.
func DebugJson(item interface{}) {
	if debugEnabled {
		b, _ := json.MarshalIndent(item, "  ", "  ")
		Debug(string(b))
	}
}

type SystemDebug struct {
	Bus                   Bus
	RingCounter           RingCounter
	PresettableCounter    PresettableCounter
	ProgramCounter        ProgramCounter
	StackPointer          StackPointer
	MemoryAddressBuffer   DoubleWordBuffer
	GeneralPurposeBuffer  DoubleWordBuffer
	MemoryAddressRegister MemoryAddressRegister
	MemoryDataRegister    MemoryDataRegister
	RandomAccessMemory    *RandomAccessMemory
	AddressReadOnlyMemory *AddressReadOnlyMemory
	ControlReadOnlyMemory *ControlReadOnlyMemory
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

func GetSystemDebugObject(s System) SystemDebug {
	return SystemDebug{
		Bus:                   s.Bus,
		RingCounter:           s.RingCounter,
		PresettableCounter:    s.PresettableCounter,
		ProgramCounter:        s.ProgramCounter,
		StackPointer:          s.StackPointer,
		MemoryAddressBuffer:   s.MemoryAddressBuffer,
		GeneralPurposeBuffer:  s.GeneralPurposeBuffer,
		MemoryAddressRegister: s.MemoryAddressRegister,
		MemoryDataRegister:    s.MemoryDataRegister,
		RandomAccessMemory:    &s.RandomAccessMemory,
		AddressReadOnlyMemory: &s.AddressReadOnlyMemory,
		ControlReadOnlyMemory: &s.ControlReadOnlyMemory,
		InstructionRegister:   s.InstructionRegister,
		ArithmeticLogicUnit:   s.ArithmeticLogicUnit,
		Accumulator:           s.Accumulator,
		TemporaryRegister:     s.TemporaryRegister,
		BRegister:             s.BRegister,
		CRegister:             s.CRegister,
		DRegister:             s.DRegister,
		ERegister:             s.ERegister,
		FRegister:             s.FRegister,
		HRegister:             s.HRegister,
		LRegister:             s.LRegister,
		InputRegister1:        s.InputRegister1,
		InputRegister2:        s.InputRegister2,
		OutputRegister3:       s.OutputRegister3,
		OutputRegister4:       s.OutputRegister4,
		ControllerSequencer:   s.ControllerSequencer,
		Halt:                  s.Halt,
	}
}
