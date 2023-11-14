package celestia

import (
	"encoding"
	"errors"
)

var (
	ErrInvalidSize    = errors.New("invalid size")
	ErrInvalidVersion = errors.New("invalid version")
)

// Framer defines a way to encode/decode a FrameRef.
type Framer interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

const (
	FrameCelestiaLegacy = iota
	FrameEthereumStd
	FrameCelestiaStd
)
