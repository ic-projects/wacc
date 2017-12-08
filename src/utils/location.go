package utils

// Location is a struct which stores a particular  This can be a
// register, an address (on the heap) or a position on the stack.
type Location struct {
	Register Register

	// If true, the stack position is used to store the address on heap
	IsOnHeap bool

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

// NewHeapLocation builds a location using the stack position which stores the
// address on the heap.
func NewHeapLocation(currentPos int) *Location {
	return &Location{
		IsOnHeap:   true,
		CurrentPos: currentPos,
	}
}

// IsRegister will return true if the location is a register
func (loc *Location) IsRegister() bool {
	return loc.Register != UNDEFINED
}

// IsAddress will return true if the location is an address on the heap
func (loc *Location) IsAddress() bool {
	return loc.IsOnHeap
}

// IsStackOffset will return true if the location a stack offset
func (loc *Location) IsStackOffset() bool {
	return !loc.IsRegister() && !loc.IsAddress()
}
