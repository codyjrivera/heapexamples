package heapexamples

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Contains(v int) bool {
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val == v {
			return true
		}
	}
	return false
}

func (head *ListNode) ContainsRecursive(v int) bool {
	if head == nil {
		return false
	}
	if head.Val == v {
		return true
	}
	return head.Next.ContainsRecursive(v)
}

func (head *ListNode) InsertFront(v int) *ListNode {
	return &ListNode{Val: v, Next: head}
}

func (head *ListNode) InsertBack(v int) *ListNode {
	newNode := &ListNode{Val: v}
	if head == nil {
		return newNode
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	return head
}

func (head *ListNode) InsertBackRecursive(v int) *ListNode {
	if head == nil {
		return &ListNode{Val: v}
	}
	head.Next = head.Next.InsertBackRecursive(v)
	return head
}

func (head *ListNode) DeleteAll(v int) *ListNode {
	for head != nil && head.Val == v {
		head = head.Next
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == v {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func (head *ListNode) DeleteAllRecursive(v int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = head.Next.DeleteAllRecursive(v)
	if head.Val == v {
		return head.Next
	}
	return head
}

func (head *ListNode) CopyList() *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Val: head.Val}
	cur := head.Next
	tail := newHead
	for cur != nil {
		tail.Next = &ListNode{Val: cur.Val}
		tail = tail.Next
		cur = cur.Next
	}
	return newHead
}

func (head *ListNode) CopyListRecursive() *ListNode {
	if head == nil {
		return nil
	}
	return &ListNode{Val: head.Val, Next: head.Next.CopyListRecursive()}
}
