package ds

// Interval represents a single interval.
type Interval struct {
	Low, High int
}

// TreeNode represents a single node in the interval tree.
type TreeNode struct {
	Interval Interval
	Max      int
	Left     *TreeNode
	Right    *TreeNode
}

// IntervalTree represents the root of an interval tree.
type IntervalTree struct {
	Root *TreeNode
}

// Insert adds a new interval to the tree.
func (tree *IntervalTree) Insert(interval Interval) {
	if tree.Root == nil {
		tree.Root = &TreeNode{
			Interval: interval,
			Max:      interval.High,
		}
	} else {
		tree.Root.insert(interval)
	}
}

// Search for overlapping intervals in the tree.
func (tree *IntervalTree) Search(interval Interval) (overlaps []Interval) {
	return tree.Root.search(interval)
}

// insert a new interval into the subtree rooted at the node.
func (node *TreeNode) insert(interval Interval) {
	if interval.Low < node.Interval.Low {
		if node.Left == nil {
			node.Left = &TreeNode{
				Interval: interval,
				Max:      interval.High,
			}
		} else {
			node.Left.insert(interval)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{
				Interval: interval,
				Max:      interval.High,
			}
		} else {
			node.Right.insert(interval)
		}
	}

	// Update the Max value
	if node.Max < interval.High {
		node.Max = interval.High
	}
}

// search for overlapping intervals in the subtree rooted at the node.
func (node *TreeNode) search(interval Interval) (overlaps []Interval) {
	if node == nil {
		return nil
	}

	// If there is overlap, add the interval to the result
	if node.Interval.Low <= interval.High && node.Interval.High >= interval.Low {
		overlaps = append(overlaps, node.Interval)
	}

	// If left child is present and could have an overlapping interval, search left subtree
	if node.Left != nil && node.Left.Max >= interval.Low {
		overlaps = append(overlaps, node.Left.search(interval)...)
	}

	// If right child could have an overlapping interval, search right subtree
	if node.Right != nil && node.Interval.Low <= interval.High {
		overlaps = append(overlaps, node.Right.search(interval)...)
	}

	return overlaps
}
