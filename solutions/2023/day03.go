package solutions2023

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/landanqrew/goAdvent/osUtils"
)

var d3p1Example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var nonSymbols map[string]bool = map[string]bool{
	"1": true,
	"2": true,
	"3": true,
	"4": true,
	"5": true,
	"6": true,
	"7": true,
	"8": true,
	"9": true,
	"0": true,
	".": false,
}

type coordinate struct{
	x int
	y int
}

func (c coordinate) isNumeric(grid [][]string) bool {
	char := grid[c.y][c.x]
	v, ok := nonSymbols[char]
	if ok && v {
		return true
	}
	return false
}

func (c coordinate) isSymbol(grid [][]string) bool {
	char := grid[c.y][c.x]
	_, ok := nonSymbols[char]
	return !ok
}

func (c coordinate) hasAdjacentSymbol(num string, grid [][]string) bool {
	for y := c.y - 1; y < c.y + 2; y++ {
		for x := c.x - 1; x < c.x + len(num) + 1; x++ {
			if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
				char := grid[y][x]
				_, ok := nonSymbols[char]
				if !ok {
					return true
				}
			}
		}
	}
	return false
}

type gear struct {
	nums []int
	coor coordinate
	gearRatio int
}

func (g *gear) getGearRatio(nums map[coordinate]string, grid [][]string) {
	checked := map[coordinate]bool{}
	for x := g.coor.x - 1; x <= g.coor.x + 1; x++ {
		for y := g.coor.y - 1; y <= g.coor.y + 1; y++ {
			if x >= 0 && x < len(grid[y]) && y >= 0 && y < len(grid) {
				c := coordinate{x, y}
				if c.isNumeric(grid) && !checked[c] {
					num := getNum(nums, grid, c, checked)
					if num > 0 {
						g.nums = append(g.nums, num)
					}
				}
			}
		}
	}
	if len(g.nums) == 2 {
		g.gearRatio = g.nums[0] * g.nums[1]
	}
}

func getNum(nums map[coordinate]string, grid [][]string, coor coordinate, checked map[coordinate]bool) int {
	if checked[coor] {
		return 0
	}
	_, ok := nums[coor]
	checked[coor] = true
	if ok {
		numInt, _ := strconv.Atoi(nums[coor])
		return numInt
	}
	return getNum(nums, grid, coordinate{coor.x - 1, coor.y}, checked)
}




func Day3(useExample bool, debug bool) {
	puzzle := string(osUtils.ReadFileFromPath("puzzles/2023/day3.txt"))
	if useExample {
		puzzle = d3p1Example
	}

	lines := strings.Split(puzzle, "\n")
	nums := map[coordinate]string{}
	gears := []gear{}

	grid := [][]string{}
	for y, line := range lines {
		appendNumbers(line, y, nums)
		chars := strings.Split(line, "")
		grid = append(grid, chars)
		for x, char := range chars {
			if char == "*" {
				gears = append(gears, gear{coor: coordinate{x, y}})
			}
		}
		/*num := ""
		for x, char := range chars {
			//update
			// grid[y][x] = char
			if char != "." {
				coor := coordinate{x, y}
				if coor.isNumeric(grid) {
					num += char
				} else {
					if num != "" {
						numInt, _ := strconv.Atoi(num)
						nums[coordinate{x, y}] = numInt
						num = ""
					}
					num = ""
				}
			}
		}
		if num != "" {
			numInt, _ := strconv.Atoi(num)
			nums[coordinate{len(chars) - 1, y}] = numInt
			num = ""
		}*/
	}


	p1total := 0
	for _, coor := range sortedKeys(nums) {
		num := nums[coor]
		/*if debug {
			fmt.Println(coor, num)
		}*/
		//convert to int
		if coor.hasAdjacentSymbol(num, grid) {
			numInt, _ := strconv.Atoi(num)
			p1total += numInt
			if debug && coor.y <= 5 {
				fmt.Println("Adding ", num, " to total")
			}
		}
	}

	p2total := 0
	for _, gear := range gears {
		gear.getGearRatio(nums, grid)
		p2total += gear.gearRatio
	}

	fmt.Println("Part 1: ", p1total, "Part 2: ", p2total)
}

func appendNumbers(line string, y int, nums map[coordinate]string) {
	num := ""
	currentNumStartX := -1 // Keep track of the start of the current number being built
	for x, r := range line {
		char := string(r)
		// Check if the character is a digit
		isDigit := false
		if val, ok := nonSymbols[char]; ok && val { // It's in nonSymbols and true (is a digit)
			isDigit = true
		}

		if isDigit {
			if num == "" { // Starting a new number
				currentNumStartX = x
			}
			num += char
		} else { // Character is not a digit (it's a symbol or '.')
			if num != "" { // We just finished a number
				nums[coordinate{currentNumStartX, y}] = num
			}
			num = "" // Reset for the next potential number
			currentNumStartX = -1
		}
	}
	// After the loop, if num is not empty, it means the line ended with a number
	if num != "" {
		// currentNumStartX should already hold the correct starting X
		nums[coordinate{currentNumStartX, y}] = num
	}
}

func sortedKeys(m map[coordinate]string) []coordinate {
	keys := make([]coordinate, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Sort the slice of coordinates
	sort.Slice(keys, func(i, j int) bool {
		// Primary sort condition: by y-coordinate
		if keys[i].y != keys[j].y {
			return keys[i].y < keys[j].y
		}
		// Secondary sort condition: if y-coordinates are equal, sort by x-coordinate
		return keys[i].x < keys[j].x
	})

	return keys
}