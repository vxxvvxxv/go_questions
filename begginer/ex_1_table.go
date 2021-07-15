package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Print("   ", i*j) // or fmt.Print("\t", i*j)
		}
		fmt.Println()
	}
}
