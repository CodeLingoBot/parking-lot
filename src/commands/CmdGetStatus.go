// Package commands with 'status' command implementation
package commands

import (
	"fmt"
	"perror"
	"store"
	"strings"
)

// CmdGetStatus defined arguments and related methods
type CmdGetStatus struct {
	Command
}

// NewCmdGetStatus - new status command instance
func NewCmdGetStatus() *CmdGetStatus {
	var cmd = new(CmdGetStatus)
	cmd.Cmd = "status"
	return cmd
}

// Help to print help of status command
func (st *CmdGetStatus) Help() string {
	return `status
	Fetch all current parked vehicle details and slot numbers
	Eg: status`
}

// Parse to parse arguments
func (st *CmdGetStatus) Parse(argString string) error {
	st.Command.Parse(argString)
	return nil
}

// Verify to chech the provided parameteres are valid or not
func (st *CmdGetStatus) Verify() error {
	return nil
}

// Run to execute the command and provide result
func (st *CmdGetStatus) Run() (string, error) {
	var outPutList = []string{
		fmt.Sprintf("%-12s%-19s%-s", "Slot No.", "Registration No", "Colour"),
	}
	pC := store.Get().GetParkingCenter()
	slots, err := pC.ReportFilledSlots()
	if nil == err {
		for _, s := range slots {
			v := s.GetVehicle()
			outPutList = append(
				outPutList,
				fmt.Sprintf(
					"%-12v%-19v%-v",
					s.GetNumber(),
					v.GetNumber(),
					v.GetColour(),
				),
			)
		}
	} else {
		outPutList = []string{
			"No Data Found",
		}
	}
	st.OutPut = strings.Join(outPutList, perror.NewLine)
	return st.OutPut, nil
}
