package unitconvert

import "errors"

var defaultError error = errors.New("Unsupported unit destination.")

type UnitLength int

const (
	Meter UnitLength = iota
	Feet
)

const meterToFeet float64 = 3.28084

var conversionFactors = map[UnitLength]map[UnitLength]float64{
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
func Length(v float64, from, to UnitLength) (float64, error) {

	if from == to {
		return v, nil
	}

	factor, ok := conversionFactors[from][to]
	if !ok {
		return 0, defaultError
	}

	return v * factor, nil
}
