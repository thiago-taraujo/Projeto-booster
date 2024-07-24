package main

import (
	"fmt"
	"time"
)

func numeroAoQuadrado() {
	for num := 1; num <= 100; num ++ {
		fmt.Printf("%d\n", num*num)
	}
}

func main() {
	start := time.Now()
	go numeroAoQuadrado()
	time.Sleep(time.Second)
	end := time.Now()
	fmt.Printf("duration: %s\n", end.Sub(start))
}