package unitconvert_test

import (
	"math"
	"testing"

	"github.com/alexiusacademia/unitconvert"
)

func TestArea(t *testing.T) {
	v := 2.5

	result, err := unitconvert.Area(v, unitconvert.SquareMeter, unitconvert.SquareFeet)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	result = math.Round(result*math.Pow(10, 4)) / math.Pow(10, 4)

	if result != 26.9098 {
		t.Errorf("Expected 26.9098, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}

	result, err = unitconvert.Area(v, unitconvert.SquareFeet, unitconvert.SquareFeet)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	if result != result {
		t.Errorf("Expected %f, got %f\n", v, result)
	}

	result, err = unitconvert.Area(v, unitconvert.SquareFeet, unitconvert.SquareMeter)
	if err != nil {
		t.Errorf("An error has occured: %s\n", err)
	}

	result = math.Round(result*math.Pow(10, 3)) / math.Pow(10, 3)

	if result != 0.232 {
		t.Errorf("Expected 0.232, got %f.\n", math.Round(result*math.Pow(10, 4))/math.Pow(10, 4))
	}
}
