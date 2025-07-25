package game

type Positoin struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Move struct {
	From Positoin `json:"from"`
	To   Positoin `json:"to"`
}
