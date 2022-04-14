package pkg

import (
	"testing"
)

func TestFindBiggestScore(t *testing.T) {
	tree1 := createTree(3)
	tree2 := createTree(3)
	tree2.children = append(tree2.children, tree1)
	tree3 := createTree(3)
	tree3.children = append(tree3.children, tree1)
	// Explictly set biggest score, so we know if the answer is correct.
	tree1.children[0].children[2].score = 1337
	trees := []*Node{tree1, tree2, tree3}
	ret := findBiggestScore(trees)
	if ret != 1337 {
		t.Fail()
	}
}
