/*
	Exercise 1.2:
	Modify the echo program to to print the index and value of each of its arguments, one per line.
 */

package main

import (
	"os"
	"fmt"
)

func main()  {
	for i, arg := range os.Args {
		fmt.Printf("Args[%d] = %v\n", i, arg)
	}
}