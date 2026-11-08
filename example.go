package main

import (
	"fmt"
	fc "fullcontrol/fullcontrol"
)

func main() {
	defer fc.Close()

	c, err := fc.CreateController(fc.Dualshock3) // Create controller object
	if err != nil {
		panic(err)
	}
	ds3, ok := c.(*fc.Dualshock3Controller) // Assert typing
	if !ok {
    	panic("Wrong controller type")
	}

	alias, err := fc.GetAdapterAlias()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bluetooth Alias: %s\n", alias)
	
	if err != nil {
		panic(err)
	}

	err = ds3.Connect()
	if err != nil {
		panic(err)
	}
	
	ds3.TiltStick(0.5, 1.3, ds3.Sticks.Left) // Tilt stick to the right and very far up
	ds3.PressButton(&ds3.Buttons.Triangle) // Fully press triangle
	ds3.PressButton(&ds3.Buttons.Square, 0.5) // Press square halfway
	fmt.Printf("%s", ds3) // Print controller statue
}
