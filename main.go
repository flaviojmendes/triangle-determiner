package main

import (
	"log"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	determine(argsWithoutProg)

	log.Printf("Hello World!")
}