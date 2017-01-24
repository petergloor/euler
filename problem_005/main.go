// project problem_005 main.go
// Solution for Project Euler Problem #005, Smallest multiple.
package main

import "fmt"

func main() {
	const max = 20
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
