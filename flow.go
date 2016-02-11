package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := []int{0, 0, 0, 0, 0, 0}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000000; i++ {
		roll := rand.Intn(6) + 1
		slice[roll-1] += 1
	}

	fmt.Println("1 :", slice[0])
	fmt.Println("2 :", slice[1])
	fmt.Println("3 :", slice[2])
	fmt.Println("4 :", slice[3])
	fmt.Println("5 :", slice[4])
	fmt.Println("6 :", slice[5])
}
