package utils

// Location is a struct which stores a particular  This can be a
// register, an address (on the heap) or a position on the stack.
type Location struct {
	Register Register
	Address  int

	// Stores information needed to determine stack offset
	CurrentPos int
}

// NewStackOffsetLocation builds a location using a stack position
func NewStackOffsetLocation(currentPos int) *Location {
	return &Location{
		Register:   UNDEFINED,
		CurrentPos: currentPos,
	}
}
