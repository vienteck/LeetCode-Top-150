package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	// Build the graph and calculate in-degrees
	for _, prerequisite := range prerequisites {
		course, prereq := prerequisite[0], prerequisite[1]
		graph[prereq] = append(graph[prereq], course)
		inDegree[course]++
	}

	// Initialize a queue with courses having in-degree of 0
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Perform topological sort
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		for _, nextCourse := range graph[course] {
			inDegree[nextCourse]--

			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// Check if all courses can be completed
	for _, degree := range inDegree {
		if degree > 0 {
			return false
		}
	}

	return true
}

func main() {
	numCourses := 4
	prerequisites := [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}

	result := canFinish(numCourses, prerequisites)
	fmt.Println("Can finish all courses?", result)
}
