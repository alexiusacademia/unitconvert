package unitconvert

type UnitArea int
type UnitVolume int

const (
	SquareMeter UnitArea = iota
	SquareFoot
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
