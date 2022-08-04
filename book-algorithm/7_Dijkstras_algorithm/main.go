package main

import "math"

func main() {

}

func findShortestPath(graph map[string]map[string]int, startNode string, finishNode string) (map[string]int,
	map[string]string) {
	costs := make(map[string]int)
	costs[finishNode] = math.MaxInt32

	parents := make(map[string]string)
	parents[finishNode] = ""

	processed := make(map[string]bool)

	// Initialization of costs and parents

	for node, cost := range graph[startNode] {
		costs[node] = cost
		parents[node] = startNode
	}

	lowestCostNode := findLowestCostNode(costs, processed)
}
