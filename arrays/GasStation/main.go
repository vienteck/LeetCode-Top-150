package main

import "fmt"

func main() {
	answer := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	fmt.Println(answer)
}

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost, currentGas, start := 0, 0, 0, 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentGas += gas[i] - cost[i]

		// If not enough gas to reach the next station, start from the next station
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}

	// If totalGas is less than totalCost, it's impossible to complete the circuit
	if totalGas < totalCost {
		return -1
	}

	return start
}
