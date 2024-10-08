/*
Instructions
Write a function that takes an int min and an int max as parameters.
The function should return a slice of ints with all the values between
the min and max.

Min is included, and max is excluded.

If min is greater than or equal to max, a nil slice is returned.

make is not allowed for this exercise.

Expected function
func AppendRange(min, max int) []int {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(5, 10))
	fmt.Println(piscine.AppendRange(10, 5))
}
And its output :

$ go run .
[5 6 7 8 9]
[]
$
*/
package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	} else {

		var x []int

		for i := min; i < max; i++ {
			x = append(x, i)
		}

		return x
	}
}
