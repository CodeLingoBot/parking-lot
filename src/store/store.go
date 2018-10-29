// Package store handling center for Parking center instance
package store

import (
	"config"
	"parking"
)

// Store keeps a global reference for Parking Center
type Store struct {
	parkingCenter *parking.Center
}

// store : local storage
var store *Store

// New : New store instance can be created
// 	Singleton instance creation and return
func New() *Store {
	if nil == store {
		store = new(Store).init()
	}
	return store
}

// init : initialise created object
func (s *Store) init() *Store {
	s.parkingCenter = parking.New(
		config.Start,
		config.Capacity,
	)
	return s
}

// GetParkingCenter : Return the current parking center instance
func (s *Store) GetParkingCenter() *parking.Center {
	return s.parkingCenter
}

// SetParkingCenter : Set new ParkingCenter object in memory
func (s *Store) SetParkingCenter(pC *parking.Center) {
	s.parkingCenter = pC
}
