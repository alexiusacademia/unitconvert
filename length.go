package unitconvert

type UnitLength int

var Meter Unit = Unit{Name: "meter", Type: LengthUnit, System: "metric"}
var Feet Unit = Unit{Name: "feet", Type: LengthUnit, System: "metric"}

const meterToFeet float64 = 3.28084

var conversionFactors = map[Unit]map[Unit]float64{
	Meter: {
		Feet: meterToFeet,
		// Add more here
	},
	Feet: {
		Meter: 1 / meterToFeet,
		// Add more here
	},
	// Add more here
}

// Convert length quantities.
//
// Args:
// - v (float64): The value to be converted.
// - from (UnitLength): The unit of the value.
// - to (UnitLength): The destination unit.
//
// Returns:
// - The converted value (float64).
// - An error if there is one (error).
func Length(v float64, from, to Unit) (float64, error) {

	if from == to {
		return v, nil
	}

	factor, ok := conversionFactors[from][to]
	if !ok {
		return 0, UnsupportedDestinationError
	}

	return v * factor, nil
}
