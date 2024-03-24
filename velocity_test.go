package unitconvert_test

import (
	"math"
	"testing"

	"github.com/alexiusacademia/unitconvert"
)

func TestVelocity(t *testing.T) {
	v := 2.5

	// From meters per second
	// Convert mps -> fps
	result, err := unitconvert.Velocity(v, unitconvert.MeterPerSecond, unitconvert.FeetPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	if math.Round(result*math.Pow(10, 3))/math.Pow(10, 3) != 8.202 {
		t.Errorf("Expected 8.2021, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	// Convert mps -> kph
	result, err = unitconvert.Velocity(v, unitconvert.MeterPerSecond, unitconvert.KilometerPerHour)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	if math.Round(result*math.Pow(10, 3))/math.Pow(10, 3) != 9.0 {
		t.Errorf("Expected 9.0, got %f.\n", math.Round(result*math.Pow(10, 3))/math.Pow(10, 3))
	}

	// Convert mps -> mph
	result, err = unitconvert.Velocity(v, unitconvert.MeterPerSecond, unitconvert.MilesPerHour)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	if math.Round(result*math.Pow(10, 4))/math.Pow(10, 4) != 5.5924 {
		t.Errorf("Expected 5.5924, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	result, err = unitconvert.Velocity(v, unitconvert.FeetPerSecond, unitconvert.FeetPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	if result != result {
		t.Errorf("Expected %f, got %f\n", v, result)
	}

	result, err = unitconvert.Velocity(v, unitconvert.FeetPerSecond, unitconvert.MeterPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	if math.Round(result*math.Pow(10, 3))/math.Pow(10, 3) != 0.762 {
		t.Errorf("Expected 0.762, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	// Convert kph -> mph
	result, err = unitconvert.Velocity(v, unitconvert.KilometerPerHour, unitconvert.MilesPerHour)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}
	if math.Round(result*math.Pow(10, 3))/math.Pow(10, 3) != 1.553 {
		t.Errorf("Expected 1.553, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	// Convert kph -> meter per second
	result, err = unitconvert.Velocity(v, unitconvert.KilometerPerHour, unitconvert.MeterPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}
	if math.Round(result*math.Pow(10, 3))/math.Pow(10, 3) != 0.694 {
		t.Errorf("Expected 0.694, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	// Convert kph -> fps
	result, err = unitconvert.Velocity(v, unitconvert.KilometerPerHour, unitconvert.FeetPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}
	if math.Round(result*math.Pow(10, 3))/math.Pow(10, 3) != 2.278 {
		t.Errorf("Expected 2.278, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	// Convert mph -> kph
	result, err = unitconvert.Velocity(v, unitconvert.MilesPerHour, unitconvert.KilometerPerHour)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}
	if math.Round(result*math.Pow(10, 3))/math.Pow(10, 3) != 4.023 {
		t.Errorf("Expected 4.023, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	// Convert mph -> meters per second
	result, err = unitconvert.Velocity(v, unitconvert.MilesPerHour, unitconvert.MeterPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}
	if math.Round(result*math.Pow(10, 4))/math.Pow(10, 4) != 1.1176 {
		t.Errorf("Expected 1.117, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	// Convert mph -> feet per second
	result, err = unitconvert.Velocity(v, unitconvert.MilesPerHour, unitconvert.FeetPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}
	if math.Round(result*math.Pow(10, 4))/math.Pow(10, 4) != 3.6667 {
		t.Errorf("Expected 3.6667, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}
}
