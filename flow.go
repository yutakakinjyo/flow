package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	stock int
}

func max(stock, roll int) int {
	if stock < roll {
		roll = stock
	}
	return roll
}

func roll() int {
	return rand.Intn(6) + 1
}

func main() {

	rand.Seed(time.Now().UnixNano())
	worker := []Worker{
		Worker{0},
		Worker{0},
		Worker{0},
		Worker{0},
		Worker{0},
	}
	output := 0

	for i := 0; i < 100; i++ {

		worker[0].stock += roll()

		flow := max(worker[0].stock, roll())
		worker[0].stock -= flow
		worker[1].stock += flow

		flow = max(worker[1].stock, roll())
		worker[1].stock -= flow
		worker[2].stock += flow

		flow = max(worker[2].stock, roll())
		worker[2].stock -= flow
		worker[3].stock += flow

		flow = max(worker[3].stock, roll())
		worker[3].stock -= flow
		worker[4].stock += flow

		flow = max(worker[4].stock, roll())
		worker[4].stock -= flow
		output += flow
	}

	fmt.Println("==== Result =====")
	fmt.Println("worker[0] :", worker[0].stock)
	fmt.Println("worker[1] :", worker[1].stock)
	fmt.Println("worker[2] :", worker[2].stock)
	fmt.Println("worker[3] :", worker[3].stock)
	fmt.Println("worker[4] :", worker[4].stock)
	fmt.Println("output :", output)

}
