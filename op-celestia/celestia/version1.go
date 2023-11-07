package celestia

// FrameEthereumStdRef contains the reference to the specific frame on celestia and
// satisfies the Framer interface.
type FrameEthereumStdRef struct {
	Calldata []byte
}

var _ Framer = &FrameEthereumStdRef{}

// MarshalBinary encodes the FrameEthereumStdRef to binary
// serialization format: calldata
func (f *FrameEthereumStdRef) MarshalBinary() ([]byte, error) {
	return append([]byte{FrameEthereumStd}, f.Calldata...), nil
}

// UnmarshalBinary decodes the binary to FrameEthereumStdRef
// serialization format: calldata
func (f *FrameEthereumStdRef) UnmarshalBinary(ref []byte) error {
	f.Calldata = ref
	return nil
}
