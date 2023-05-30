package main

import "fmt"

type Bike struct {
	owner string
}

type Car struct {
	owner string
}

func (b Bike) drive() {
	fmt.Println("Driving a bike")
}

func (c Car) drive() {
	fmt.Println("Driving a car")
}

func (c Car) reverse() {
	fmt.Println("Reversing the car back")
}

type Vehicle interface {
	drive()
}

func main() {
	vehicles := []Vehicle{
		Car{
			owner: "James",
		},
		Bike{
			owner: "Ethan",
		},
	}

	for _, vehicle := range vehicles {
		vehicle.drive()
		/* calling the reverse() method on vehicle
		but, it is specific to `Car` only
		while vehicle is an interface type

		Hence, we need to assert that the vehicle should be `Car`
		in order to call the reverse method
		*/
		car, ok := vehicle.(Car)
		if ok {
			car.reverse()
		}

		/*
			Idiomatic way :)
			if car, ok := vehicle.(Car); ok {
					car.reverse()
			}
		*/

	}
}
