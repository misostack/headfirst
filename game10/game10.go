package main

import (
	"fmt"
	"time"
)

func theLinearTime() {
	start := time.Now().Nanosecond()
	for i:=0; i<10000; i++ {
		// do nothing
	}
	end := time.Now().Nanosecond()

	fmt.Println("The run time of the algorithm:", end - start, " miliseconds")
}

func theHaNoiTower(num int, src string, dst string, tmp string) {
	// The Tower of Hanoi (also called the Tower of Brahma) We are given three rods and N number
	// of disks, initially all the disks are added to first rod (the leftmost one) is in decreasing size order.
	// The objective is to transfer the entire stack of disks from first tower to third tower (the rightmost
	// one), moving only one disk at a time and never a larger one onto a smaller.

	// Analysis: If we want to move N disks from source to destination, then we first move N-1 disks
	// from source to temp, then move the lowest Nth disk from source to destination. Then it will
	// move N-1 disks from temp to destination.
	
	if num < 1 {
		return
	}
	theHaNoiTower(num - 1, src, tmp, dst)
	fmt.Println("\n Move [", num, "] disk from peg ", src, " to peg ", dst)
	theHaNoiTower(num - 1, tmp, dst, src)
}

func main () {
	fmt.Println("[Headfirst Series][Game10]: Algorithms Analysis")
	theLinearTime()
	num := 3
	theHaNoiTower(num, "A", "C", "B")
}