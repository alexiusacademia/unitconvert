package unitconvert_test

import (
	"math"
	"testing"

	"github.com/alexiusacademia/unitconvert"
)

func TestVelocity(t *testing.T) {
	v := 2.5

	result, err := unitconvert.Velocity(v, unitconvert.MeterPerSecond, unitconvert.FeetPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	if math.Round(result*math.Pow(10, 3))/math.Pow(10, 3) != 8.202 {
		t.Errorf("Expected 8.2021, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
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
}
