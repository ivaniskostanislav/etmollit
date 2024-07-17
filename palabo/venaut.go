package main

import (
	"fmt"
	"strconv"
)

func octalToBinary(octal string) (string, error) {
	decimal, err := strconv.ParseInt(octal, 8, 64)
	if err != nil {
		return "", err
	}
	binary := strconv.FormatInt(decimal, 2)
	return binary, nil
}

func main() {
	octal := "17" // Example octal number
	binary, err := octalToBinary(octal)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(binary) // Output: 1111
}
