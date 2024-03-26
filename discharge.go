package unitconvert

var CubicMeterPerSecond Unit = Unit{Name: "cubic meter per second", Type: DischargeUnit, System: "metric"}
var CubicFeetPerSecond Unit = Unit{Name: "cubic feet per second", Type: DischargeUnit, System: "english"}

// Convert Discharge quantities.
//
// Args:
// - v (float64): The value to be converted.
// - from (Unit): The unit of the value.
// - to (Unit): The destination unit.
//
// Returns:
// - The converted value (float64).
// - An error if there is one (error).
func Discharge(v float64, from, to Unit) (float64, error) {
	if from.Type != DischargeUnit || to.Type != DischargeUnit {
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
