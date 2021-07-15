package main

import "fmt"

func main() {
	x := []int{0, 0, -4, -2, 4, 0, 2, -4, 3, 0, 0, -4, 2, 2, 0, -3, -3, 3, -3, -4}

	plus := 0
	minus := 0
	zero := 0

	for i := 0; i < len(x); i++ {
		if x[i] < 0 {
			minus += 1
		} else if x[i] > 0 {
			plus += 1
		} else {
			zero += 1
		}
	}

	fmt.Println("Положительных: ", plus)
	fmt.Println("Отрицательных: ", minus)
	fmt.Println("Равных нулю: ", zero)

	y := []int{-4, 2, 2, 0, -3, -3, 3, -3, -4, 3, 2, -3, 2, -1, 3, 4, -4, 3, 0, 0}

	plus = 0
	minus = 0
	zero = 0

	for i := 0; i < len(y); i++ {
		if y[i] < 0 {
			minus += 1
		} else if y[i] > 0 {
			plus += 1
		} else {
			zero += 1
		}
	}

	fmt.Println("Положительных: ", plus)
	fmt.Println("Отрицательных: ", minus)
	fmt.Println("Равных нулю: ", zero)
}
