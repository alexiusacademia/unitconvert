package unitconvert

import "errors"

var defaultError error = errors.New("Unsupported unit destination.")

func Length(v float64, from UnitLength, to UnitLength) (float64, error) {
	if from == to {
		return v, nil
	}

	if from == Feet {
		if to == Meter {
			return v / 3.28, nil
		} else {
			return 0, defaultError
		}
	}

	if from == Meter {
		if to == Feet {
			return v * 3.28, nil
		} else {
			return 0, defaultError
		}
	}

	return 0, nil
}
