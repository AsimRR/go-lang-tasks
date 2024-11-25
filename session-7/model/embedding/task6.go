package embedding

import "fmt"

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicle
	FuelType string
}

func (c *Car) GetCarDetails() {
	fmt.Printf("Make: %s \nModel: %s\nYear: %d\nFuel Type: %s\n", c.Make, c.Model, c.Year, c.FuelType)
}
