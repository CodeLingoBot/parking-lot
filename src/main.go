package main

// Capacity of parking center data type
type Capacity uint64

// Index is slot index number type
type Index uint64

// Slot defines a Number and a Vehicle
// If vehicle object is allocated then slot is used
type Slot struct {
	Number  Index
	Vehicle *Vehicle
}

// ParkingCenter object holds the config of all required properties
type ParkingCenter struct {
	Capacity Capacity
	Counter,
	startIndex,
	allocationIndex Index
	slots []*Slot
}

func main() {

}
