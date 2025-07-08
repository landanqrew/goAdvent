package solutions2023

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/landanqrew/goAdvent/osUtils"
)

type game struct {
	Sets []set `json:"sets"`
	ID int `json:"id"`
}

func (g game) print() {
	json, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		log.Fatal("cannot marshal game: ", err)
	}
	fmt.Println(string(json))
}

func (g game) getPower() (minRed int, minGreen int, minBlue int) {
	for _, set := range g.Sets {
		if minRed == 0 || set.Red > minRed {
			minRed = set.Red
		}
		if minGreen == 0 || set.Green > minGreen {
			minGreen = set.Green
		}
		if minBlue == 0 || set.Blue > minBlue {
			minBlue = set.Blue
		}
	}
	return minRed, minGreen, minBlue
}

/*func (g *game) toMap() map[string]interface{} {
	setMaps := []map[string]int{}
	for _, set := range g.sets {
		setMaps = append(setMaps, map[string]int{
			"red": set.red,
			"green": set.green,
			"blue": set.blue,
		})
	}
	gMap := map[string]interface{}{}
	gMap["id"] = g.id
	gMap["sets"] = setMaps
	return gMap
}*/

type set struct {
	Red int `json:"red"`
	Green int `json:"green"`
	Blue int `json:"blue"`
}

func (s set) print() {
	json, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatal("cannot marshal set: ", err)
	}
	fmt.Println(string(json))
}

func (gm game) isPossible(r int, g int, b int) bool {
	for _, set := range gm.Sets {
		if set.Red > r || set.Green > g || set.Blue > b {
			return false
		}
	}
	return true
}

var p1Example string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func Day2(useExample bool, debug bool) {
	puzzle := string(osUtils.ReadFileFromPath("puzzles/2023/day2.txt"))
	if useExample {
		puzzle = p1Example
	} 

	lines := strings.Split(puzzle, "\n")
	games := []game{}
	for _, line := range lines {
		game := parseGame(line)
		games = append(games, game)
	}

	if debug {
		for _, game := range games {
			fmt.Println("Game (",game.ID,"):")
			game.print()
		}
	}

	p1total := 0
	p2total := 0
	for _, game := range games {
		minRed, minGreen, minBlue := game.getPower()
		p2total += minRed * minGreen * minBlue
		if game.isPossible(12, 13, 14) {
			p1total += game.ID
		}
	}
	fmt.Println("Part 1: ", p1total, "Part 2: ", p2total)
}

func parseGame(line string) game {
	// get id
	parts := strings.Split(line, ":")
	id := extractNumbers(parts[0])[0]

	// get sets
	sets := extractSets(parts[1])

	// return game
	return game{
		ID: id,
		Sets: sets,
	}
}

func extractSets(line string) []set {
	setStrings := strings.Split(line, ";")
	
	sets := []set{}
	for _, setString := range setStrings {
		s := set{}
		items := strings.Split(setString, ",")
		for _, item := range items {
			alphaChars := extractAlphaChars(item)
			nums := extractNumbers(item)
			switch alphaChars[0] {
			case "red":
				s.Red = nums[0]
			case "green":
				s.Green = nums[0]
			case "blue":
				s.Blue = nums[0]
			default:
				continue
			}
		}

		sets = append(sets, s)
	}
	return sets
}

func extractNumbers(numString string) []int {
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		log.Fatal("cannot compile regex: ", err)
	}
	matches := re.FindAllString(numString, -1)
	nums := []int{}
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			log.Fatal("cannot convert string to int: ", err)
		}
		nums = append(nums, num)
	}
	return nums
}

func extractAlphaChars(str string) []string {
	re, err := regexp.Compile(`[a-zA-Z]+`)
	if err != nil {
		log.Fatal("cannot compile regex: ", err)
	}
	matches := re.FindAllString(str, -1)
	return matches
}




