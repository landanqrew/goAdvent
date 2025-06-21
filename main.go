package main

import (
	"github.com/landanqrew/goAdvent/osUtils"
)

func main() {
	osUtils.ListDirectoryContents(".", true).Print()
}
