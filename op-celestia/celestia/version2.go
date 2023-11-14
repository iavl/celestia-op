package celestia

import (
	"encoding/binary"
)

// FrameCelestiaStdRef contains the reference to the specific frame on celestia and
// satisfies the Framer interface.
type FrameCelestiaStdRef struct {
	BlockHeight  uint64
	TxCommitment []byte
}

var _ Framer = &FrameCelestiaStdRef{}

// MarshalBinary encodes the FrameCelestiaStdRef to binary
// serialization format: height + commitment
//
//	----------------------------------------
//
// | 8 byte uint64  |  32 byte commitment   |
//
//	----------------------------------------
//
// | <-- height --> | <-- commitment -->    |
//
//	----------------------------------------
func (f *FrameCelestiaStdRef) MarshalBinary() ([]byte, error) {
	ref := make([]byte, 8+len(f.TxCommitment))

	binary.LittleEndian.PutUint64(ref, f.BlockHeight)
	copy(ref[8:], f.TxCommitment)

	return append([]byte{FrameCelestiaStd}, ref...), nil
}

// UnmarshalBinary decodes the binary to FrameCelestiaStdRef
// serialization format: height + commitment
//
//	----------------------------------------
//
// | 8 byte uint64  |  32 byte commitment   |
//
//	----------------------------------------
//
// | <-- height --> | <-- commitment -->    |
//
//	----------------------------------------
func (f *FrameCelestiaStdRef) UnmarshalBinary(ref []byte) error {
	if len(ref) <= 8 {
		return ErrInvalidSize
	}
	f.BlockHeight = binary.LittleEndian.Uint64(ref[:8])
	f.TxCommitment = ref[8:]
	return nil
}
