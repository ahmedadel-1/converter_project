package graph

import (
	"container/heap"
	"math"
)

// Edge represents a conversion with a quality (weight).
type Edge struct {
	To, Weight int
}

// Graph represents the graph structure.
type Graph struct {
	Nodes map[int][]Edge
}

// NewConversionGraph initializes a new conversion graph.
func NewConversionGraph() *Graph {
	return &Graph{Nodes: make(map[int][]Edge)}
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(from, to, weight int) {
	g.Nodes[from] = append(g.Nodes[from], Edge{To: to, Weight: weight})
}

// PriorityQueue implements a min-heap for Dijkstra's algorithm.
type PriorityQueue []Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Distance < pq[j].Distance }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Item represents a node in the priority queue.
type Item struct {
	Node, Distance int
}

// Dijkstra finds the shortest paths from a source node.
func Dijkstra(g *Graph, source int) map[int]int {
	distances := make(map[int]int)
	for node := range g.Nodes {
		distances[node] = math.MaxInt32
	}
	distances[source] = 0

	pq := &PriorityQueue{{Node: source, Distance: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)
		currentNode := current.Node
		currentDistance := current.Distance

		if currentDistance > distances[currentNode] {
			continue
		}

		for _, edge := range g.Nodes[currentNode] {
			newDistance := currentDistance + edge.Weight
			if newDistance < distances[edge.To] {
				distances[edge.To] = newDistance
				heap.Push(pq, Item{Node: edge.To, Distance: newDistance})
			}
		}
	}

	return distances
}
