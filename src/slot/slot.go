package slot

import (
	"perror"
	"ptypes"
	"vehicle"
)

// SlotNumberLowerLimit is slot lower bound defined as a constant
const SlotNumberLowerLimit = 1

// Slot defines a Number and a Vehicle
// If vehicle object is allocated then slot is used
type Slot struct {
	Number  ptypes.Index
	Vehicle *vehicle.Vehicle
}

// New Slot Object creation function
func New() *Slot {
	return new(Slot).init()
}

// init - initialise Object with default values
func (sl *Slot) init() *Slot {
	sl.Number = SlotNumberLowerLimit - 1
	sl.Vehicle = nil

	return sl
}

// SetNumber - Set slot number to slot object
func (sl *Slot) SetNumber(number ptypes.Index) (*Slot, error) {
	if !IsValidSlotNumber(number) {
		return sl, perror.ErrSlotNumberInvalid
	}

	sl.Number = number
	return sl, nil
}

// IsValidSlotNumber - help to check the slot number is valid or not
func IsValidSlotNumber(Number ptypes.Index) bool {
	return (Number >= SlotNumberLowerLimit)
}

// GetNumber - get eslot number from slot object
func (sl *Slot) GetNumber() ptypes.Index {
	return sl.Number
}

// IsFree Verifies that slot is free or not, if no vehicle allocated
// then vehicle property will be nil
func (sl *Slot) IsFree() bool {
	return sl.Vehicle == nil
}

// Allocate set a vehicle object to slot, so that slot will be used
// Slot without valid slot number should show error
// Already using slot should not reused until slot is free
func (sl *Slot) Allocate(vehicle *vehicle.Vehicle) (*Slot, error) {
	if !sl.IsValid() {
		return sl, perror.ErrVehicleAssignInvalidSlot
	}

	if nil != sl.Vehicle {
		return sl, perror.ErrSlotAlreadyAllocated
	}

	sl.Vehicle = vehicle
	return sl, nil
}

// IsValid help to check the slot is valid or not
// Mainly check slot number allocated or not
func (sl *Slot) IsValid() bool {
	return sl.Number >= SlotNumberLowerLimit
}

// Free remove vehicle object from slot
func (sl *Slot) Free() *Slot {
	sl.Vehicle = nil
	return sl
}

// GetVehicle - get vehicle object from allocated slot
func (sl *Slot) GetVehicle() *vehicle.Vehicle {
	return sl.Vehicle
}
