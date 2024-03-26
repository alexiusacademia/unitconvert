package unitconvert_test

import (
	"math"
	"testing"

	"github.com/alexiusacademia/unitconvert"
)

func TestDischarge(t *testing.T) {
	v := 2.5

	result, err := unitconvert.Discharge(v, unitconvert.CubicMeterPerSecond, unitconvert.CubicFeetPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	result = math.Round(result*math.Pow(10, 4)) / math.Pow(10, 4)

	if result != 88.2867 {
		t.Errorf("Expected 88.2867, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	result, err = unitconvert.Discharge(v, unitconvert.CubicFeetPerSecond, unitconvert.CubicFeetPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	if result != result {
		t.Errorf("Expected %f, got %f\n", v, result)
	}

	result, err = unitconvert.Discharge(v, unitconvert.CubicFeetPerSecond, unitconvert.CubicMeterPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	result = math.Round(result*math.Pow(10, 4)) / math.Pow(10, 4)

	if result != 0.0708 {
		t.Errorf("Expected 0.0708, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}
}