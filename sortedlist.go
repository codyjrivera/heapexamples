package heapexamples

// this is a sorted linked list which allows no duplicates (to make
// specification and verification easier).

func (head *ListNode) SortedInsert(v int) *ListNode {
	if head == nil || v < head.Val {
		return &ListNode{Val: v, Next: head}
	}
	if v == head.Val {
		return head
	}
	cur := head
	for cur.Next != nil && cur.Next.Val < v {
		cur = cur.Next
	}
	if cur.Next != nil && cur.Next.Val == v {
		return head
	}
	cur.Next = &ListNode{Val: v, Next: cur.Next}
	return head
}

func (head *ListNode) SortedInsertRecursive(v int) *ListNode {
	if head == nil || v < head.Val {
		return &ListNode{Val: v, Next: head}
	}
	if v == head.Val {
		return head
	}
	head.Next = head.Next.SortedInsertRecursive(v)
	return head
}

func (head *ListNode) SortedDelete(v int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == v {
		return head.Next
	}
	cur := head
	for cur.Next != nil && cur.Next.Val != v {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	return head
}

func (head *ListNode) SortedDeleteRecursive(v int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == v {
		return head.Next
	}
	head.Next = head.Next.SortedDeleteRecursive(v)
	return head
}

func (a *ListNode) Merge(b *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

func (a *ListNode) MergeRecursive(b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val <= b.Val {
		a.Next = a.Next.MergeRecursive(b)
		return a
	}
	b.Next = a.MergeRecursive(b.Next)
	return b
}
