package object

import (
	"fmt"
	"unsafe"
)

// Object is the root of the object hierarchy, similar to java.lang.Object.
type Object interface {
	// String returns a string representation of the object.
	String() string
	// Equals returns true if the other object is equal to this one.
	Equals(other Object) bool
	// HashCode returns a hash code value for the object.
	HashCode() int
}

// BaseObject provides a default implementation of the Object interface.
// Structs can embed this to "inherit" the default behavior.
type BaseObject struct{}

// String returns a default string representation, which is the memory address of the object.
// It is recommended to override this in embedding structs.
func (o *BaseObject) String() string {
	return fmt.Sprintf("%T@%p", o, o)
}

// Equals provides a default equality check based on reference equality.
// It is recommended to override this for value-based equality.
func (o *BaseObject) Equals(other Object) bool {
	return o == other
}

// HashCode provides a default hash code based on the object's pointer address.
// It is recommended to override this.
func (o *BaseObject) HashCode() int {
	return int(uintptr(unsafe.Pointer(o)))
}
