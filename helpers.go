package heapexamples

// --- List helpers (for tests) ---

func (head *ListNode) toSlice() []int {
	var out []int
	for cur := head; cur != nil; cur = cur.Next {
		out = append(out, cur.Val)
	}
	return out
}

func fromSlice(vals []int) *ListNode {
	var head *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		head = &ListNode{Val: vals[i], Next: head}
	}
	return head
}
