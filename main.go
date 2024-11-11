package main

import (
	"converter_project/graph"
	"fmt"
)

func main() {
	// Initialize conversion graph and add conversion edges with weights
	conversionGraph := graph.NewConversionGraph()
	graph.SetupConversions(conversionGraph)

	// Source and target formats
	const (
		Excel = 1
		Word  = 2
		PDF   = 3
		HTML  = 4
	)

	sourceFormat := HTML
	distances := graph.Dijkstra(conversionGraph, sourceFormat)

	// Display shortest paths
	fmt.Printf("Shortest paths (lowest-cost conversions) from format %d:\n", sourceFormat)
	for format, distance := range distances {
		fmt.Printf("To format %d: Cost %d\n", format, distance)
	}
}
