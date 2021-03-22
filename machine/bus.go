package machine

import "SystemOnePoc/sap3/machine/types"

// Bus is a transport for moving Value around the system.
type Bus struct {
	Value types.DoubleWord
}

// Set puts the Value on Bus.
func (b *Bus) Set(v types.DoubleWord) {
	b.Value = v
}

// Get gets the Value on Bus.
func (b Bus) Get() types.DoubleWord {
	return b.Value
}
