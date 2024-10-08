/*
Instructions
Write a function that takes an int min and an int max as parameters. The function must return a slice of ints with all the values between min and max.

Min is included, and max is excluded.

If min is greater than or equal to max, a nil slice is returned.

append is not allowed for this exercise.

Expected function
func MakeRange(min, max int) []int {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.MakeRange(5, 10))
	fmt.Println(piscine.MakeRange(10, 5))
}
And its output :

$ go run .
[5 6 7 8 9]
[]
$
*/
package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	} else {
		x := make([]int, max-min)
		for v := range x {
			x[v] = min
			min = min + 1
		}
		return x
	}
}
