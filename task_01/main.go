package main

import "fmt"

func main() {
	s := generateSlice(10)

	for _, v := range s {
		if isEven(v) {
			fmt.Printf("%v is even \n", v)
		} else {
			fmt.Printf("%v is odd \n", v)
		}
	}
}

func generateSlice(size int) []int {
	s := []int{0}
	for i := 0; i < size; i++ {
		s = append(s, s[i]+1)
	}
	return s
}

func isEven(number int) bool {
	return (number % 2) == 0
}
