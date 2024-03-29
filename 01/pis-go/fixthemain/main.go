package main

import "github.com/01-edu/z01"

// ------------
// package piscine
// ---------------------------------
func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// -------------------------------
func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = "OPEN"
	// return true
}

// ---------------------------------------
func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = "CLOSE"
	// return true
}

// -----------------------------------------
func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?\n")
	return Door.state == "OPEN"
}

// -----------------------------------------
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state == "CLOSE"
}

// -----------------------------------------
type Door struct {
	state string
}

// --------------
func main() {
	door := &Door{}

	OpenDoor(door)

	if IsDoorClose(door) { // bool
		OpenDoor(door)
	}

	if IsDoorOpen(door) { // bool
		CloseDoor(door)
	}

	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
