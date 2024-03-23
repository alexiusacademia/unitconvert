package unitconvert

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
