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

func theHaNoiTower(num int, src string, dst string, tmp string, totalTime *int64, steps *int) {
	// The Tower of Hanoi (also called the Tower of Brahma) We are given three rods and N number
	// of disks, initially all the disks are added to first rod (the leftmost one) is in decreasing size order.
	// The objective is to transfer the entire stack of disks from first tower to third tower (the rightmost
	// one), moving only one disk at a time and never a larger one onto a smaller.

	// Analysis: If we want to move N disks from source to destination, then we first move N-1 disks
	// from source to temp, then move the lowest Nth disk from source to destination. Then it will
	// move N-1 disks from temp to destination.
	start := int64(time.Now().Nanosecond())
	if num < 1 {
		fmt.Println("The run time of the algorithm:", *steps, " steps")
		fmt.Println("The run time of the algorithm:", *totalTime, " nanoseconds")
		return
	}
	theHaNoiTower(num - 1, src, tmp, dst, totalTime, steps)
	fmt.Println("\n Move [", num, "] disk from peg ", src, " to peg ", dst)
	*steps ++
	theHaNoiTower(num - 1, tmp, dst, src, totalTime, steps)
	end := int64(time.Now().Nanosecond())
	*totalTime += (end - start)	
}

func main () {
	fmt.Println("[Headfirst Series][Game10]: Algorithms Analysis")
	theLinearTime()
	num := 8
	var totalTime int64 = 0
	var steps int = 0
	theHaNoiTower(num, "A", "C", "B", &totalTime, &steps)
}