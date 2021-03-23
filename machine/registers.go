package machine

import (
	"Go-SAP3/machine/types"
)

// Register is a general-purpose register.
type Register struct {
	Value       types.Word
	ReadEnable  types.State
	WriteEnable types.State
}

// Update reads Value of Register from Bus when ReadEnable is High.
// Update writes Value of Register to Bus when WriteEnable is High.
func (r *Register) Update(bus *Bus) {
	if r.ReadEnable == types.High {
		r.Value = types.Word(bus.Value)
	}

	if r.WriteEnable == types.High {
		bus.Value = types.DoubleWord(r.Value)
	}
}

// AluRegister describes contract for any register that can act as ACC or TEMP in ArithmeticLogicUnit.
type AluRegister interface {
	Get() types.Word
}

// Get implements AluRegister contract for a Register.
func (r Register) Get() types.Word {
	return r.Value
}