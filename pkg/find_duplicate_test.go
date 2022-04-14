package pkg

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	s := "abcdefg"
	ret := findDuplicate(s)
	if ret {
		t.Fail()
	}
	// 'a' is a duplicate
	s = "abcdefga"
	ret = findDuplicate(s)
	if !ret {
		t.Fail()
	}
}
