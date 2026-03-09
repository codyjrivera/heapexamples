package heapexamples

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) bstMin() *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func (root *TreeNode) bstInorder() []int {
	if root == nil {
		return nil
	}
	var out []int
	out = append(out, root.Left.bstInorder()...)
	out = append(out, root.Val)
	out = append(out, root.Right.bstInorder()...)
	return out
}

func (root *TreeNode) BSTContains(v int) bool {
	if root == nil {
		return false
	}
	if v < root.Val {
		return root.Left.BSTContains(v)
	}
	if v > root.Val {
		return root.Right.BSTContains(v)
	}
	return true
}

func (root *TreeNode) BSTInsert(v int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if v < root.Val {
		root.Left = root.Left.BSTInsert(v)
	} else if v > root.Val {
		root.Right = root.Right.BSTInsert(v)
	}
	return root
}

func (root *TreeNode) BSTDelete(v int) *TreeNode {
	if root == nil {
		return nil
	}
	if v < root.Val {
		root.Left = root.Left.BSTDelete(v)
	} else if v > root.Val {
		root.Right = root.Right.BSTDelete(v)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		successor := root.Right.bstMin()
		root.Val = successor.Val
		root.Right = root.Right.BSTDelete(successor.Val)
	}
	return root
}
