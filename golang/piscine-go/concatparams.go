/*Instructions
Write a function that takes the arguments received in parameters and
returns them as a string. The string is the result of all the arguments
concatenated with a newline (\n) between.

Expected function
func ConcatParams(args []string) string {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
}
And its output :

$ go run .
Hello
how
are
you?
$
*/
package piscine

func ConcatParams(args []string) string {
	var x string

	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			x += args[i]
		} else {
			x += args[i] + "\n"
		}
	}
	return x
}
