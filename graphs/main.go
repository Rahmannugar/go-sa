package main

import "fmt"

// AdjacencyMatrix represents an undirected graph with a matrix.
type AdjacencyMatrix struct {
	matrix [][]bool
}

// NewAdjacencyMatrix creates a graph with the given number of vertices.
func NewAdjacencyMatrix(numVertices int) *AdjacencyMatrix {
	matrix := make([][]bool, numVertices)

	for i := range matrix {
		matrix[i] = make([]bool, numVertices)
	}

	return &AdjacencyMatrix{
		matrix: matrix,
	}
}

// AddEdge connects two vertices in the matrix.
// Time complexity: O(1), Space complexity: O(1)
func (g *AdjacencyMatrix) AddEdge(u, v int) {
	g.matrix[u][v] = true
	g.matrix[v][u] = true
}

// EdgeExists reports whether two vertices are connected in the matrix.
// Time complexity: O(1), Space complexity: O(1)
func (g *AdjacencyMatrix) EdgeExists(u, v int) bool {
	return g.matrix[u][v]
}

// AdjacencyList represents an undirected graph with neighbor lists.
type AdjacencyList struct {
	graph map[int][]int
}

// NewAdjacencyList creates an empty graph.
func NewAdjacencyList() *AdjacencyList {
	return &AdjacencyList{
		graph: make(map[int][]int),
	}
}

// AddVertex adds a vertex if it is not already present.
func (g *AdjacencyList) AddVertex(vertex int) {
	if _, exists := g.graph[vertex]; !exists {
		g.graph[vertex] = []int{}
	}
}

// AddEdge connects two vertices in the adjacency list.
// Time complexity: O(1) average amortized, Space complexity: O(1)
func (g *AdjacencyList) AddEdge(u, v int) {
	g.AddVertex(u)
	g.AddVertex(v)

	g.graph[u] = append(g.graph[u], v)
	g.graph[v] = append(g.graph[v], u)
}

// EdgeExists checks a vertex's neighbors for an edge to another vertex.
// Time complexity: O(degree(u)), Space complexity: O(1)
func (g *AdjacencyList) EdgeExists(u, v int) bool {
	for _, neighbor := range g.graph[u] {
		if neighbor == v {
			return true
		}
	}

	return false
}

func main() {
	matrix := NewAdjacencyMatrix(5)
	list := NewAdjacencyList()

	edges := [][2]int{
		{0, 1},
		{0, 4},
		{1, 2},
		{1, 3},
		{1, 4},
	}

	for _, edge := range edges {
		matrix.AddEdge(edge[0], edge[1])
		list.AddEdge(edge[0], edge[1])
	}

	fmt.Println("Adjacency Matrix:")

	for _, row := range matrix.matrix {
		fmt.Println(row)
	}

	fmt.Println("\nAdjacency List:")

	for vertex, neighbors := range list.graph {
		fmt.Printf("%d: %v\n", vertex, neighbors)
	}

	fmt.Println("\nEdge checks:")

	fmt.Println("Matrix: 2 connected to 3?", matrix.EdgeExists(2, 3))
	fmt.Println("List: 2 connected to 3?", list.EdgeExists(2, 3))

	fmt.Println("Matrix: 1 connected to 4?", matrix.EdgeExists(1, 4))
	fmt.Println("List: 1 connected to 4?", list.EdgeExists(1, 4))
}
