package heapexamples

import (
	"reflect"
	"testing"
)

func checkAVL(t *testing.T, root *AVLNode) {
	t.Helper()
	if root == nil {
		return
	}
	bf := root.balanceFactor()
	if bf < -1 || bf > 1 {
		t.Errorf("AVL violation at node %d: balance factor = %d", root.Val, bf)
	}
	checkAVL(t, root.Left)
	checkAVL(t, root.Right)
}

func TestAVLInsert(t *testing.T) {
	var root *AVLNode
	for _, v := range []int{10, 20, 30, 40, 50, 25} {
		root = root.AVLInsert(v)
	}
	got := root.avlInorder()
	want := []int{10, 20, 25, 30, 40, 50}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AVLInsert: got %v, want %v", got, want)
	}
	checkAVL(t, root)
}

func TestAVLInsertLeftLeft(t *testing.T) {
	var root *AVLNode
	for _, v := range []int{30, 20, 10} {
		root = root.AVLInsert(v)
	}
	got := root.avlInorder()
	want := []int{10, 20, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AVL LL: got %v, want %v", got, want)
	}
	checkAVL(t, root)
	if root.Val != 20 {
		t.Errorf("AVL LL root: got %d, want 20", root.Val)
	}
}

func TestAVLInsertRightRight(t *testing.T) {
	var root *AVLNode
	for _, v := range []int{10, 20, 30} {
		root = root.AVLInsert(v)
	}
	checkAVL(t, root)
	if root.Val != 20 {
		t.Errorf("AVL RR root: got %d, want 20", root.Val)
	}
}

func TestAVLInsertLeftRight(t *testing.T) {
	var root *AVLNode
	for _, v := range []int{30, 10, 20} {
		root = root.AVLInsert(v)
	}
	checkAVL(t, root)
	if root.Val != 20 {
		t.Errorf("AVL LR root: got %d, want 20", root.Val)
	}
}

func TestAVLInsertRightLeft(t *testing.T) {
	var root *AVLNode
	for _, v := range []int{10, 30, 20} {
		root = root.AVLInsert(v)
	}
	checkAVL(t, root)
	if root.Val != 20 {
		t.Errorf("AVL RL root: got %d, want 20", root.Val)
	}
}

func TestAVLDelete(t *testing.T) {
	var root *AVLNode
	for _, v := range []int{10, 20, 30, 40, 50, 25} {
		root = root.AVLInsert(v)
	}

	root = root.AVLDelete(50)
	got := root.avlInorder()
	want := []int{10, 20, 25, 30, 40}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AVLDelete leaf: got %v, want %v", got, want)
	}
	checkAVL(t, root)

	root = root.AVLDelete(30)
	got = root.avlInorder()
	want = []int{10, 20, 25, 40}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AVLDelete two children: got %v, want %v", got, want)
	}
	checkAVL(t, root)

	for _, v := range []int{10, 20, 25, 40} {
		root = root.AVLDelete(v)
		checkAVL(t, root)
	}
	if root != nil {
		t.Errorf("AVLDelete all: expected nil, got %v", root)
	}
}
