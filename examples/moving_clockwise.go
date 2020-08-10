package main

import (
	"machine"
	"github.com/niclaslind/tinystepper/stepper"
	)

func main() {
	motor := stepper.NewStepper(4, machine.D2, machine.D3, machine.D4, machine.D5)
}
