package fullcontrol

import (
	"fmt"
	"math"
	"strings"
)

func (c *Dualshock3Controller) String() string {
	return fmt.Sprintf(
		strings.Replace(`
			=== Dualshock 3 Controller ===
			---          Info          ---
			Mac Address: %17s
			Connected:   %17t

			---    Buttons  Pressed    ---
			%sL2                          %sR2
			%sR1                          %sL1

			  %s↑   %sSTART       %sSELECT   %s△
			%s←   %s→         %sPS         %s□   %s◯
			  %s↓                        %s✕
			         %sL3        %sR3
			%s
			---         Sticks         ---
			Left Stick:  %17s
			Right Stick: %17s
            
		`, "\t", "", -1),
		c.Info.MacAddress,
		c.Info.Connected,
		buttonToAnsi(c.Buttons.L2),
		buttonToAnsi(c.Buttons.R2),
		buttonToAnsi(c.Buttons.L1),
		buttonToAnsi(c.Buttons.R1),
		buttonToAnsi(c.Buttons.Up),
		buttonToAnsi(c.Buttons.Start),
		buttonToAnsi(c.Buttons.Select),
		buttonToAnsi(c.Buttons.Triangle),
		buttonToAnsi(c.Buttons.Left),
		buttonToAnsi(c.Buttons.Right),
		buttonToAnsi(c.Buttons.PS),
		buttonToAnsi(c.Buttons.Square),
		buttonToAnsi(c.Buttons.Circle),
		buttonToAnsi(c.Buttons.Down),
		buttonToAnsi(c.Buttons.Cross),
		buttonToAnsi(c.Buttons.L3),
		buttonToAnsi(c.Buttons.R3),
		"\x1b[0m",
		fmt.Sprintf(
			"(%.3f, %.3f)", 
			c.Sticks.Left.X,
			c.Sticks.Left.Y,
		),
		fmt.Sprintf(
			"(%.3f, %.3f)", 
			c.Sticks.Right.X,
			c.Sticks.Right.Y,
		),
	)
}

func (c *Dualshock3Controller) TiltStick(x float64, y float64, stick ControllerStick) {
	vectorLength := math.Sqrt((x*x) + (y*y))
	if vectorLength > 1 {
		if stick == c.Sticks.Left {
        	c.Sticks.Left.X = x / vectorLength
        	c.Sticks.Left.Y = y / vectorLength
    	} else {
        	c.Sticks.Right.X = x
        	c.Sticks.Right.Y = y
    	}
	} else {
		if stick == c.Sticks.Left {
        	c.Sticks.Left.X = x
        	c.Sticks.Left.Y = y
    	} else {
        	c.Sticks.Right.X = x
        	c.Sticks.Right.Y = y
    	}
	}
}

func (c *Dualshock3Controller) PressButton(button any) {
	switch b := button.(type) {
    case *uint8:
        *b = 255
    case *bool:
        *b = true
    default:
        panic("Button must be *uint8 or *bool")
    }
}

func (c *Dualshock3Controller) ReleaseButton(button any) {
	switch b := button.(type) {
    case *uint8:
        *b = 0
    case *bool:
        *b = false
    default:
        panic("Button must be *uint8 or *bool")
    }
}
