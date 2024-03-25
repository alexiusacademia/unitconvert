package unitconvert

var SquareMeter Unit = Unit{Name: "square meter", Type: AreaUnit, System: "metric"}
var SquareFeet Unit = Unit{Name: "square feet", Type: AreaUnit, System: "english"}

// Area calculates the conversion of a value from one area unit to another.
// It takes a value (v) in the source unit (from) and converts it to the destination unit (to).
// Parameters:
//   - v: The value to be converted.
//   - from: The unit of the input value.
//   - to: The unit to which the input value is to be converted.
//
// Returns:
//   - float64: The converted value.
//   - error: An error indicating any conversion or input validation issues.
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
