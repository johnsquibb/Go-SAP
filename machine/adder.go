package machine

import (
	"Go-SAP3/machine/types"
)

type BinaryAdder struct {
}

func (ba *BinaryAdder) Calculate(a, b uint8, carryIn types.State) (sum types.Word, carryOut types.State) {
	var carry = types.DoubleWord(0)
	if carryIn {
		carry = 1
	}

	result := types.DoubleWord(a) + types.DoubleWord(b) + carry
	carryOut = result&0b1_0000_0000 != 0
	sum = types.Word(result)
	
	return
}
