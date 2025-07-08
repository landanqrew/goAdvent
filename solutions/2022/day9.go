package solutions2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/landanqrew/goAdvent/osUtils"
)

type knot struct {
	coor [2]int
	last [2]int
}

func readPuzzle() string {
	sb := osUtils.ReadFileFromPath("/Users/landanquartemont/Desktop/Development/go_projects/goAdvent/puzzles/2022/day9.txt")
	return string(sb)
} 

func Day9(useExample bool) {
	startTS := time.Now()
	dirMap := map[string][2]int{
		"U": {0, 1},
		"D": {0, -1},
		"L": {-1, 0},
		"R": {1, 0},
	}
	puzzle := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	puzzle = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	if !useExample {
		puzzle = readPuzzle()
	}
	tailVisited := map[[2]int]struct{}{}

	head, tail, last := [2]int{0,0}, [2]int{0,0}, [2]int{0,0}
	tailVisited[tail] = struct{}{}
	
	for _, ins := range strings.Split(puzzle, "\n") {
		parts := strings.Split(ins, " ")
		d := parts[0]
		n, _ := strconv.Atoi(parts[1])
		
		for j := 0; j < n; j++ {
			last[0], last[1] = head[0], head[1]
			head[0] += dirMap[d][0]
			head[1] += dirMap[d][1]

			if math.Abs(float64(head[0] - tail[0])) > 1.5 || math.Abs(float64(head[1] - tail[1])) > 1.5 {
				tail[0], tail[1] = last[0], last[1]
				tailVisited[tail] = struct{}{}
			}
			// fmt.Printf("instruction: %s head (%d, %d) | tail (%d, %d)\n", ins, head[0], head[1], tail[0], tail[1])
		}
	}
	fmt.Printf("2022 Day 9; Part 1: head (%d, %d) | tail (%d, %d) | tailVisited (%d)\n", head[0], head[1], tail[0], tail[1], len(tailVisited))

	//part 2
	tailVisited = map[[2]int]struct{}{}

	knots := make([]knot, 10) // A slice to hold pointers to all 10 knots

	for k := 0; k < 10; k++ {
		knots[k] = knot{coor: [2]int{0, 0}, last: [2]int{0, 0}}
	}

	for _, ins := range strings.Split(puzzle, "\n") {
		parts := strings.Split(ins, " ")
		//fmt.Printf("instruction: %s\n", ins)
		d := parts[0]
		n, _ := strconv.Atoi(parts[1])
		
		for j := 0; j < n; j++ {
			// fmt.Println("iteration", j + 1)
			for i := 0; i < 10; i++ {
				if i == 0 {
					knots[i].last[0], knots[i].last[1] = knots[i].coor[0], knots[i].coor[1]
					knots[i].coor[0] += dirMap[d][0]
					knots[i].coor[1] += dirMap[d][1]
					
				} else {
					prevCoor := knots[i - 1].coor
					if math.Abs(float64(knots[i].coor[0] - prevCoor[0])) > 1.5 || math.Abs(float64(knots[i].coor[1] - prevCoor[1])) > 1.5 {
						// set last to current coor
						knots[i].last[0], knots[i].last[1] = knots[i].coor[0], knots[i].coor[1]
						// get dir based on last knot minus current (max +1, -1)
						dir := [2]int{0, 0}
						switch comp := knots[i - 1].coor[0] - knots[i].coor[0]; {
						case comp > 0:
							dir[0] = 1
						case comp < 0:
							dir[0] = -1
						}
						switch comp := knots[i - 1].coor[1] - knots[i].coor[1]; {
						case comp > 0:
							dir[1] = 1
						case comp < 0:
							dir[1] = -1
						}
						// move current knot to curKnotCoor + dir
						knots[i].coor[0], knots[i].coor[1] = knots[i].coor[0] + dir[0], knots[i].coor[1] + dir[1]
					}
					
					if i == 9 {
						tailVisited[knots[i].coor] = struct{}{}
					}
				}
				// fmt.Printf("knot[%d]: (%d, %d) => (%d, %d)\n", i, knots[i].last[0], knots[i].last[1], knots[i].coor[0], knots[i].coor[1])
			}
		}
	}
	endTS := time.Now()
	fmt.Printf("2022 Day 9; Part 2: head (%d, %d) | tailVisited (%d) | time: %s\n", knots[0].coor[0], knots[0].coor[1], len(tailVisited), endTS.Sub(startTS))
}