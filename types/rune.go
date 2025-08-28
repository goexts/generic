package types

// Rune is an alias for Int32, as rune is an alias for int32 in Go.
type Rune = Int32

// NewRune creates a new Rune (Int32) object.
func NewRune(value rune) *Rune {
	return NewInt32(value)
}
