package pkg

// PROBLEM STATEMENT:
//
// Given a list of trees (as pointers to their roots), return the biggest score
// of all nodes in all trees.
//
// SOLUTION:
//
// Traverse the tree keeping track of the biggest value in each tree. A helper
// function is used in order to recursively traverse the tree and keep track of
// visited nodes.

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

/**
 * Definition for a Node.
 */
type Node struct {
	id       uuid.UUID
	score    int
	children []*Node
}

func (n Node) String() string {
	return fmt.Sprintf("Node(score: %d)", n.score)
}

func NewNode() *Node {
	node := new(Node)
	node.id = uuid.New()
	node.score = rand.Intn(100)
	node.children = []*Node{}
	return node
}

func createTreeHelper(parent *Node, depth int) *Node {
	if depth <= 1 {
		return parent
	}
	for i := 1; i <= 3; i++ {
		parent.children = append(parent.children, createTree(depth-1))
	}
	return parent
}

func createTree(depth int) *Node {
	root := NewNode()
	return createTreeHelper(root, depth)
}

// Given multiple trees. Find the biggest node score.
func findBiggestScore(trees []*Node) int {
	biggest := 0
	for _, root := range trees {
		visited := make(map[*Node]bool)
		curr := findBiggestScoreHelper(root, visited, 0)
		if curr > biggest {
			biggest = curr
		}
	}
	return biggest
}

func findBiggestScoreHelper(node *Node, visited map[*Node]bool, biggest int) int {
	_, ok := visited[node]
	if ok {
		return biggest
	} else {
		visited[node] = true
	}
	if node.score > biggest {
		biggest = node.score
	}
	for _, n := range node.children {
		curr := findBiggestScoreHelper(n, visited, biggest)
		if curr > biggest {
			biggest = curr
		}
	}
	return biggest
}

func main() {
}
