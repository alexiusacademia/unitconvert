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

// String representation of a unit.
//
// Returns:
// - The string representation of the unit (string).
func (u UnitLength) String() string {
	return [...]string{
		"meter",
		"feet",
	}[u]
}
