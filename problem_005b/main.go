// project problem_005b main.go
// Solution for Project Euler Problem #005, Smallest multiple.
// TODO: Try to make it faster using concurrency with Go routines.
/*
 W O R K  I N  P R O G R E S S  -  D O N T  U S E  T H I S  C O D E
 */
package main

import (
	"fmt"
	"runtime"
)

func find(max int) {
	var k int
	for i := 1; true; i++ {
		// Let the user know we are still in search of the solution
		if i%1000000 == 0 {
			fmt.Println(i/1000000, " million numbers processed")
		}
		k = 0
		for j := 1; j <= max; j++ {
			if i%j == 0 {
				k++
			}
		}
		if k == max {
			fmt.Println("Hoorrray, here it comes:", i)
			fmt.Println("Actually, Go supports concurrency by design and could do that much faster.")
			break
		}
	}
}

func main() {
	max := 10
	maxInt, numCPUs := hardwareCheck() // Set hardware specific values
	nGoRoutines := numCPUs             // Calculate and Set the # of Go routines
	fmt.Println("# Go routines:", nGoRoutines)

	c := make(chan int)

	fmt.Println(maxInt / 256*1024)

	// Reservoir for unused variables to avoid compiler errors.
	// Should be empty by the end of the day.
	_, _, _ = c, maxInt, max

}

// hardwareCheck returns the highest integer value (maxInt)
// and the number of CPU's on this computer.
func hardwareCheck() (int, int) {
	return int(^uint(0) >> 1), runtime.NumCPU()
}

// preserve keeps parts of the main() code used in problem_005a
func preserve(max int) {
	var k int
	for i := 1; true; i++ {
		// Let the user know we are still in search of the solution
		if i%1000000 == 0 {
			fmt.Println(i/1000000, " million numbers processed")
		}
		k = 0
		for j := 1; j <= max; j++ {
			if i%j == 0 {
				k++
			}
		}
		if k == max {
			fmt.Println("Hoorrray, here it comes:", i)
			fmt.Println("Actually, Go supports concurrency by design and could do that much faster.")
			break
		}
	}

}

/*
 Smallest multiple
 Problem 5
 2520 is the smallest number that can be divided by each of the numbers
 from 1 to 10 without any remainder.

 What is the smallest positive number that is evenly divisible by all of
 the numbers from 1 to 20?
*/
