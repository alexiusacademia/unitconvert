package unitconvert

type UnitLength int
type UnitVolume int

const (
	Meter UnitLength = iota
	Feet
)

const (
	CubicMeter UnitVolume = iota
	CubicFeet
)

// String representation of a length unit (UnitLength).
//
// Returns:
// - The string representation of the unit (string).
func (u UnitLength) String() string {
	return [...]string{
		"meter",
		"feet",
	}[u]
}

// String representation of a volume unit (UnitVolume).
//
// Returns:
// - The string representation of the unit (string).
func (u UnitVolume) String() string {
	return [...]string{
		"cubic meter",
		"cubic feet",
	}[u]
}
