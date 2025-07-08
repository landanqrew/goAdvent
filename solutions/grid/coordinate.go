package grid

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
	Char rune `json:"charRune"`
}

func (c Coordinate) GetCharAsStr() string {
	return string(c.Char)
}