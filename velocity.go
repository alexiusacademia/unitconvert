package unitconvert

var MeterPerSecond Unit = Unit{Name: "meter per second", Type: VelocityUnit, System: "metric"}
var FeetPerSecond Unit = Unit{Name: "feet per second", Type: VelocityUnit, System: "english"}
var KilometerPerHour Unit = Unit{Name: "kilometer per hour", Type: VelocityUnit, System: "metric"}
var MilesPerHour Unit = Unit{Name: "miles per hour", Type: VelocityUnit, System: "english"}

// Convert velocity quantities.
//
// Args:
// - v (float64): The value to be converted.\n
// - from (Unit): The unit of the value.\n
// - to (Unit): The destination unit.\n
//
// Returns:
// - The converted value (float64).
// - An error if there is one (error).
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
