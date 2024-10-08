/* Instructions
Write a function that returns the concatenation of all the strings of a slice of strings separated by the separator passed as the argument sep.

Expected function
func Join(strs []string, sep string) string {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
}
And its output :

$ go run .
Hello!: How: are: you?
$
*/
package piscine

func Join(strs []string, sep string) string {
	var asd string
	for i := 0; i < len(strs); i++ {
		if i < len(strs)-1 {
			asd += string(strs[i]) + sep
		} else {
			asd += string(strs[i])
		}
	}
	return asd
}
