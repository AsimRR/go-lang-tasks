package embedding

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Employee struct {
	EmployeeID int
	Position   string
	Person
}

func (p Employee) GetEmployeeDetails() {

	fullName := fmt.Sprintf("%s %s", p.FirstName, p.LastName)

	fmt.Printf("Name: %s \nEmployee ID: %d\nPosition: %s\n", fullName, p.EmployeeID, p.Position)
}
