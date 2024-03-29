package ds

type Interval[A any] interface {
	Start() A
	End() A
}

type SimpleInterval[A any] struct {
	start A
	end   A
}

func (s SimpleInterval[A]) Start() A {
	return s.start
}

func (s SimpleInterval[A]) End() A {
	return s.end
}

// ContiguousIntervalTree is a non thread-save binary search tree storing contiguous intervals. This means that the
// intervals are always directly following each other without overlap.
// The special case are intervals with zero length, of which an unlimited number may be added.
// When iterating the tree, it will also return the empty intervals formed by the space in between inserted elements.
type ContiguousIntervalTree[K any, V any] struct {
	Comparator Comparator[K]
	Root       *ContiguousIntervalNode[K, V]
}

// NewContiguousIntervalTree creates a new tree for the given key type K, capable of storing values of type V.
func NewContiguousIntervalTree[K any, V any](comparator Comparator[K]) *ContiguousIntervalTree[K, V] {
	return &ContiguousIntervalTree[K, V]{
		Root:       nil,
		Comparator: comparator,
	}
}

// ContiguousIntervalNode is internally used in the ContiguousIntervalTree to store its data.
type ContiguousIntervalNode[K any, V any] struct {
	Interval Interval[K]
	Value    V
	Left     *ContiguousIntervalNode[K, V]
	Right    *ContiguousIntervalNode[K, V]
}

// size recursively calculates the size of the tree rooted in this node.
func (n *ContiguousIntervalNode[K, V]) size() int {
	if n == nil {
		return 0
	}
	return n.Left.size() + n.Right.size() + 1
}

func (n *ContiguousIntervalNode[K, V]) alphaBalanced(a float64) bool {
	left := float64(n.Left.size())
	right := float64(n.Right.size())
	alpha := a * float64(left+right+1)
	return left <= alpha && right <= alpha
}

// Insert inserts an interval into the tree. If the interval overlaps with an existing interval, this operation fails
// and returns false.
func (t *ContiguousIntervalTree[K, V]) Insert(i Interval[K], value V) bool {
	newNode := &ContiguousIntervalNode[K, V]{Interval: i, Value: value}

	if t.Root == nil {
		t.Root = newNode
		return true
	}

	currentNode := t.Root
	for {
		if t.Comparator(i.Start(), currentNode.Interval.Start()) <= 0 {
			if t.Comparator(i.End(), currentNode.Interval.Start()) > 0 {
				return false // overlap
			}
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return true
			}
			currentNode = currentNode.Left
		} else if t.Comparator(i.Start(), currentNode.Interval.End()) >= 0 {
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return true
			}
			currentNode = currentNode.Right
		} else {
			return false // overlap
		}
	}
}

func (t *ContiguousIntervalTree[K, V]) Find(point K) bool {
	return t.search(point) != nil
}

// search finds the node that contains a given point K
func (t *ContiguousIntervalTree[K, V]) search(point K) *ContiguousIntervalNode[K, V] {
	if t.Root == nil {
		return nil
	}

	sum := 0
	height := 0
	currentNode := t.Root
	for {
		if t.Comparator(point, currentNode.Interval.Start()) < 0 {
			if currentNode.Left == nil {
				if -sum == height {
					return nil
				}
				return &ContiguousIntervalNode[K, V]{}
			}
			currentNode = currentNode.Left
			sum -= 1
			height += 1
		} else if t.Comparator(point, currentNode.Interval.End()) > 0 {
			if currentNode.Right == nil {
				if sum == height {
					return nil
				}
				return &ContiguousIntervalNode[K, V]{}
			}
			currentNode = currentNode.Right
			sum += 1
			height += 1
		} else {
			return currentNode
		}
	}
}

// Size returns the number of nodes in the tree.
// This might be less than the number of contiguous intervals expressed by this tree. To get that, use NumIntervals.
func (t *ContiguousIntervalTree[K, V]) Size() int {
	if t.Root == nil {
		return 0
	}
	return t.Root.size()
}

// NumIntervals returns the number of contiguous intervals in the node. This includes intervals that are in between
// active intervals and that might not have any data associated with them.
func (t *ContiguousIntervalTree[K, V]) NumIntervals() int {
	count := 0
	t.TraverseInOrder(func(interval Interval[K], v V) {
		count += 1
	})
	return count
}

// TraverseInOrder visits the intervals in order, including the empty intervals formed by the space in between inserted
// elements.
func (t *ContiguousIntervalTree[K, V]) TraverseInOrder(f func(interval Interval[K], value V)) {
	var lastInterval Interval[K]

	var inOrder func(node *ContiguousIntervalNode[K, V])
	inOrder = func(node *ContiguousIntervalNode[K, V]) {
		if node == nil {
			return
		}
		inOrder(node.Left)

		if lastInterval != nil && t.Comparator(lastInterval.End(), node.Interval.Start()) < 0 {
			var zero V
			f(SimpleInterval[K]{lastInterval.End(), node.Interval.Start()}, zero)
		}

		f(node.Interval, node.Value)

		lastInterval = node.Interval
		inOrder(node.Right)
	}

	inOrder(t.Root)
}

func (t *ContiguousIntervalTree[K, V]) TraverseBetween(start, end Interval[K], f func(interval Interval[K], value V)) {
	// TODO: implement
}
