package avl

func (n *Node) rRotation() *Node {
	x := n.leftChild
	t := x.rigthChild

	n.leftChild = t
	x.rigthChild = n

	n.updateHeight()
	x.updateHeight()

	return x
}

func (n *Node) lRotation() *Node {
	x := n.rigthChild
	t := x.leftChild

	n.rigthChild = t
	x.leftChild = n

	n.updateHeight()
	x.updateHeight()

	return x
}
