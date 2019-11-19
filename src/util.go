package avl

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func strPreOrder(tree *Node) string {
	if tree == nil {
		return ""
	}
	return fmt.Sprintf("%d %s%s", tree.key, strPreOrder(tree.leftChild), strPreOrder(tree.rigthChild))

}
