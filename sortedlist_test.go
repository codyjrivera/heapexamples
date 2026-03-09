package heapexamples

import (
	"reflect"
	"testing"
)

func TestSortedInsert(t *testing.T) {
	var head *ListNode
	for _, v := range []int{5, 1, 3, 2, 4} {
		head = head.SortedInsert(v)
	}
	got := head.toSlice()
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedInsert: got %v, want %v", got, want)
	}

	head = head.SortedInsert(3)
	got = head.toSlice()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedInsert duplicate: got %v, want %v", got, want)
	}
}

func TestSortedInsertRecursive(t *testing.T) {
	var head *ListNode
	for _, v := range []int{5, 1, 3, 2, 4} {
		head = head.SortedInsertRecursive(v)
	}
	got := head.toSlice()
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedInsertRecursive: got %v, want %v", got, want)
	}

	head = head.SortedInsertRecursive(3)
	got = head.toSlice()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedInsertRecursive duplicate: got %v, want %v", got, want)
	}
}

func TestSortedDelete(t *testing.T) {
	head := fromSlice([]int{1, 2, 3, 4, 5})
	head = head.SortedDelete(3)
	got := head.toSlice()
	want := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedDelete middle: got %v, want %v", got, want)
	}

	head = head.SortedDelete(1)
	got = head.toSlice()
	want = []int{2, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedDelete head: got %v, want %v", got, want)
	}

	head = head.SortedDelete(9)
	got = head.toSlice()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedDelete not found: got %v, want %v", got, want)
	}
}

func TestSortedDeleteRecursive(t *testing.T) {
	head := fromSlice([]int{1, 2, 3, 4, 5})
	head = head.SortedDeleteRecursive(3)
	got := head.toSlice()
	want := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedDeleteRecursive middle: got %v, want %v", got, want)
	}

	head = head.SortedDeleteRecursive(1)
	got = head.toSlice()
	want = []int{2, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SortedDeleteRecursive head: got %v, want %v", got, want)
	}
}

func testMergeFunc(t *testing.T, mergeFn func(*ListNode, *ListNode) *ListNode) {
	t.Helper()
	tests := []struct {
		name string
		a, b []int
		want []int
	}{
		{"interleaved", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"a nil", nil, []int{1, 2}, []int{1, 2}},
		{"b nil", []int{1, 2}, nil, []int{1, 2}},
		{"both nil", nil, nil, nil},
		{"a shorter", []int{1}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
		{"b shorter", []int{1, 2, 3}, []int{4}, []int{1, 2, 3, 4}},
		{"duplicates", []int{1, 3}, []int{1, 2, 3}, []int{1, 1, 2, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeFn(fromSlice(tt.a), fromSlice(tt.b)).toSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T)          { testMergeFunc(t, (*ListNode).Merge) }
func TestMergeRecursive(t *testing.T) { testMergeFunc(t, (*ListNode).MergeRecursive) }
