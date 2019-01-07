package main

import (
	"log"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	triangleType := determine(argsWithoutProg)
	log.Printf("The type of the Triangle is %s", triangleType)
}