// Package commands with 'registration_numbers_for_cars_with_colour'
// command implementation
package commands

import (
	"fmt"
	"perror"
	"store"
	"strings"
)

// CmdGetRegNumWithColour defined arguments and related methods
type CmdGetRegNumWithColour struct {
	Command
	Color string
}

// NewCmdGetRegNumWithColour new command instance
func NewCmdGetRegNumWithColour() *CmdGetRegNumWithColour {
	var cmd = new(CmdGetRegNumWithColour)
	cmd.Cmd = "registration_numbers_for_cars_with_colour"
	return cmd
}

// Help to print command help information
func (crc *CmdGetRegNumWithColour) Help() string {
	return `registration_numbers_for_cars_with_colour <colour>
	Searching registration number of vehicle by using colour.
	Eg: registration_numbers_for_cars_with_colour White`
}

// Parse to parse arguments
func (crc *CmdGetRegNumWithColour) Parse(argString string) error {
	crc.Command.Parse(argString)
	crc.Color = crc.Args[0]
	return nil
}

// Verify to check the provided parameters are valid or not
func (crc *CmdGetRegNumWithColour) Verify() error {
	if perror.Empty == crc.Color {
		return perror.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (crc *CmdGetRegNumWithColour) Run() (string, error) {
	var outPutList = []string{}
	pC := store.Get().GetParkingCenter()
	slots, err := pC.ReportVehicleByColor(crc.Color)
	if nil == err {
		for _, s := range slots {
			v := s.GetVehicle()
			outPutList = append(
				outPutList,
				fmt.Sprintf("%v", v.GetNumber()),
			)
		}
	} else {
		outPutList = []string{
			err.Error(),
		}
	}
	crc.OutPut = strings.Join(outPutList, perror.Comma)
	return crc.OutPut, err
}
