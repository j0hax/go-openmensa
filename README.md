# go-openmensa

[![Go Report Card](https://goreportcard.com/badge/github.com/j0hax/go-openmensa)](https://goreportcard.com/report/github.com/j0hax/go-openmensa)

Go API for OpenMensa

## Install

The `openmensa` module functions purely as a library with no executables. To use it in a project, run

```console
go get github.com/j0hax/go-openmensa
```

## Example

Following code snippet fetches today's menu and prices for a cafeteria in Hannover:

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/j0hax/go-openmensa"
)

// openmensa.org ID for Contine Hannover
const contine = 7

func main() {
	date := time.Now().Format("2006-01-02")

	menu, err := openmensa.GetMeals(contine, date)

	if err != nil {
		log.Fatal(err)
	}

	for _, meal := range *menu {
		price := meal.Prices["students"]
		fmt.Printf("- %s (%0.2f€)\n", meal, price)
	}
}
```
