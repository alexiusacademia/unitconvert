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

	result, err = unitconvert.Discharge(v, unitconvert.LiterPerSecond, unitconvert.CubicMeterPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	result = math.Round(result*math.Pow(10, 4)) / math.Pow(10, 4)

	if result != 0.0025 {
		t.Errorf("Expected 0.0025, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	result, err = unitconvert.Discharge(v, unitconvert.LiterPerSecond, unitconvert.LiterPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	result = math.Round(result*math.Pow(10, 1)) / math.Pow(10, 1)

	if result != 2.5 {
		t.Errorf("Expected 2.5, got %f.\n", math.Round(result*math.Pow(10, 1))/math.Pow(10, 1))
	}

	result, err = unitconvert.Discharge(v, unitconvert.CubicMeterPerSecond, unitconvert.LiterPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	result = math.Round(result*math.Pow(10, 3)) / math.Pow(10, 3)

	if result != 2500.000 {
		t.Errorf("Expected 2500.000, got %f.\n", math.Round(result*math.Pow(10, 3))/math.Pow(10, 3))
	}

	result, err = unitconvert.Discharge(v, unitconvert.LiterPerSecond, unitconvert.CubicFeetPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	result = math.Round(result*math.Pow(10, 4)) / math.Pow(10, 4)

	if result != 0.0883 {
		t.Errorf("Expected 0.0882867, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	result, err = unitconvert.Discharge(v, unitconvert.CubicFeetPerSecond, unitconvert.LiterPerSecond)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	result = math.Round(result*math.Pow(10, 1)) / math.Pow(10, 1)

	if result != 70.8 {
		t.Errorf("Expected 70.8, got %f.\n", math.Round(result*math.Pow(10, 1))/math.Pow(10, 1))
	}
}
