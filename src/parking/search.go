// Package parking manages the complete logic of a parking center
// Search, Status Report etc. are implemented here
package parking

import (
	"perror"
	"ptypes"
	"slot"
)

// ReportVehicleByColor Get filled slots with vehicle colour
func (pc *Center) ReportVehicleByColor(color string) ([]*slot.Slot, error) {
	oSlots, err := pc.GetSlotsBy("color", color)
	if perror.Zero == len(oSlots) {
		err = perror.ErrNotFound
	}
	return oSlots, err
}

// ReportVehicleByNumber Get filled slots with vehicle number
func (pc *Center) ReportVehicleByNumber(number string) ([]*slot.Slot, error) {
	oSlots, err := pc.GetSlotsBy("number", number)
	if perror.Zero == len(oSlots) {
		err = perror.ErrNotFound
	}
	return oSlots, err
}

// ReportFreeSlots Get all free slots in center
func (pc *Center) ReportFreeSlots() []*slot.Slot {
	return pc.getAllFreeSlot()
}

// ReportFilledSlots Get all filled slots
func (pc *Center) ReportFilledSlots() ([]*slot.Slot, error) {
	allocSlots := make([]*slot.Slot, 0)
	if ptypes.Index(perror.Zero) == pc.Counter {
		return nil, perror.ErrNoFilledSlots
	}
	for _, s := range pc.slots {
		if !s.IsFree() {
			allocSlots = append(allocSlots, s)
		}
	}
	return allocSlots, nil
}
