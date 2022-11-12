package machine

import (
	"Go-SAP3/machine/types"
	"math/bits"
)

// ArithmeticLogicUnit is asynchronous (un-clocked) ; Its value will change as soon as input words change.
type ArithmeticLogicUnit struct {
	A           types.Word
	B           types.Word
	Sum         types.Word
	WriteEnable types.State
	AluMode     AluMode
	Flags       Flags
}

type Flags struct {
	Carry        types.State
	Sign         types.State
	Zero         types.State
	Parity       types.State
	CarryEnable  types.State
	SignEnable   types.State
	ZeroEnable   types.State
	ParityEnable types.State
}

// AluMode bits control the operating mode of the ALU.
type AluMode struct {
	Bit0 types.State // LSB
	Bit1 types.State
	Bit2 types.State
	Bit3 types.State
	Bit4 types.State // MSB
}

// AluMode modes are integer value of AluMode Bit flags.
const (
	AluAdd              = 0b00000
	AluSubtract         = 0b00001
	AluIncrement        = 0b00010
	AluDecrement        = 0b00011
	AluComplement       = 0b00100
	AluAnd              = 0b00101
	AluOr               = 0b00110
	AluExclusiveOr      = 0b00111
	AluRotateLeft       = 0b01000
	AluRotateRight      = 0b01001
	AluAddCarry         = 0b01010
	AluSubtractBorrow   = 0b01011
	AluSetCarry         = 0b01100
	AluComplementCarry  = 0b01101
	AluRotateLeftCarry  = 0b01110
	AluRotateRightCarry = 0b01111
	AluCompare          = 0b10000
	AluIncrementPair    = 0b10001
	AluDecrementPair    = 0b10010
)

// Calculate returns the result of the arithmetic operation based on AluMode mode.
func (a *ArithmeticLogicUnit) Calculate() {

	adder := BinaryAdder{}
	var carryOut types.State

	switch a.AluMode.getMode() {
	case AluAdd:
		a.Sum, carryOut = adder.Calculate(uint8(a.A), uint8(a.B), types.Low)
	case AluAddCarry:
		a.Sum, carryOut = adder.Calculate(uint8(a.A), uint8(a.B), a.Flags.Carry)
	case AluSubtract:
		bComplement2 := (^uint8(a.B)) + 1
		a.Sum, carryOut = adder.Calculate(uint8(a.A), bComplement2, types.Low)
		// 8085 complements the carry flag on subtraction.
		carryOut = !carryOut
	case AluSubtractBorrow:
		// Combine B and Carry, subtract from A.
		b := uint8(a.B)
		if a.Flags.Carry {
			b += 1
		}
		bComplement2 := (^b) + 1
		a.Sum, carryOut = adder.Calculate(uint8(a.A), bComplement2, types.Low)

		// 8085 complements the carry flag on subtraction.
		carryOut = !carryOut
	case AluIncrement:
		a.Sum, carryOut = adder.Calculate(uint8(a.A), 1, types.Low)
	case AluDecrement:
		complement2 := (^uint8(1)) + 1
		a.Sum, carryOut = adder.Calculate(uint8(a.A), complement2, types.Low)
		// 8085 complements the carry flag on subtraction.
		carryOut = !carryOut
	case AluIncrementPair:
		// 8085 INX does not affect any flags by design (since typically used for memory addresses).
		var c types.State
		a.B, c = adder.Calculate(uint8(a.B), 1, types.Low)
		if c {
			a.A, _ = adder.Calculate(uint8(a.A), 1, types.Low)
		}
	case AluDecrementPair:
		// 8085 INX does not affect any flags by design (since typically used for memory addresses).
		var c types.State
		complement2 := (^uint8(1)) + 1
		a.B, c = adder.Calculate(uint8(a.B), complement2, types.Low)
		// 8085 complements the carry flag on subtraction.
		c = !c
		if c {
			a.A, _ = adder.Calculate(uint8(a.A), complement2, types.Low)
		}
	case AluComplement:
		a.Sum = ^a.A
	case AluAnd:
		a.Sum = a.A & a.B
	case AluOr:
		a.Sum = a.A | a.B
	case AluExclusiveOr:
		a.Sum = a.A ^ a.B
	case AluSetCarry:
		a.Flags.Carry = types.High
	case AluComplementCarry:
		a.Flags.Carry = !a.Flags.Carry
	case AluRotateLeft:
		carryOut = a.rotateAllLeft()
	case AluRotateRight:
		carryOut = a.rotateAllRight()
	case AluRotateLeftCarry:
		carryOut = a.rotateLeftCarry()
	case AluRotateRightCarry:
		carryOut = a.rotateRightCarry()
	case AluCompare:
		// Compare does B-A.
		aComplement2 := (^uint8(a.A)) + 1

		a.Sum, carryOut = adder.Calculate(uint8(a.B), aComplement2, types.Low)

		// 8085 complements the carry flag on subtraction.
		carryOut = !carryOut
	}

	if a.Flags.SignEnable == types.High {
		a.Flags.Sign = a.Sum&0x80 != 0
	}

	if a.Flags.ZeroEnable == types.High {
		a.Flags.Zero = a.Sum == 0
	}

	if a.Flags.CarryEnable == types.High {
		a.Flags.Carry = carryOut
	}

	if a.Flags.ParityEnable == types.High {
		p := bits.OnesCount8(uint8(a.Sum))
		a.Flags.Parity = p%2 == 0
	}
}

