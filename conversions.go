package unitconvert

import "math"

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
	SquareMeter: {
		SquareFeet: math.Pow(meterToFeet, 2),
		// Add more here
	},
	SquareFeet: {
		SquareMeter: 1 / math.Pow(meterToFeet, 2),
		// Add more here
	},
	MeterPerSecond: {
		FeetPerSecond: meterToFeet,
	},
	FeetPerSecond: {
		MeterPerSecond: 1 / meterToFeet,
	},
	// Add more here
}
