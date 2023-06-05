// Program to iterate over slice/array

package main

import "fmt"

func main() {
	var x = 0
	var sampleSlice = []string{"cap","iron","spidey"}
	for x < len(sampleSlice) {
		fmt.Printf("Checking iter elem = %v value\n",sampleSlice[x])
		x++
	}
}