package location

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
