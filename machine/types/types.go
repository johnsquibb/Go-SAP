package types

// Word is the SAP-2 word size of 8 bits, used across all Register(s).
type Word uint8

// DoubleWord is a double word used for Bus, ProgramCounter, and RandomAccessMemory addresses.
type DoubleWord uint16

// QuadrupleWord is a quadruple word of 32 bits.
type QuadrupleWord uint32

// OctupleWord is a octuple word of 64 bits.
type OctupleWord uint64

// State is High or Low.
type State bool

// High is on state.
const High State = true

// Low is off state.
const Low State = false

// ramSize is the total size of Ram.
// The first 2K of Ram in SAP-2 is ROM (0x0000 - 0x07FF).
// Usable address space begins at 0x0800.
const ramSize = 65536

// Ram is allocated in Word(s).
type Ram [ramSize]Word

// addressRomSize is the total size of AddressRom.
const addressRomSize = 256

// AddressRom stores Word(s).
type AddressRom [addressRomSize]DoubleWord

// controlRomSize is the total size of the ControlRom.
const controlRomSize = 1024

// ControlRom stores OctupleWord(s).
type ControlRom [controlRomSize]OctupleWord
