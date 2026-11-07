package fullcontrol

import (
	"fmt"
	"errors"
	"regexp"
	"runtime"
	"math/rand/v2"
)

func init() {
	os := runtime.GOOS
	if os != "linux" {
		panic("This module takes advantage of the linux bluetooth stack to emulate a bluetooth peripheral. Please use a linux machine to run this module!")
	}
}

func randomMacAddress() string, error {
	mac := ""
	for i := 0; i < 6; i++ {
		if i != 0 {
			mac += ":"
		}
		mac += fmt.Sprintf("%02X", rand.IntN(256))
	}
	if isValidMacAddress(mac) {
		return mac, nil
	} else {
		return "", errors.New("Failed to generate randomized mac address")
	}
}

macRegex := regexp.MustCompile(`^([0-9A-Fa-f]{2}:){5}[0-9A-Fa-f]{2}$`)
func isValidMacAddress(mac string) bool {
	return macRegex.MatchString(mac)
}

func CreateController(model ControllerModel) Controller, error {
	switch model {
	case Dualshock3:
		return Dualshock3Controller{
			Controller{
				Dualshock3,
				randomMacAddress(),
				false,
				0,
				238 // Charging
			},
			Dualshock3Buttons{},
			Dualshock3Sticks
		}, nil
	case default:
		return nil, errors.New("Unknown controller type")
	}
}
