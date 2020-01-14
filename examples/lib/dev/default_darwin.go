package dev

import (
	"github.com/Flips01/ble"
	"github.com/Flips01/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
