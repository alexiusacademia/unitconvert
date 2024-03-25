package unitconvert

type UnitType int

// Constants representing different types of measurement units.
const (
	LengthUnit UnitType = iota
	AreaUnit
	VolumeUnit
	VelocityUnit
	DischargeUnit
)