// getMode returns the Word value determined from Bit0 thru Bit3 flags.
func (a AluMode) getMode() types.Word {
	v := 0

	if a.Bit4 == types.High {
		v |= 0b10000
	}

	if a.Bit3 == types.High {
		v |= 0b01000
	}

	if a.Bit2 == types.High {
		v |= 0b00100
	}

	if a.Bit1 == types.High {
		v |= 0b00010
	}

	if a.Bit0 == types.High {
		v |= 0b00001
	}

	return types.Word(v)
}

// Update refreshes the state of ArithmeticLogicUnit, summing the values currently set to its registers.
// When ArithmeticLogicUnit WriteEnable is High, the value is moved immediately to Bus Value.
func (a *ArithmeticLogicUnit) Update(acc AluRegister, temp AluRegister, s *System) {

	a.A = acc.Get()
	a.B = temp.Get()

	if a.WriteEnable == types.High {
		a.GetFlags(s.FRegister)
		a.Calculate()
		a.SetFlags(&s.FRegister)
		s.Bus.Value = types.DoubleWord(a.Sum)
	}
}

// GetFlags reads flag bits from Register and applies to ArithmeticLogicUnit.
// Register bits: [S Z 0 0 0 P 0 C]
func (a *ArithmeticLogicUnit) GetFlags(f Register) {
	a.Flags.Carry = f.Value&0b0000_0001 != 0
	a.Flags.Parity = f.Value&0b0000_0100 != 0
	a.Flags.Zero = f.Value&0b0100_0000 != 0
	a.Flags.Sign = f.Value&0b1000_0000 != 0
}

// SetFlags writes flag bits to Register from settings in ArithmeticLogicUnit.
// Register bits: [S Z 0 0 0 P 0 C]
func (a *ArithmeticLogicUnit) SetFlags(f *Register) {
	if a.Flags.Sign {
		f.Value |= 0b1000_0000
	} else {
		f.Value &= 0b0111_1111
	}

	if a.Flags.Zero {
		f.Value |= 0b0100_0000
	} else {
		f.Value &= 0b1011_1111
	}

	if a.Flags.Parity {
		f.Value |= 0b0000_0100
	} else {
		f.Value &= 0b1111_1011
	}

	if a.Flags.Carry {
		f.Value |= 0b0000_0001
	} else {
		f.Value &= 0b1111_1110
	}
}

func (a *ArithmeticLogicUnit) rotateAllLeft() types.State {
	carryOut := a.A&0b1000_0000 != 0

	a.Sum = a.A << 1
	if a.Flags.Carry {
		a.Sum += 1
	}

	return types.State(carryOut)
}

func (a *ArithmeticLogicUnit) rotateAllRight() types.State {
	carryOut := a.A&0b0000_0001 != 0

	a.Sum = a.A >> 1
	if a.Flags.Carry {
		a.Sum |= 0b1000_0000
	}

	return types.State(carryOut)
}

func (a *ArithmeticLogicUnit) rotateLeftCarry() types.State {
	carryOut := a.A&0b1000_0000 != 0

	a.Sum = a.A << 1

	return types.State(carryOut)
}

func (a *ArithmeticLogicUnit) rotateRightCarry() types.State {
	carryOut := a.A&0b0000_0001 != 0

	a.Sum = a.A >> 1

	return types.State(carryOut)
}
