package main

import (
	"fmt"

	"github.com/jeremyimmanuel/golang-bootcamp/printer"
)

// Ok if unused
var (
	// Ma is something
	Ma int
)

func main() {
	var (
		speed, velocity int
		heat            float64
	)
	safe := true
	var name = "jeremys"

	a, b := 1, 2

	fmt.Print(speed, heat, velocity, safe, name)

	text := "Jeremy"
	fmt.Print(text)
	printer.Lol()

	fmt.Print(a, b)
	a = 10
}
