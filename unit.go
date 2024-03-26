package unitconvert

// Unit represents a unit of measurement with its name, type, and system.
// Parameters:
//   - Name: The name of the unit.
//   - Type: The type of the unit (e.g., Meter, SquareMeter, MilesPerHour).
//   - System: The system to which the unit belongs (e.g., metric, english).
type Unit struct {
	Name   string
	Type   UnitType
	System string
}
