package unitconvert

var SquareMeter Unit = Unit{Name: "square meter", Type: AreaUnit, System: "metric"}
var SquareFeet Unit = Unit{Name: "square feet", Type: AreaUnit, System: "metric"}

// Convert length quantities.
//
// Args:
// - v (float64): The value to be converted.
// - from (Unit): The unit of the value.
// - to (Unit): The destination unit.
//
// Returns:
// - The converted value (float64).
// - An error if there is one (error).
func Area(v float64, from, to Unit) (float64, error) {
	if from.Type != AreaUnit || to.Type != AreaUnit {
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
