package heapexamples

type AVLNode struct {
	Val    int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

func (n *AVLNode) height() int {
	if n == nil {
		return -1
	}
	return n.Height
}

func (n *AVLNode) updateHeight() {
	l, r := n.Left.height(), n.Right.height()
	if l > r {
		n.Height = l + 1
	} else {
		n.Height = r + 1
	}
}

func (n *AVLNode) balanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.height() - n.Right.height()
}

func (n *AVLNode) rotateRight() *AVLNode {
	l := n.Left
	n.Left = l.Right
	l.Right = n
	n.updateHeight()
	l.updateHeight()
	return l
}

func (n *AVLNode) rotateLeft() *AVLNode {
	r := n.Right
	n.Right = r.Left
	r.Left = n
	n.updateHeight()
	r.updateHeight()
	return r
}

func (n *AVLNode) rebalance() *AVLNode {
	bf := n.balanceFactor()
	if bf > 1 {
		if n.Left.balanceFactor() < 0 {
			n.Left = n.Left.rotateLeft()
		}
		return n.rotateRight()
	}
	if bf < -1 {
		if n.Right.balanceFactor() > 0 {
			n.Right = n.Right.rotateRight()
		}
		return n.rotateLeft()
	}
	return n
}

func (n *AVLNode) avlMin() *AVLNode {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (root *AVLNode) avlInorder() []int {
	if root == nil {
		return nil
	}
	var out []int
	out = append(out, root.Left.avlInorder()...)
	out = append(out, root.Val)
	out = append(out, root.Right.avlInorder()...)
	return out
}

func (root *AVLNode) AVLInsert(v int) *AVLNode {
	if root == nil {
		return &AVLNode{Val: v}
	}
	if v < root.Val {
		root.Left = root.Left.AVLInsert(v)
	} else if v > root.Val {
		root.Right = root.Right.AVLInsert(v)
	} else {
		return root
	}
	root.updateHeight()
	return root.rebalance()
}

func (root *AVLNode) AVLDelete(v int) *AVLNode {
	if root == nil {
		return nil
	}
	if v < root.Val {
		root.Left = root.Left.AVLDelete(v)
	} else if v > root.Val {
		root.Right = root.Right.AVLDelete(v)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		successor := root.Right.avlMin()
		root.Val = successor.Val
		root.Right = root.Right.AVLDelete(successor.Val)
	}
	root.updateHeight()
	return root.rebalance()
}
