package location

// Location is a struct which stores a particular location. This can be a
// register, an address (on the heap) or a position on the stack.
type Location struct {
	Register Register
	Address  int

	// Stores information needed to determine stack offset
	CurrentPos int
}

func NewRegisterLocation(register Register) *Location {
	return &Location{
		Register: register,
	}
}

func NewAddressLocation(address int) *Location {
	return &Location{
		Register: UNDEFINED,
		Address:  address,
	}
}

func NewStackOffsetLocation(currentPos int) *Location {
	return &Location{
		Register:   UNDEFINED,
		CurrentPos: currentPos,
	}
}
