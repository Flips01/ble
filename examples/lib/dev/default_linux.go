package dev

import (
	"github.com/Flips01/ble"
	"github.com/Flips01/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
