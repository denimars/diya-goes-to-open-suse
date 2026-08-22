package main

import (
	"fmt"
)

func UrutanDeploy(jumlahModul int, ketergantungan [][]int) []int {
	queue := make([]int, 0)
	state := make([]int, jumlahModul)
	moduleDependecies := make([][]int, jumlahModul)
	hasCycle := false
	for _, d := range ketergantungan {
		moduleDependecies[d[0]] = append(moduleDependecies[d[0]], d[1])
	}
	for i := range jumlahModul {
		if state[i] == 0 {
			explore(&queue, i, &moduleDependecies, &state, &hasCycle)
		}
	}
	fmt.Println("has cycle", hasCycle)
	if hasCycle {
		return []int{}
	}

	return queue

}

func explore(queue *[]int, node int, moduleDependenciesP *[][]int, statusP *[]int, hasCycle *bool) {
	status := *statusP
	if *hasCycle {
		return
	}
	if status[node] == 1 {
		*hasCycle = true
		return

	}
	if status[node] == 2 {
		return
	}
	status[node] = 1
	moduleDependencies := *moduleDependenciesP
	for _, p := range moduleDependencies[node] {

		explore(queue, p, moduleDependenciesP, &status, hasCycle)
	}
	status[node] = 2
	*queue = append(*queue, node)

}

func main() {
	dependecies := [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		// {4, 5},
		// {3, 1},
		// {3, 5}}
	}
	fmt.Println(UrutanDeploy(4, dependecies))
}
