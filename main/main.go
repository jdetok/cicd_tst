package main

import (
	"fmt"
	"log"
)

func add_two_pos_nums(n1, n2 int) (int, error) {
	if n1 <= 0 || n2 <= 0 {
		return 0, fmt.Errorf("both numbers greater than or equal to 0: %d | %d", n1, n2)
	}
	return (n1 + n2), nil
}

func main() {
	res, err := add_two_pos_nums(5, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result: ", res)
}
