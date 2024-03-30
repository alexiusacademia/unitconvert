package unitconvert

import "errors"

var unsupportedDestinationError error = errors.New("Unsupported unit destination.")
var invalidInputTypes error = errors.New("Invalid unit types given for the conversion.")
var errIncompatibleUnitTypes error = errors.New("unit types given are incompatible.")
