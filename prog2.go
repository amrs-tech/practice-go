// Program to generate range of numbers

package main

import "fmt"

func main() {
	var x = 0
	for x < 5 {
		fmt.Println("Checking incr loop of x = %v value",x)
		x+=2
	}
}