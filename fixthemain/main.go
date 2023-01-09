package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	CLOSE = false
	OPEN  = true
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	if ptrDoor.state == CLOSE {
		return false
	} else {
		return true
	}
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	if ptrDoor.state == OPEN {
		return false
	} else {
		return true
	}
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) == true {
		OpenDoor(door)
	}
	if IsDoorOpen(door) == true {
		CloseDoor(door)
	}
}
