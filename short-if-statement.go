package main

import "fmt"

func main() {
	if x := 10; x%2 == 0 {
		fmt.Printf("%d is even\n", x)
	} else {
		fmt.Printf("%d is odd\n", x)
	}
}
