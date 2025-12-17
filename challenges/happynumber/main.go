package main

import (
	"fmt"
	"strconv"
	"strings"
)

// A happy number is defined by the following process: starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle that does not include 1. Those numbers for which this process ends in 1 are happy numbers.
//
// Write a function that determines if a number is happy.

// input: 19 or 4
// output: true for 19, false for 4
// Explanation:
// 1. divide the number into its digits 1 and 9, 4
// 2. square each digit and sumt them up: 1^2 + 9^2 = 82, 4^2 = 16
// 3. repeat the process until the number equal 1 or loops ends
// 4. if the number equal 1 return true, else return false
// 4. to detect loops deep or if the number is already seen.
// Example:
// 0 -> not happy
// 1 -> happy
// 2 -> not happy yet -> 2^2 = 4 -> not happy yet -> 4^2 = 16 -> not happy yet
// 16 -> 1^2 + 6^2 = 37 -> not happy yet
// 123 -> 1^2 + 2^2 + 3^2 = 14 -> not happy yet
// Approach:
// use a set to track seen numbers
// split the number into its digits
// if number is 1 return true
// if number is seen return false
// square each digit and sum them up
// add to number to the seen numbers
// repeat again

func main() {
	fmt.Println("Calculating happy numbers")
	number := 123456789 // 0, 1, 2, 4, 19, 123456789
	happy := isHappy(number)
	fmt.Printf("%d is happy: %v\n", number, happy)
}

func isHappy(number int) bool {
	seen := map[int]struct{}{}

	counter := 0
	for {
		fmt.Printf("operation %d: %d\n", counter, number)
		if number == 0 {
			return false
		}

		if number == 1 {
			return true
		}

		digits := strings.Split(strconv.Itoa(number), "") // 19
		var sum int
		for _, digit := range digits { // 1, 9
			numericalDigit, _ := strconv.Atoi(digit)
			sum = sum + (numericalDigit * numericalDigit) // 0 + 1*1 = 1 + 9*9 = 82
		}

		number = sum

		if _, ok := seen[number]; ok {
			return false
		}
		seen[number] = struct{}{}

		counter++
	}
}
