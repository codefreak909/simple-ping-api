package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("not divisible by zero")
	}

	return x / y, nil
}

func main() {
	output, err := divide(5, 0)
	fmt.Println(output, err)
}
