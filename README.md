# go-openmensa
[![Go](https://github.com/j0hax/go-openmensa/actions/workflows/go.yml/badge.svg)](https://github.com/j0hax/go-openmensa/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/j0hax/go-openmensa.svg)](https://pkg.go.dev/github.com/j0hax/go-openmensa)
[![Go Report Card](https://goreportcard.com/badge/github.com/j0hax/go-openmensa)](https://goreportcard.com/report/github.com/j0hax/go-openmensa)

Go API for OpenMensa

## Install

The `openmensa` module functions purely as a library with no executables. To use it in a project, run

```console
$ go get github.com/j0hax/go-openmensa
```

## Example

Following code snippet fetches today's menu and prices for a cafeteria in Hannover:

```go
package main

import (
	"fmt"
	"github.com/j0hax/go-openmensa"
	"log"
)

func main() {
	// Contine Hannover has ID 7
	contine, err := openmensa.GetCanteen(7)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the current menu
	menu, err := contine.CurrentMenu()
	if err != nil {
		log.Fatal(err)
	}

	// Print out structured data
	fmt.Printf("%s: %s\n", contine.Name, menu.Day)
	for _, meal := range menu.Meals {
		price := meal.Prices["students"]
		fmt.Printf("- %s (%0.2f€)\n", meal, price)
	}
}
```

## See Also

* OpenMensa's official [API Documentation](https://docs.openmensa.org/)
* [kiliankoe/openmensa](https://github.com/kiliankoe/openmensa), which I only discovered long after starting this project. It appears to be unmaintained.
