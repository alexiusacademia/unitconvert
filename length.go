package unitconvert

import "errors"

var defaultError error = errors.New("Unsupported unit destination.")

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
func Length(v float64, from UnitLength, to UnitLength) (float64, error) {
	const meterToFeet float64 = 3.2808

	if from == to {
		return v, nil
	}

	if from == Feet {
		if to == Meter {
			return v / meterToFeet, nil
		} else {
			return 0, defaultError
		}
	}

	if from == Meter {
		if to == Feet {
			return v * meterToFeet, nil
		} else {
			return 0, defaultError
		}
	}

	return 0, nil
}
