package game

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Move struct {
	From *Position `json:"from"`
	To   *Position `json:"to"`
}
