// project clocksequence main.go
// Solution for Project Euler Problem #506
/*
  Currently it makes not much sense to try to solve the problem
  as it obviousely needs something like the math/big package.
  The current solution works for sequences up to S(49) only.
  However, I'll keep this for further investigations.
*/
package main

import (
	"fmt"
)

// main is the entry point of the clocksequence app
func main() {
	n := 49       // n in S(n)
	var stub int  // stub of digits
	var total int // sum of all stubs

	p := 1 // Next position in the sequence of digits
	for i := 1; i <= n; i++ {
		stub, p = getNextStub(i, p)
		fmt.Println(i, stub)
		total += stub
	}
	fmt.Printf("The resulting total of S(%d) is %d\n", n, total)
	fmt.Printf("The resulting total of S(%d) mod 123454321 is %d\n",
		n, total%123454321)

}

// getNextStub returns a stub of digits from the sequence of digits.
// It starts at position p in the sequence of digits and returns the
// stub of digits that builds the sum of n. It also returns the next
// position in the thought sequence of digits.
func getNextStub(n, p int) (int, int) {
	var sum int   // sum of digits in the stub
	var digit int // digit at current position
	var stub int  // stub to return

	for sum < n {
		digit = getDigit(p)
		p++
		sum += digit
		if sum < n {
			stub += digit
			stub = stub * 10
		} else {
			stub += digit
		}
	}

	return stub, p
}

// getDigit returns the value of a given position from a thought
// infinitive repeating sequence of digits 123432123432...
// The position is based on 1; e.g. getDigit(5) == 3.
func getDigit(p int) int {
	i := (p - 1) % 6
	return []int{1, 2, 3, 4, 3, 2}[i]
}

/*
 Clock sequence
 Problem 506
 Published on Sunday, 8th March 2015, 04:00 am;
 Solved by 522; Difficulty rating: 30%
 Consider the infinite repeating sequence of digits:
 1234321234321234321...

 Amazingly, you can break this sequence of digits into a sequence of
 integers such that the sum of the digits in the n'th value is n.

 The sequence goes as follows:
 1, 2, 3, 4, 32, 123, 43, 2123, 432, 1234, 32123, ...

 Let vn be the n'th value in this sequence.
 For example, v2 = 2, v5 = 32 and v11 = 32123.

 Let S(n) be v1 + v2 + ... + vn.
 For example, S(11) = 36120, and S(1000) mod 123454321 = 18232686.

 Find S(1014) mod 123454321.
*/
