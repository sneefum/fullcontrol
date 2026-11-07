package main

import (
	"fmt"
	fc "fullcontrol/fullcontrol"
)

func main() {
	c, err := fc.CreateController(fc.Dualshock3)
	if err != nil {
		panic(err)
	}
	ds3, ok := c.(*fc.Dualshock3Controller)
	if !ok {
    	panic("Wrong controller type")
	}
	fmt.Printf("Actual type: %T\n", ds3)
	ds3.TiltStick(0.5, 1.3, ds3.Sticks.Left)
	ds3.PressButton(&ds3.Buttons.Triangle)
	fmt.Printf("%s", ds3)
}
