package main

import (
	"github.com/landanqrew/goAdvent/osUtils"
	solutions2023 "github.com/landanqrew/goAdvent/solutions/2023"
)

func main() {
	dir := osUtils.ListDirectoryContents(".", true)
	if false {
		dir.Print()
	} 

	// day 1
	solutions2023.Day1()
	solutions2023.Day2(false, false)
	solutions2023.Day3(false, true)
}
