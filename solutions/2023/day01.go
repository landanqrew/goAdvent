package solutions2023

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/landanqrew/goAdvent/osUtils"
)

var stringMap = map[string]string{
	"one":   "on1ne",
	"two":   "tw2so",
	"three": "thr3ree",
	"four":  "fo4ur",
	"five":  "fi5ve",
	"six":   "si6ix",
	"seven": "sev7ven",
	"eight": "eig8ght",
	"nine":  "ni9ne",
	"zero":  "ze0ro",
}

/*var example string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`*/

func Day1() {
	puzzle := string(osUtils.ReadFileFromPath("puzzles/2023/day1.txt"))
	//puzzle = example
	lines := strings.Split(puzzle, "\n")

	part1, part2 := 0, 0

	for _, line := range lines {
		// fmt.Println(line)
		transformedLine := line
		for num, str := range stringMap {
			transformedLine = strings.ReplaceAll(transformedLine, num, str)
		}
		p1Nums := extractNumberStrings(line)
		if len(p1Nums) >= 1 {
			part1 += convertNumberStringToNumber(p1Nums[0] + p1Nums[len(p1Nums)-1])
		}

		p2Nums := extractNumberStrings(transformedLine)
		if len(p2Nums) >= 1 {
			part2 += convertNumberStringToNumber(p2Nums[0] + p2Nums[len(p2Nums)-1])
		}
		// fmt.Println(part1, part2)
	}
	fmt.Println(part1, part2)
}

func extractNumberStrings(str string) []string {
	//nums := []string{}
	rePattern, err := regexp.Compile(`\d`)
	if err != nil {
		log.Fatal("cannot compile regex pattern", err)
	}
	nums := rePattern.FindAllString(str, -1)
	// matches := rePattern.FindAllString(str, -1)
	/*for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			log.Printf("cannot convert string '%s' to int: %v", match, err)
		}
		nums = append(nums, num)
	}*/
	return nums
}

func convertNumberStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("cannot convert string '%s' to int: %v", str, err)
	}
	return num
}