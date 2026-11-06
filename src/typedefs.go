package fullcontrol

type ControllerModel int
type Controller struct {
	Model        ControllerModel
	MacAddress   string
	Connected    bool
	PlayerNumber int
	BatteryLevel int
}

type Dualshock3Buttons struct { 
	// Some buttons are uint8 and not bool because the buttons on the Dualshock 3 
	// controller were pressure sensitive and thus require numbers.

	Circle   uint8
	Cross    uint8
	Square   uint8
	Triangle uint8
	Up       uint8
	Left     uint8
	Down     uint8
	Right    uint8
	L1       uint8
	L2       uint8
	L3       bool
	R1       uint8
	R2       uint8
	R3       bool
	Start    bool
	Select   bool
	PS       bool
}

type Dualshock3Sticks struct {
	LeftX  float64
	LeftY  float64
	RightX float64
	RightY float64
}

type Dualshock3Controller struct {
	Info    Controller
	Buttons Dualshock3Buttons
	Sticks  Dualshock3Sticks
}

const (
	// Only playstation for now! These ones are the easiest from what I've heard
	Dualshock3 ControllerModel = iota
)
