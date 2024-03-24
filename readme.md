# unitconvert

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
