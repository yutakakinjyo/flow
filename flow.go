package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	stock := []int{0, 0, 0, 0, 0}
	output := 0

	for i := 0; i < 100; i++ {

		stock[0] += roll()

		flow := max(stock[0], roll())
		stock[0] -= flow
		stock[1] += flow

		flow = max(stock[1], roll())
		stock[1] -= flow
		stock[2] += flow

		flow = max(stock[2], roll())
		stock[2] -= flow
		stock[3] += flow

		flow = max(stock[3], roll())
		stock[3] -= flow
		stock[4] += flow

		flow = max(stock[4], roll())
		stock[4] -= flow
		output += flow
	}

	fmt.Println("==== Result =====")
	fmt.Println("stock[0] :", stock[0])
	fmt.Println("stock[1] :", stock[1])
	fmt.Println("stock[2] :", stock[2])
	fmt.Println("stock[3] :", stock[3])
	fmt.Println("stock[4] :", stock[4])
	fmt.Println("output :", output)

}
