package fullcontrol

import (
	"cmp"
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

func Clamp[T cmp.Ordered](val, minVal, maxVal T) T {
	if val > maxVal {
		return maxVal
	} else if val < minVal {
		return minVal
	} else {
		return val
	}
}

func randomMacAddress() (string, error) {
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

var macRegex *regexp.Regexp = regexp.MustCompile(`^([0-9A-Fa-f]{2}:){5}[0-9A-Fa-f]{2}$`)
func isValidMacAddress(mac string) bool {
	return macRegex.MatchString(mac)
}

func CreateController(model ControllerModel) (any, error) {
	controllerMac, err := randomMacAddress()
	if err != nil {
		return nil, err
	}
	switch model {
	case Dualshock3:
		return &Dualshock3Controller{
			Controller{
				Dualshock3,
				controllerMac,
				false,
				0,
				238, // Charging
			},
			Dualshock3Buttons{},
			Dualshock3Sticks{},
		}, nil
	default:
		return nil, errors.New("Unknown controller type")
	}
}

func buttonToAnsi(button any) string {
	// uint8() and bool() are neccessary here or this function
	// fails to compile
	switch v := button.(type) {
    case uint8:
        if v == 0 {
			return "\x1b[0;31m"
		} else if v == 255 {
			return "\x1b[1;32m"
		} else {
			return "\x1b[1;34m"
		}
    case bool:
        if v {
			return "\x1b[1;32m"
		} else {
			return "\x1b[0;31m"
		}
    default:
        panic("Argument button must be either type uint8 or bool!")
    }	
}
