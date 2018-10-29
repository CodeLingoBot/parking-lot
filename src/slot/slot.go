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
	return new(Slot).Init()
}

// Init - Initialise Object with default values
func (sl *Slot) Init() *Slot {
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
