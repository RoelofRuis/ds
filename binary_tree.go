package ds

// TreeNode represents a node in a binary tree.
type TreeNode[A any] struct {
	Value A
	Left  *TreeNode[A]
	Right *TreeNode[A]
}

// BinaryTree is an unbalanced binary tree.
type BinaryTree[A any] struct {
	Root       *TreeNode[A]
	Comparator Comparator[A]
}

// NewBinaryTree creates a new BinaryTree that will use the given Comparator to order its nodes.
func NewBinaryTree[A any](comparator Comparator[A]) *BinaryTree[A] {
	return &BinaryTree[A]{
		Root:       nil,
		Comparator: comparator,
	}
}

// Insert adds a new node with the specified value into the tree based on the Comparator.
func (t *BinaryTree[A]) Insert(value A) {
	newNode := &TreeNode[A]{Value: value}

	if t.Root == nil {
		t.Root = newNode
		return
	}

	currentNode := t.Root
	for {
		compResult := t.Comparator(value, currentNode.Value)
		if compResult < 0 {
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return
			}
			currentNode = currentNode.Left
		} else { // Equal or greater values go to the right
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return
			}
			currentNode = currentNode.Right
		}
	}
}

// TraverseInOrder performs an in-order traversal of the tree, applying the given function to each node.
func (t *BinaryTree[A]) TraverseInOrder(f func(A)) {
	var inOrder func(node *TreeNode[A])
	inOrder = func(node *TreeNode[A]) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		f(node.Value)
		inOrder(node.Right)
	}

	inOrder(t.Root)
}
