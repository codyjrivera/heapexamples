package heapexamples

import (
	"reflect"
	"testing"
)

func TestBSTContains(t *testing.T) {
	var root *TreeNode
	for _, v := range []int{5, 3, 7, 1, 4, 6, 8} {
		root = root.BSTInsert(v)
	}
	for _, v := range []int{5, 3, 7, 1, 4, 6, 8} {
		if !root.BSTContains(v) {
			t.Errorf("BSTContains(%d) = false, want true", v)
		}
	}
	for _, v := range []int{0, 2, 9, -1} {
		if root.BSTContains(v) {
			t.Errorf("BSTContains(%d) = true, want false", v)
		}
	}
	var nilRoot *TreeNode
	if nilRoot.BSTContains(1) {
		t.Error("BSTContains on nil tree = true, want false")
	}
}

func TestBSTInsert(t *testing.T) {
	var root *TreeNode
	for _, v := range []int{5, 3, 7, 1, 4, 6, 8} {
		root = root.BSTInsert(v)
	}
	got := root.bstInorder()
	want := []int{1, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BSTInsert: got %v, want %v", got, want)
	}
}

func TestBSTDelete(t *testing.T) {
	var root *TreeNode
	for _, v := range []int{5, 3, 7, 1, 4, 6, 8} {
		root = root.BSTInsert(v)
	}

	root = root.BSTDelete(1)
	got := root.bstInorder()
	want := []int{3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BSTDelete leaf: got %v, want %v", got, want)
	}

	root = root.BSTDelete(3)
	got = root.bstInorder()
	want = []int{4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BSTDelete one child: got %v, want %v", got, want)
	}

	root = root.BSTDelete(5)
	got = root.bstInorder()
	want = []int{4, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BSTDelete two children: got %v, want %v", got, want)
	}

	root = root.BSTDelete(99)
	got = root.bstInorder()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BSTDelete not found: got %v, want %v", got, want)
	}
}
