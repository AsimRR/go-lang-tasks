package structs

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func CalculateRectangle(rectangle Rectangle) {
	perimetr := 2 * (rectangle.Width + rectangle.Height)
	area := rectangle.Width * rectangle.Height
	fmt.Printf("Width: %f , Height: %f \n Area: %f \nPerimeter: %f \n", rectangle.Width, rectangle.Height, area, perimetr)
}
