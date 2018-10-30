// Package commands test wich test functionality test for
// Command File processing
package commands

import (
	"io/ioutil"
	"os"
	"testing"
)

var cmdStrings = []byte(
	`create_parking_lot 6
	park KA-01-HH-1234 White
	park KA-01-HH-9999 White
	park KA-01-BB-0001 Black
	park KA-01-HH-7777 Red
	park KA-01-HH-2701 Blue
	park KA-01-HH-3141 Black
	leave 4
	status
	park KA-01-P-333 White
	park DL-12-AA-9999 White
	registration_numbers_for_cars_with_colour White
	slot_numbers_for_cars_with_colour White
	slot_number_for_registration_number KA-01-HH-3141
	slot_number_for_registration_number MH-04-AY-1111`)

func TestFileCommandProcessor_Process(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "test.data")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(cmdStrings); err != nil {
		t.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}

	err = NewFileCommandProcessor(tmpFile.Name()).Process()
	if err != nil {
		t.Error("Should not have error on process")
	}
}
