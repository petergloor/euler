// project problem_005 main.go
// Solution for Project Euler Problem #005, Smallest multiple.
package main

import "fmt"

func main() {
	var k int
	for i := 1; true; i++ {
		// Let the user know we are still in search of the solution
		if i%1000000 == 0 {
			fmt.Println(i/1000000, " million numbers processed")
		}
		k = 0
		for j := 1; j <= 20; j++ {
			if i%j == 0 {
				k++
			}
		}
		if k == 20 {
			fmt.Println("Hoorrray, here it comes:", i)
			fmt.Println("Actually, Go supports concurrency by design and could do that much faster.")
			break
		}
	}
}
