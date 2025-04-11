package main

import (
	"container/heap"
	"dijkstras-alg-go/model"
	"dijkstras-alg-go/priority-queue"
	"fmt"
)

const INFINITY int = 999

var EDGES = [...]model.NodeID{
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
}

func main() {
	graphMap, distanceMap, visitedMap, predeccessorMap := initGraph()

	pq := make(priority_queue.PriorityQueue, 0)
	heap.Init(&pq)

	var sourceNode model.NodeID = "u"
	distanceMap[sourceNode] = 0
	fmt.Printf("Source is %s\n", sourceNode)

	heap.Push(&pq, &priority_queue.Item{
		Value: sourceNode,
	})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*priority_queue.Item)

		if visitedMap[current.Value] {
			continue
		}

		visitedMap[current.Value] = true

		for neighborID, neighborWeight := range graphMap[current.Value] {
			newDistance := distanceMap[current.Value] + neighborWeight

			if newDistance < distanceMap[neighborID] {
				distanceMap[neighborID] = newDistance
				predeccessorMap[neighborID] = current.Value
				heap.Push(&pq, &priority_queue.Item{
					Value:    neighborID,
					Priority: newDistance,
				})
			}
		}
	}

	fmt.Println("\nShortest distances from source:")
	for _, node := range EDGES {
		fmt.Printf("Distance to %s: %d\n", node, distanceMap[node])
	}

	fmt.Println("\nShortest paths from source:")
	for _, node := range EDGES {
		if node == sourceNode {
			continue // Skip the source node
		}

		// Reconstruct the path
		path := []model.NodeID{node}
		current := node
		for current != sourceNode {
			current = predeccessorMap[current]
			path = append([]model.NodeID{current}, path...)
		}

		// Print the path
		fmt.Printf("Path to %s: ", node)
		for i, pathNode := range path {
			if i > 0 {
				fmt.Print(" -> ")
			}
			fmt.Print(pathNode)
		}
		fmt.Println()
	}

}

func initGraph() (map[model.NodeID]map[model.NodeID]int, map[model.NodeID]int, map[model.NodeID]bool, map[model.NodeID]model.NodeID) {
	graphMap := make(map[model.NodeID]map[model.NodeID]int)
	addGraphEdge(graphMap, "u", "v", 1)
	addGraphEdge(graphMap, "u", "w", 3)
	addGraphEdge(graphMap, "u", "x", 5)

	addGraphEdge(graphMap, "v", "u", 1)
	addGraphEdge(graphMap, "v", "w", 3)
	addGraphEdge(graphMap, "v", "y", 4)

	addGraphEdge(graphMap, "y", "x", 7)
	addGraphEdge(graphMap, "y", "w", 12)
	addGraphEdge(graphMap, "y", "v", 4)
	addGraphEdge(graphMap, "y", "z", 2)

	addGraphEdge(graphMap, "z", "y", 2)
	addGraphEdge(graphMap, "z", "x", 3)

	addGraphEdge(graphMap, "x", "u", 5)
	addGraphEdge(graphMap, "x", "w", 1)
	addGraphEdge(graphMap, "x", "y", 7)
	addGraphEdge(graphMap, "x", "z", 3)

	distanceMap := make(map[model.NodeID]int)
	for _, edge := range EDGES {
		distanceMap[edge] = INFINITY
	}

	visitedMap := make(map[model.NodeID]bool)
	for _, edge := range EDGES {
		visitedMap[edge] = false
	}

	predecessorMap := make(map[model.NodeID]model.NodeID)
	for _, edge := range EDGES {
		predecessorMap[edge] = ""
	}

	return graphMap, distanceMap, visitedMap, predecessorMap
}

func addGraphEdge(graph map[model.NodeID]map[model.NodeID]int, from model.NodeID, to model.NodeID, weight int) {
	if _, exists := graph[from]; !exists {
		graph[from] = make(map[model.NodeID]int)
	}
	graph[from][to] = weight
}
