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

type UnitType int

// Constants representing different types of measurement units.
const (
	LengthUnit UnitType = iota
	AreaUnit
	VolumeUnit
	VelocityUnit
	DischargeUnit
)

// Length units
var Meter Unit = Unit{Name: "meter", Type: LengthUnit, System: "metric"}
var Feet Unit = Unit{Name: "feet", Type: LengthUnit, System: "english"}

// Area Units
var SquareMeter Unit = Unit{Name: "square meter", Type: AreaUnit, System: "metric"}
var SquareFeet Unit = Unit{Name: "square feet", Type: AreaUnit, System: "english"}

// Velocity Units
var MeterPerSecond Unit = Unit{Name: "meter per second", Type: VelocityUnit, System: "metric"}
var FeetPerSecond Unit = Unit{Name: "feet per second", Type: VelocityUnit, System: "english"}
var KilometerPerHour Unit = Unit{Name: "kilometer per hour", Type: VelocityUnit, System: "metric"}
var MilesPerHour Unit = Unit{Name: "miles per hour", Type: VelocityUnit, System: "english"}

// Discharge Units
var CubicMeterPerSecond Unit = Unit{Name: "cubic meter per second", Type: DischargeUnit, System: "metric"}
var CubicFeetPerSecond Unit = Unit{Name: "cubic feet per second", Type: DischargeUnit, System: "english"}
var LiterPerSecond Unit = Unit{Name: "liter per second", Type: DischargeUnit, System: "metric"}
