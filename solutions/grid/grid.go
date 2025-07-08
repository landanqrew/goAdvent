package grid

import (
	"errors"
	"slices"
	"strings"
)

type Grid struct {
	GridMap map[Coordinate]string `json:"gridMap"`
}

func (g *Grid) BuildGrid(s string) error {
	if len(s) == 0 {
		return errors.New("cannot build grid. empty string provided")
	}
	var lines []string = strings.Split(s, "\n")
	slices.Reverse(lines)
	for y, line := range lines {
		for x, r := range line {
			c := Coordinate{
				X: x,
				Y: y,
				Char: r,
			}
			g.GridMap[c]=c.GetCharAsStr()
		}
	}
	return nil
}

