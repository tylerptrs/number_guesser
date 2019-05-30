package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	max = 100
	min = 0
)

var entry string

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func test(entry string, g int) {
	fmt.Printf("max = %d min = %d\n", max, min)
	switch entry {
	case "<":
		max = g
		guess(min, max)
	case ">":
		min = g
		guess(min, max)
	case "=":
		fmt.Printf("Your number is %d\n", g)
	default:
		fmt.Printf("%s is not <, > or =.  Try again\n", entry)
		fmt.Scan(&entry)
		test(entry, g)
	}
}

func guess(min, max int) {
	g := randInt(min, max)
	fmt.Printf("Is you number <, >, or = %d?\n", g)
	fmt.Scan(&entry)
	test(entry, g)
}

func main() {
	println("I will ask you several questions to guess your number, 0-100")
	guess(0, 100)
}
