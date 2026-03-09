package heapexamples

import (
	"reflect"
	"testing"
)

func testContainsFunc(t *testing.T, containsFn func(*ListNode, int) bool) {
	t.Helper()
	tests := []struct {
		name string
		in   []int
		v    int
		want bool
	}{
		{"found at head", []int{1, 2, 3}, 1, true},
		{"found in middle", []int{1, 2, 3}, 2, true},
		{"found at tail", []int{1, 2, 3}, 3, true},
		{"not found", []int{1, 2, 3}, 9, false},
		{"empty list", nil, 1, false},
		{"single found", []int{5}, 5, true},
		{"single not found", []int{5}, 6, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := fromSlice(tt.in)
			if got := containsFn(head, tt.v); got != tt.want {
				t.Errorf("Contains(%v, %d) = %v, want %v", tt.in, tt.v, got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T)          { testContainsFunc(t, (*ListNode).Contains) }
func TestContainsRecursive(t *testing.T) { testContainsFunc(t, (*ListNode).ContainsRecursive) }

func TestInsertFront(t *testing.T) {
	var head *ListNode
	head = head.InsertFront(3)
	head = head.InsertFront(2)
	head = head.InsertFront(1)
	got := head.toSlice()
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertFront: got %v, want %v", got, want)
	}
}

func TestInsertBack(t *testing.T) {
	var head *ListNode
	head = head.InsertBack(1)
	head = head.InsertBack(2)
	head = head.InsertBack(3)
	got := head.toSlice()
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertBack: got %v, want %v", got, want)
	}
}

func TestInsertBackRecursive(t *testing.T) {
	var head *ListNode
	head = head.InsertBackRecursive(1)
	head = head.InsertBackRecursive(2)
	head = head.InsertBackRecursive(3)
	got := head.toSlice()
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InsertBackRecursive: got %v, want %v", got, want)
	}
}

func TestDeleteAll(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		v    int
		want []int
	}{
		{"middle", []int{1, 2, 3, 2, 4}, 2, []int{1, 3, 4}},
		{"head", []int{5, 5, 1, 2}, 5, []int{1, 2}},
		{"tail", []int{1, 2, 3, 3}, 3, []int{1, 2}},
		{"all same", []int{7, 7, 7}, 7, nil},
		{"not found", []int{1, 2, 3}, 9, []int{1, 2, 3}},
		{"empty", nil, 1, nil},
		{"single match", []int{4}, 4, nil},
		{"single no match", []int{4}, 5, []int{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := fromSlice(tt.in)
			got := head.DeleteAll(tt.v).toSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAll(%v, %d) = %v, want %v", tt.in, tt.v, got, tt.want)
			}
		})
	}
}

func testCopyFunc(t *testing.T, copyFn func(*ListNode) *ListNode) {
	t.Helper()
	tests := []struct {
		name string
		in   []int
	}{
		{"empty", nil},
		{"single", []int{1}},
		{"multiple", []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := fromSlice(tt.in)
			cp := copyFn(orig)
			if !reflect.DeepEqual(cp.toSlice(), orig.toSlice()) {
				t.Errorf("copy %v != original %v", cp.toSlice(), orig.toSlice())
			}
			// Mutate copy and verify original is unchanged.
			if cp != nil {
				cp.Val = -999
				if orig.Val == -999 {
					t.Error("mutating copy affected original — not a deep copy")
				}
			}
		})
	}
}

func TestCopyList(t *testing.T)          { testCopyFunc(t, (*ListNode).CopyList) }
func TestCopyListRecursive(t *testing.T) { testCopyFunc(t, (*ListNode).CopyListRecursive) }

func TestDeleteAllRecursive(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		v    int
		want []int
	}{
		{"middle", []int{1, 2, 3, 2, 4}, 2, []int{1, 3, 4}},
		{"head", []int{5, 5, 1, 2}, 5, []int{1, 2}},
		{"tail", []int{1, 2, 3, 3}, 3, []int{1, 2}},
		{"all same", []int{7, 7, 7}, 7, nil},
		{"not found", []int{1, 2, 3}, 9, []int{1, 2, 3}},
		{"empty", nil, 1, nil},
		{"single match", []int{4}, 4, nil},
		{"single no match", []int{4}, 5, []int{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := fromSlice(tt.in)
			got := head.DeleteAllRecursive(tt.v).toSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteAllRecursive(%v, %d) = %v, want %v", tt.in, tt.v, got, tt.want)
			}
		})
	}
}
