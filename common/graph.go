package common

// Based On: https://www.geeksforgeeks.org/find-longest-path-directed-acyclic-graph/
import (
	"math"

	"github.com/golang-collections/collections/stack"
)

type AdjNode struct {
	value, weight int
}

type Graph struct {
	verticies_count int
	nodes           [][]AdjNode
}

func CreateGraph(size int) Graph {
	return Graph{
		verticies_count: size,
		nodes:           make([][]AdjNode, size),
	}
}

func (g *Graph) Size() int {
	return g.verticies_count
}
func (g *Graph) AddEdge(u, value, weight int) {
	g.nodes[u] = append(g.nodes[u], AdjNode{value, weight})
}

func (g *Graph) topologicalSortUtil(v int, visited []bool,
	stack *stack.Stack) {
	// Mark the current node as visited
	visited[v] = true

	// Recur for all the vertices adjacent to this vertex
	for _, node := range g.nodes[v] {
		if !visited[node.value] {
			g.topologicalSortUtil(node.value, visited, stack)
		}
	}

	// Push current vertex to stack which stores topological
	// sort
	stack.Push(v)
}

func (g *Graph) LongestPath(starting, target int) int {

	stack := stack.New()
	dist := make([]int, g.verticies_count)

	// Mark all the vertices as not visited
	visited := make([]bool, g.verticies_count)
	for i := 0; i < g.verticies_count; i++ {
		visited[i] = false
	}

	// Call the recursive helper function to store Topological
	// Sort starting from all vertices one by one
	for i := 0; i < g.verticies_count; i++ {
		if !visited[i] {
			g.topologicalSortUtil(i, visited, stack)
		}
	}

	// Initialize distances to all vertices as infinite and
	// distance to source as 0
	for i := 0; i < g.verticies_count; i++ {
		dist[i] = math.MinInt
	}

	dist[starting] = 0
	// Process vertices in topological order
	for stack.Len() != 0 {
		// Get the next vertex from topological order
		u := stack.Pop().(int)

		// Update distances of all adjacent vertices
		if dist[u] != math.MinInt {
			for _, node := range g.nodes[u] {
				if dist[node.value] < dist[u]+node.weight {
					dist[node.value] = dist[u] + node.weight
				}
			}
		}
	}

	// Print the calculated longest distances
	return dist[target]
}
