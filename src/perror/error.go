// Package perror error messages
package perror

import "errors"

var (
	ErrSlotNumberInvalid        = errors.New("slot: Please provide valid slot number 1 or greater")
	ErrSlotAlreadyAllocated     = errors.New("slot: Slots already allocated")
	ErrParkingFullCapacity      = errors.New("Sorry, parking lot is full")
	ErrVehicleAssignInvalidSlot = errors.New("vehicle: Cannot allocate vehicle in invalid slot")
)
