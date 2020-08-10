package stepper

import (
	"machine"
	"math"
	"time"
)

type StepperConfig struct {
	direction     int
	stepDelay     int64
	lastStepTime  int64
	numberOfSteps int64
	pinCount      int
	stepNumber    int64
}

type Stepper struct {
	motorPins []machine.Pin
	config    StepperConfig
}

func NewStepper(numberOfSteps int64, motorPins ...machine.Pin) *Stepper {
	s := new(Stepper)
	s.config.stepNumber = 0
	s.config.direction = 0
	s.config.lastStepTime = 0
	s.config.numberOfSteps = numberOfSteps
	s.motorPins = motorPins
	for _, pin := range s.motorPins {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	return s
}

/*
	Sets the speed in revs per minute
*/
func (s *Stepper) SetSpeed(speed int64) {
	s.config.stepDelay = 60 * 1000 * 1000 / s.config.numberOfSteps / speed
}

/*
	Move the motor steps_to_move steps. If the number is negative,
	the motor moves in the reverse direction
*/
func (s *Stepper) Step(stepsToMove float64) {
	stepsLeft := math.Abs(stepsToMove)

	if stepsToMove > 0 {
		s.config.direction = 1
	} else {
		s.config.direction = 0
	}

	for stepsLeft > 0 {
		now := time.Now()

		if time.Since(now).Microseconds()-s.config.lastStepTime >= s.config.stepDelay {
			// get the timeStamp of when you stepped
			s.config.lastStepTime = time.Since(now).Microseconds()

			// depending on direction, increment or decrement the step number
			if s.config.direction == 1 {
				// If we have reached the max limit (360 degrees), start over at 0
				if s.config.stepNumber == s.config.numberOfSteps {
					s.config.stepNumber = 0
				}

				s.config.stepNumber++
			} else {
				// If we have reached the min limit, start over at max step
				if s.config.stepNumber == 0 {
					s.config.stepNumber = s.config.numberOfSteps
				}
				s.config.stepNumber--
			}

			// decrement the steps left
			stepsLeft--

			// step the motor to step number 0, 1, ..., { 3 or 10 }
			s.stepMotor()

		}
	}
}

func (s *Stepper) stepMotor() {
	var step int64
	if len(s.motorPins) == 5 {
		step = s.config.stepNumber % 10
	} else {
		step = s.config.stepNumber % 4
	}
	nrOfPins := len(s.motorPins)

	if nrOfPins == 2 {
		switch step {
		case 0: //01
			s.motorPins[0].Low()
			s.motorPins[1].High()
		case 1: //11
			s.motorPins[0].High()
			s.motorPins[1].High()
		case 2: //10
			s.motorPins[0].High()
			s.motorPins[1].Low()
		case 3: //00
			s.motorPins[0].Low()
			s.motorPins[1].High()
		}
	} else if nrOfPins == 4 {
		switch step {
		case 0: //1010
			s.motorPins[0].High()
			s.motorPins[1].Low()
			s.motorPins[2].High()
			s.motorPins[3].Low()
		case 1: //0110
			s.motorPins[0].Low()
			s.motorPins[1].High()
			s.motorPins[2].High()
			s.motorPins[3].Low()
		case 2: //0101
			s.motorPins[0].Low()
			s.motorPins[1].High()
			s.motorPins[2].Low()
			s.motorPins[3].High()
		case 3: //1001
			s.motorPins[0].High()
			s.motorPins[1].Low()
			s.motorPins[2].Low()
			s.motorPins[3].High()
		}
	} else if nrOfPins == 5 {
		switch step {
		case 0: // 01101
			s.motorPins[0].Low()
			s.motorPins[1].High()
			s.motorPins[2].High()
			s.motorPins[3].Low()
			s.motorPins[4].High()
		case 1: // 01001
			s.motorPins[0].Low()
			s.motorPins[1].High()
			s.motorPins[2].Low()
			s.motorPins[3].Low()
			s.motorPins[4].High()
		case 2: // 01011
			s.motorPins[0].Low()
			s.motorPins[1].High()
			s.motorPins[2].Low()
			s.motorPins[3].High()
			s.motorPins[4].High()
		case 3: // 01010
			s.motorPins[0].Low()
			s.motorPins[1].High()
			s.motorPins[2].Low()
			s.motorPins[3].High()
			s.motorPins[4].Low()
		case 4: // 11010
			s.motorPins[0].High()
			s.motorPins[1].High()
			s.motorPins[2].Low()
			s.motorPins[3].High()
			s.motorPins[4].Low()
		case 5: // 10010
			s.motorPins[0].High()
			s.motorPins[1].Low()
			s.motorPins[2].Low()
			s.motorPins[3].High()
			s.motorPins[4].Low()
		case 6: // 10110
			s.motorPins[0].High()
			s.motorPins[1].Low()
			s.motorPins[2].High()
			s.motorPins[3].High()
			s.motorPins[4].Low()
		case 7: // 10100
			s.motorPins[0].High()
			s.motorPins[1].Low()
			s.motorPins[2].High()
			s.motorPins[3].Low()
			s.motorPins[4].Low()
		case 8: // 10101
			s.motorPins[0].High()
			s.motorPins[1].Low()
			s.motorPins[2].High()
			s.motorPins[3].Low()
			s.motorPins[4].High()
		case 9: // 00101
			s.motorPins[0].Low()
			s.motorPins[1].Low()
			s.motorPins[2].High()
			s.motorPins[3].Low()
			s.motorPins[4].High()
		}
	}
}
