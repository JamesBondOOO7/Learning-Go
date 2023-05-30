package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // an embedded field
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// do something
	return []Employee{}
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}

	fmt.Println(m.ID)
	fmt.Println(m.Description())
}
