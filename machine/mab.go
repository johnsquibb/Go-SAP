package machine

import (
	"Go-SAP/machine/types"
)

// DoubleWordBuffer is a 16-bit register for combining 2 8-bit values (LSB, MSB) into a 16-bit memory address.
// Values should be read serially over multiple T cycles while enabling LSBReadEnable or MSBReadEnable, and always from
// the LSB (8-bits) of the Bus.
// Values can be written simultaneously or individually to the Bus using LSBWriteEnable and/or MSBWriteEnable.
type DoubleWordBuffer struct {
	LSBValue       types.Word
	MSBValue       types.Word
	LSBReadEnable  types.State
	LSBWriteEnable types.State
	MSBReadEnable  types.State
	MSBWriteEnable types.State
}

// Update reads 8-bit LSBValue from Bus when LSBReadEnable is High.
// Update reads 8-bit MSBValue from Bus when MSBReadEnable is High.
// Update reads 16-bit MSBValue and LSBValue when MSBReadEnable and LSBReadEnable are High.
// Update writes (Bus | LSBValue) to Bus when LSBWriteEnable is High.
// Update writes (Bus | MSBValue) to Bus when MSBWriteEnable is High.
// Update writes MSBValue to MSB of Bus and LSBValue to LSB of Bus when LSBWriteEnable and MSBWriteEnable are High.
func (r *DoubleWordBuffer) Update(bus *Bus) {

	// Read full 16-bit bus value as MSB, LSB.
	if r.LSBReadEnable && r.MSBReadEnable {
		r.LSBValue = types.Word(bus.Value)
		r.MSBValue = types.Word(bus.Value >> 8)
		return
	}

	// Read current 8-bit bus value as LSB.
	if r.LSBReadEnable == types.High {
		r.LSBValue = types.Word(bus.Value)
	}

	// Read current 8-bit bus value as MSB.
	if r.MSBReadEnable == types.High {
		r.MSBValue = types.Word(bus.Value)
	}

	// Write MSB and LSB, overwrite bus.
	if r.MSBWriteEnable == types.High && r.LSBWriteEnable == types.High {
		v := types.DoubleWord(r.MSBValue)<<8 | types.DoubleWord(r.LSBValue)
		bus.Value = v
		return
	}

	// Write only LSB to bus.
	if r.LSBWriteEnable == types.High {
		bus.Value = types.DoubleWord(r.LSBValue)
	}

	// Write only MSB to bus.
	if r.MSBWriteEnable == types.High {
		bus.Value = types.DoubleWord(r.MSBValue)
	}
}
