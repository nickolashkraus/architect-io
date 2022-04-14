# Architect

Solutions for a technical interview

## Problem 1

Create a function. Takes in a string. Returns true if there are duplicate
characters in the string otherwise returns false.

**Solution**

Create a hash set (hash table). Add characters to the hash set if they are not
already in the hash set, else return true.

## Problem 2

Given a list of trees (as pointers to their roots), return the biggest score of
all nodes in all trees.

Definition of a `Node`:

```go
type Node struct {
	id       uuid.UUID
	score    int
	children []*Node
}
```

```go
func findBiggestScore(trees []*Node) int {
	return -1
}
```

**Solution**

Traverse the tree keeping track of the biggest value in each tree. A helper
function is used in order to recursively traverse the tree and keep track of
visited nodes.
