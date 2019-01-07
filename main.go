package main

import (
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	determine(argsWithoutProg)
}