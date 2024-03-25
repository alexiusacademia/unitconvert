package unitconvert

import "math"

const meterToFeet float64 = 3.28084
const kilometerToMile float64 = 0.621371

// conversionFactors is a map that stores conversion factors between different units.
// It maps from a unit (source) to another map, which in turn maps to the destination unit and its conversion factor.
// For example, conversionFactors[Meter][Feet] gives the conversion factor from meters to feet.
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
		FeetPerSecond:    meterToFeet,
		KilometerPerHour: 1 / 0.277778,
		MilesPerHour:     2.236942,
	},
	FeetPerSecond: {
		MeterPerSecond:   1 / meterToFeet,
		KilometerPerHour: 1 / 0.911344,
		MilesPerHour:     1 / 1.46667,
	},
	KilometerPerHour: {
		MilesPerHour:   kilometerToMile,
		MeterPerSecond: 0.277778,
		FeetPerSecond:  0.911344,
	},
	MilesPerHour: {
		KilometerPerHour: 1 / kilometerToMile,
		MeterPerSecond:   0.447038,
		FeetPerSecond:    1.466667,
	},
	// Add more here
}
