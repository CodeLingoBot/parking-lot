// Package perror error messages
package perror

import "errors"

var (
	ErrSlotNumberInvalid = errors.New("slot: Please provide valid slot number 1 or greater")
)
