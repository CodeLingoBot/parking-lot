// Package parking manages the complete logic of a parking center
// Create, Add Vehicle, Remove vehicle, Search, Status Report etc.
package parking

import (
	"perror"
	"ptypes"
	"slot"
	"vehicle"
)

// Center object holds the config of all required properties
type Center struct {
	Capacity ptypes.Capacity
	Counter,
	startIndex,
	allocationIndex ptypes.Index
	slots []*slot.Slot
}

// New parking center instance
func New(start ptypes.Index, capacity ptypes.Capacity) *Center {
	return new(Center).Init(start, capacity)
}

// Init - Initialise parking center instance
func (pc *Center) Init(start ptypes.Index, capacity ptypes.Capacity) *Center {
	pc.Capacity = capacity
	pc.allocationIndex = 0
	pc.startIndex = start

	defer func() {
		recover()
	}()

	pc.slots = make([]*slot.Slot, uint64(capacity))
	for idx := range pc.slots {
		pc.slots[idx] = slot.New()
		pc.slots[idx].SetNumber(start)
		start = start + 1
	}

	return pc
}

// getFreeSlot get next serial slot or free in between
func (pc *Center) getFreeSlot() (*slot.Slot, error) {
	if !pc.allocatedAll() {
		return pc.nextSlot()
	}

	return pc.findNextFreeSlot()
}

// allocatedAll to check all slots are allocated at least once or not
func (pc *Center) allocatedAll() bool {
	return !(uint64(pc.Capacity) >= uint64(pc.allocationIndex)+1)
}

// nextSlot get next free slot
func (pc *Center) nextSlot() (*slot.Slot, error) {
	objSlot := pc.slots[pc.allocationIndex]
	pc.allocationIndex = pc.allocationIndex + 1

	return objSlot, nil
}

// findNextFreeSlot to get next free slot
func (pc *Center) findNextFreeSlot() (*slot.Slot, error) {
	for _, objSlot := range pc.slots {
		if objSlot.IsFree() {
			return objSlot, nil
		}
	}

	return nil, perror.ErrParkingFullCapacity
}

// getAllFreeSlot - get all free slots serial and in between
func (pc *Center) getAllFreeSlot() []*slot.Slot {
	freeSlots := make([]*slot.Slot, 0)
	for _, slot := range pc.slots {
		if slot.IsFree() {
			freeSlots = append(freeSlots, slot)
		}
	}

	return freeSlots
}

// AddVehicle - add vehicle to parking center
func (pc *Center) AddVehicle(vehicle *vehicle.Vehicle) (*slot.Slot, error) {
	var (
		err     error
		objSlot *slot.Slot
	)

	objSlot, err = pc.getFreeSlot()
	if err == nil && objSlot != nil {
		objSlot, err = objSlot.Allocate(vehicle)
		if err == nil {
			pc.Counter = pc.Counter + 1
		}
	}

	return objSlot, err
}

// remove - remove vehicle from center and decrement counter
func (pc *Center) remove(s *slot.Slot) {
	s.Free()
	pc.Counter = pc.Counter - 1
}
