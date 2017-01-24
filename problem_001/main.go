// project problem_001 main.go
// Solution for Project Euler Problem #101, Multiples of 3 and 5.
package main

import "fmt"

func main() {
	var sum int

	for i := 1; i < 1000; i++ {
		switch {
		case i%3 == 0:
			sum += i
			fmt.Println(i, sum)

		case i%5 == 0:
			sum += i
			fmt.Println(i, sum)
		}
	}
	fmt.Println("The sum is:", sum)

}

/*
 Multiples of 3 and 5
 Problem 1
 If we list all the natural numbers below 10 that are multiples of 3 or 5, we
 get 3, 5, 6 and 9. The sum of these multiples is 23.

 Find the sum of all the multiples of 3 or 5 below 1000.
*/
