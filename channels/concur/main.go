package main

import (
	"fmt"
)

func calcCubes(num int, cuchan chan int) {
	sq := 0;
	for num!= 0 {
		digit := num%10
		num = num/10
		sq = sq + digit * digit * digit
	}
	cuchan <- sq
}

func calcSqares(num int, sqchan  chan int) {
	sq := 0;
	for num!= 0 {
		digit := num%10
		num = num/10
		sq = sq + digit * digit
	}
	sqchan <- sq
}

func main() {
	num := 1001
	sqchan := make(chan int)
	cuchan := make(chan int)
	go calcSqares(num, sqchan)
	go calcCubes(num, cuchan)	
	squares, cubes:= <- sqchan, <- cuchan
	fmt.Println("Final vales", squares+cubes)
}

