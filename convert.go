package unitconvert

// Convert quantities.
//
// Args:
// - v (float64): The value to be converted.
// - from (Unit): The unit of the value.
// - to (Unit): The destination unit.
//
// Returns:
// - The converted value (float64).
// - An error if there is one (error).
func Convert(v float64, from, to Unit) (float64, error) {
	if from.Type != to.Type {
		return 0, errIncompatibleUnitTypes
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

//

func Length(v float64, from, to Unit) (float64, error) {
	if from.Type != LengthUnit || to.Type != LengthUnit {
		return 0, invalidInputTypes
	}

	return Convert(v, from, to)
}

func Area(v float64, from, to Unit) (float64, error) {
	if from.Type != AreaUnit || to.Type != AreaUnit {
		return 0, invalidInputTypes
	}

	return Convert(v, from, to)
}

func Velocity(v float64, from, to Unit) (float64, error) {
	if from.Type != VelocityUnit || to.Type != VelocityUnit {
		return 0, invalidInputTypes
	}

	return Convert(v, from, to)
}

func Discharge(v float64, from, to Unit) (float64, error) {
	if from.Type != DischargeUnit || to.Type != DischargeUnit {
		return 0, invalidInputTypes
	}

	return Convert(v, from, to)
}
