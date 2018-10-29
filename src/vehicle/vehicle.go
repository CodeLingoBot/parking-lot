package vehicle

// Vehicle stores number and Color of vehicle
type Vehicle struct {
	// Number and Color is defined as string
	// eg
	// 	Number: "KA-01-HH-01234"
	// 	Color: "White"
	Number, Color string
}

// New - Package based New Object creation function
// 	@params:
// 		Number: string
// 		Color: string
// 	@return:
// 		Vehicle: *Object
func New(number, color string) *Vehicle {
	return new(Vehicle).Init(number, color)
}

// Init - Initialise created object
// 	@params:
// 		Number: string
// 		Color: string
// 	@return:
// 		Vehicle: *Object
func (v *Vehicle) Init(number, color string) *Vehicle {
	v.Color = color
	v.Number = number

	return v
}

// GetColour - Get value of colour vehicle property
func (v *Vehicle) GetColour() string {
	return v.Color
}

// GetNumber - Get value of vehicle number property
func (v *Vehicle) GetNumber() string {
	return v.Number
}

// IsEquals - check object equality
func (v *Vehicle) IsEquals(vehicle *Vehicle) bool {
	return v.Number == vehicle.GetNumber() && v.GetColour() == vehicle.GetColour()
}
