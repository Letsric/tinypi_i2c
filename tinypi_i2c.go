package tinypii2c

import (
	"github.com/ftl/i2c"
)

type pi_I2C struct {
	busid int
}

func New(bus int) *pi_I2C {
	return &pi_I2C{
		busid: bus,
	}
}

func (p *pi_I2C) Tx(addr uint16, w, r []byte) error {
	bus, error := i2c.Open(uint8(addr), p.busid)

	if error != nil {
		return error
	}

	if w != nil {
		// Write
		bus.Write(w)
	}

	if r != nil {
		// Read
		bus.Read(r)
	}

	return nil
}
