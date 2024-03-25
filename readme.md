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

## Contributing

Firstly, fork the repository and clone to your own machine.

> To contribute to a selected version, create a new branch (e.g. v0.5.0) and then work from there. To submit the version for review, create a pull request. Make sure to follow the main feature indicated in the `roadmap.md` file.

To avoid conflict in pushing branch name in a tag format, please create your branch names in the format v000 for v0.0.0. Since this creates issues when pushing changes to the branch due to same command for pushing branch and pushing tag to the origin.

### Pushing Updates/Changes

Before pushing your changes to your branch, first fetch the changes from the original repository then merge the updates to your branch to minimize or eliminate conflicts.
