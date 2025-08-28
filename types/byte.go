package types

// Byte is an alias for UInt8, as byte is an alias for uint8 in Go.
type Byte = UInt8

// NewByte creates a new Byte (UInt8) object.
func NewByte(value byte) *Byte {
	return NewUInt8(value)
}
