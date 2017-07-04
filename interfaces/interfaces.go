package main

import (
	"fmt"
	"reflect"
)

// Examine how interfaces work in conjunction with structs and methods
// The examples include electronic machinery and vehicles

// Machine represents any piece of machinery with a weight
type Machine struct {
	weight int
}

// GenerateElectricity returns true is the machine's weight is over 10
// NOTE: *Machine fully implements the "Generator" interface
func (m *Machine) GenerateElectricity() bool {
	if m.weight > 10 {
		fmt.Println("I am generating electricity.")
		return true
	}
	fmt.Println("I do not weigh enough to generate electricity.")
	return false
}

// AddValues adds two values and returns the result
// NOTE: *Machine does NOT fully implement the "Computer" interface
func (m *Machine) AddValues(value1, value2 int) int {
	result := value1 + value2
	fmt.Printf("Adding %d with %d yields: %d.\n", value1, value2, result)
	return result
}

type Vehicle struct {
	Machine
	mph, mpg int
	model    string
}

type Transport interface {
	LoadCargo(cargo string)
	WeighCargo() int
	Drive()
}
type Generator interface {
	GenerateElectricity() bool
}
type Computer interface {
	AddValues(value1, value2 int) int
	DetermineGPS(lat, long float32) float32
}
type Motorcycle interface {
	Drive()
}

func main() {
	bigMachine, smallMachine := &Machine{25}, &Machine{4} // create a big and small *Machine
	var bigGenerator, smallGenerator Generator

	// these WORK because "Machine" fully implements the "Generator" interface
	bigGenerator = bigMachine
	smallGenerator = smallMachine

	// examine how each different implementation of the same interface uses the GenerateElectricity() method
	x := []Generator{bigGenerator, smallGenerator}
	for _, value := range x {
		value.GenerateElectricity()
	}

	// examine how empty interfaces implement 0 methods, so any 'type' can also be an interface
	var blankInterface interface{}
	i := 14                                     // a random int
	s := "just a lonely string"                 // a random string
	v := Vehicle{Machine{300}, 85, 21, "Mazda"} // a random Vehicle
	blankInterface = i                          // this is valid
	blankInterface = s                          // so is this
	blankInterface = v                          // this too
	fmt.Println(blankInterface)                 // lets see what the interface is now

	// examine how to 'assert' that a certain type is stored in an interface variable (since it can potentially be anything)
	// 1. Assertion-ok pattern
	isAVehicle, ok := blankInterface.(Vehicle) // ok will be true if the interface variable contains the type of interface after the dot-assertion
	if ok {
		fmt.Printf("The assertion passed i.e. blankInterface is a Vehicle with the following attributes: %v\n", isAVehicle)
	} else {
		fmt.Printf("The assertion failed i.e. blankInterface is not currently an interface of type Vehicle.\n")
	}

	// 2. Reflection package (not necessary for all use cases)
	var reflectMachine interface{} = Machine{35}
	fmt.Println(reflect.TypeOf(reflectMachine))
}
