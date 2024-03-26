# unitconvert

[![Go Reference](https://pkg.go.dev/badge/github.com/alexiusacademia/unitconvert.svg)](https://pkg.go.dev/github.com/alexiusacademia/unitconvert)

`unitconvert` let's you convert quantities from one unit to another for engineering calculations.

Each unit type is restricted so that you can only convert units of the same type (e.g. length to length) without accidentally converting different unit types (e.g. length to area).

This library tries to exclude unit conversions that are already available in the [https://github.com/bcicen/go-units](https://github.com/bcicen/go-units) repository.

## Installation

Like any standard go module installation:

```
go get github.com/alexiusacademia/unitconvert
```

If you want to update the package, just go with

```
go get -u github.com/alexiusacademia/unitconvert
```

## Sample

1. In this example, we are going to convert 5.5 feet to meters.

```
package main

import (
	"fmt"

	uc "github.com/alexiusacademia/unitconvert"
)

func main() {
	value := 5.5 // Sample value in feet

	// Let's convert it to meter
	newValue, err := uc.Length(value, uc.Feet, uc.Meter)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The converted value from meter to feet is: %f\n", newValue)
}
```

The output shoud be 1.6764

## Run tests

To run all the tests in the library, run this in the root of this codebase.

```
go test -v ./...
```

## Contributing

Firstly, fork the repository and clone to your own machine.

> To contribute to a selected version, create a new branch (e.g. v0.5.0) and then work from there. To submit the version for review, create a pull request. Make sure to follow the main feature indicated in the `roadmap.md` file.

To avoid conflict in pushing branch name in a tag format, please create your branch names in the format v000 for v0.0.0. Since this creates issues when pushing changes to the branch due to same command for pushing branch and pushing tag to the origin.

### Pushing Updates/Changes

Before pushing your changes to your branch, first fetch the changes from the original repository then merge the updates to your branch to minimize or eliminate conflicts.

### Providing Tests

Tests should be provided in case it is not provided by default. The number of tests in case of conversion functions should vary depending on the number of conversions per UnitType.

For example, you have 4 conversions in a unit type, there should be 16. So the sqaure of the number of conversions from the `conversionFactors` variable.

### conversionFactors variable

If you are adding conversion factors from the `conversionFactors` variable, please take note of the following:

- The number of conversion factor per unit shall be the same for all `Unit`s. Meaning if you are converting a `Unit` on a `UnitType` with 2 other units, you should have 3 unit conversions for that `UnitType` and 2 conversion factors for each of those `Unit`s in the `UnitType` like in the code below:

```go
Meter: {
	Feet: meterToFeet,
	Kilometer: 1000,
},
Kilometer: {
	Meter: 1/1000,
	Feet: 3280.84,
},
Feet: {
	Meter: 1 / meterToFeet,
	Kilometer: 0.0003048,
},
```
