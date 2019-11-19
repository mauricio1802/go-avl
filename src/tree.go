package avl

//If existe a node with key == value return true and the node with key == value, if there is not return false and the node that would be it's father
func (n *Node) nodeByKey(value uint64) (bool, *Node) {
	if value == n.key || (value > n.key && !n.hasRChild()) || (value < n.key && !n.hasLChild()) {
		return true, n
	}
	if value < n.key {
		return n.leftChild.nodeByKey(value)
	}
	return n.rigthChild.nodeByKey(value)
}

func (n *Node) HasKey(value uint64) bool {
	answ, _ := n.nodeByKey(value)
	return answ
}

func (n *Node) updateHeight() {
	n.height = max(getHeight(n.rigthChild), getHeight(n.leftChild)) + 1
}

func Insert(tree, newNode *Node) *Node {
	if tree == nil {
		return newNode
	}

	if newNode.key < tree.key {
		tree.leftChild = Insert(tree.leftChild, newNode)
	} else if newNode.key > tree.key {
		tree.rigthChild = Insert(tree.rigthChild, newNode)
	} else {
		return tree
	}

	tree.updateHeight()

	bl := tree.balanceFactor()

	if bl > 1 && (newNode.key < tree.leftChild.key) {
		return tree.rRotation()
	}

	if bl < -1 && (newNode.key > tree.rigthChild.key) {
		return tree.lRotation()
	}

	if bl > 1 && (newNode.key > tree.leftChild.key) {
		tree.leftChild = tree.leftChild.lRotation()
		return tree.rRotation()
	}

	if bl < -1 && (newNode.key < tree.rigthChild.key) {
		tree.rigthChild = tree.rigthChild.rRotation()
		return tree.lRotation()
	}

	return tree
}

func Delete(tree *Node, key uint64) *Node {
	if tree == nil {
		return tree
	}

	if key < tree.key {
		tree.leftChild = Delete(tree.leftChild, key)
	} else if key > tree.key {
		tree.rigthChild = Delete(tree.rigthChild, key)
	} else {
		if tree.leftChild == nil || tree.rigthChild == nil {
			if tree.leftChild != nil {
				return tree.leftChild
			} else if tree.rigthChild != nil {
				return tree.rigthChild
			} else {
				return nil
			}
		}
		newRoot := getMinNode(tree)
		newRoot.leftChild = tree.leftChild
		newRoot.rigthChild = Delete(tree, newRoot.key)
		tree = newRoot
	}

	tree.updateHeight()

	bl := tree.balanceFactor()

	if bl > 1 && (tree.leftChild.balanceFactor() >= 0) {
		return tree.rRotation()
	}

	if bl < -1 && (tree.rigthChild.balanceFactor() <= 0) {
		return tree.lRotation()
	}

	if bl > 1 && (tree.leftChild.balanceFactor() < 0) {
		tree.leftChild = tree.leftChild.lRotation()
		return tree.rRotation()
	}

	if bl < -1 && (tree.rigthChild.balanceFactor() > 0) {
		tree.rigthChild = tree.rigthChild.rRotation()
		return tree.lRotation()
	}

	return tree

}

func (n *Node) balanceFactor() int {
	return getHeight(n.leftChild) - getHeight(n.rigthChild)
}

func getMinNode(tree *Node) *Node {
	for tree != nil && tree.leftChild != nil {
		tree = tree.leftChild
	}
	return tree
}
func getSuccessorNode(tree *Node) *Node {
	return getMinNode(tree.rigthChild)
}
