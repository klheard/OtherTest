package test

type Ship struct {
	Name string
}

type Boat struct {
	Name string
}

type CruiseShip struct {
	Ship
	Boat
}
