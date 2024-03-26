package unitconvert

var MeterPerSecond Unit = Unit{Name: "meter per second", Type: VelocityUnit, System: "metric"}
var FeetPerSecond Unit = Unit{Name: "feet per second", Type: VelocityUnit, System: "english"}
var KilometerPerHour Unit = Unit{Name: "kilometer per hour", Type: VelocityUnit, System: "metric"}
var MilesPerHour Unit = Unit{Name: "miles per hour", Type: VelocityUnit, System: "english"}

// Velocity calculates the conversion of a value from one velocity unit to another.
// It takes a value (v) in the source unit (from) and converts it to the destination unit (to).
// Parameters:
//   - v: The value to be converted.
//   - from: The unit of the input value.
//   - to: The unit to which the input value is to be converted.
//
// Returns:
//   - float64: The converted value.
//   - error: An error indicating any conversion or input validation issues.
//
// Example:
//
//	v := 2.5 // Value to be converted
//	origUnit := unitconvert.MeterPerSecond // Original unit
//	destUnit := unitconvert.FeetPerSecond  // Destination unit
//	newValue, err := unitconvert.Length(v, origUnit, destUnit) // Conversion
func Velocity(v float64, from, to Unit) (float64, error) {
	if from.Type != VelocityUnit || to.Type != VelocityUnit {
		return 0, invalidInputTypes
	}

	if from == to {
		return v, nil
	}

	factor, ok := conversionFactors[from][to]
	if !ok {
		return 0, unsupportedDestinationError
	}

	return v * factor, nil
}
