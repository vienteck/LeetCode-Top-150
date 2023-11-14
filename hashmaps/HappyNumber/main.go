package main

import "fmt"

func main() {
	isHappy(2)
}

func isHappy(n int) bool {
	newNumber := 0
	for newNumber != 1 {
		if newNumber == 4 {
			return false
		}
		newNumber = 0
		for n > 0 {
			k := n % 10
			fmt.Println(k)
			n = n / 10
			newNumber += k * k
		}
		n = newNumber
	}

	return true
}
