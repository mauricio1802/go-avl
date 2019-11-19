package avl

func InsertTest() bool {
	tree := NewNode(10, struct{}{})
	tree = Insert(tree, NewNode(20, struct{}{}))
	tree = Insert(tree, NewNode(30, struct{}{}))
	tree = Insert(tree, NewNode(40, struct{}{}))
	tree = Insert(tree, NewNode(50, struct{}{}))
	tree = Insert(tree, NewNode(25, struct{}{}))

	return strPreOrder(tree) == "30 20 10 25 40 50 "

}

func DeleteTest() bool {
	tree := NewNode(9, struct{}{})
	tree = Insert(tree, NewNode(5, struct{}{}))
	tree = Insert(tree, NewNode(10, struct{}{}))
	tree = Insert(tree, NewNode(1, struct{}{}))
	tree = Insert(tree, NewNode(6, struct{}{}))
	tree = Insert(tree, NewNode(11, struct{}{}))
	tree = Insert(tree, NewNode(0, struct{}{}))
	tree = Insert(tree, NewNode(2, struct{}{}))
	tree = Insert(tree, NewNode(3, struct{}{}))

	buildOk := strPreOrder(tree) == "9 2 1 0 5 3 6 10 11 "
	//fmt.Println(strPreOrder(tree))
	tree = Delete(tree, 10)

	deleteOk := strPreOrder(tree) == "2 1 0 9 5 3 6 11 "

	return buildOk && deleteOk
}
