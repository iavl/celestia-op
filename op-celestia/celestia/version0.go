package celestia

import (
	"encoding/binary"
)

// FrameCelestiaLegacyRef contains the reference to the specific frame on celestia and
// satisfies the Framer interface.
type FrameCelestiaLegacyRef struct {
	BlockHeight uint64
	TxIndex     uint32
}

var _ Framer = &FrameCelestiaLegacyRef{}

// MarshalBinary encodes the FrameCelestiaLegacyRef to binary
// serialization format: height + index
//
//	----------------------------------------
//
// | 8 byte uint64  |  4 byte index   |
//
//	----------------------------------------
//
// | <-- height --> | <-- tx index -->|
//
//	----------------------------------------
func (f *FrameCelestiaLegacyRef) MarshalBinary() ([]byte, error) {
	ref := make([]byte, 8+4)

	binary.LittleEndian.PutUint64(ref, f.BlockHeight)
	binary.LittleEndian.PutUint32(ref[8:], f.TxIndex)

	return append([]byte{FrameCelestiaLegacy}, ref...), nil
}

// UnmarshalBinary decodes the binary to FrameCelestiaLegacyRef
// serialization format: height + index
//
//	----------------------------------------
//
// | 8 byte uint64  |  4 byte index   |
//
//	----------------------------------------
//
// | <-- height --> | <-- tx index -->|
//
//	----------------------------------------
func (f *FrameCelestiaLegacyRef) UnmarshalBinary(ref []byte) error {
	if len(ref) != 12 {
		return ErrInvalidSize
	}
	f.BlockHeight = binary.LittleEndian.Uint64(ref[:8])
	f.TxIndex = binary.LittleEndian.Uint32(ref[8:])
	return nil
}
