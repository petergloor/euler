// project problem_003 main.go
// Solution for Project Euler Problem #003, Largest prime factor
package main

import "fmt"

// main is the entry point of the app
func main() {
	n := 600851475143
	fmt.Println("Find largest primefactor of ", n)
	for i := 2; i < (n / i); i++ {
		for n%i == 0 {
			n = n / i
			fmt.Printf("%d x %d\n", n, i)
		}
	}
	if n > 1 {
		fmt.Println("Largest prime factor:", n)
	}

}

/*
 Largest prime factor
 Problem 3
 The prime factors of 13195 are 5, 7, 13 and 29.
 What is the largest prime factor of the number 600851475143 ?
*/
