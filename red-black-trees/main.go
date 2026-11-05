package main

import "fmt"

type User struct {
	ID   int
	Name string
}

type RBNode struct {
	val    *User
	red    bool
	parent *RBNode
	left   *RBNode
	right  *RBNode
}

type RBTree struct {
	root *RBNode
	nil  *RBNode
}

// NewRBTree creates an empty tree with a black sentinel node.
func NewRBTree() *RBTree {
	nilNode := &RBNode{}
	nilNode.red = false

	nilNode.left = nilNode
	nilNode.right = nilNode
	nilNode.parent = nilNode

	return &RBTree{
		root: nilNode,
		nil:  nilNode,
	}
}

// Insert adds a user by ID and ignores duplicate IDs.
// Time complexity: O(log n), Space complexity: O(1)
func (t *RBTree) Insert(user *User) {
	if user == nil {
		return
	}

	// Create the new node.
	newNode := &RBNode{
		val:   user,
		red:   true,
		left:  t.nil,
		right: t.nil,
	}

	// Find where the new node belongs.
	parent := t.nil
	current := t.root

	for current != t.nil {
		parent = current

		if newNode.val.ID < current.val.ID {
			current = current.left
		} else if newNode.val.ID > current.val.ID {
			current = current.right
		} else {
			return
		}
	}

	// Connect the new node to its parent.
	newNode.parent = parent

	if parent == t.nil {
		t.root = newNode
	} else if newNode.val.ID < parent.val.ID {
		parent.left = newNode
	} else {
		parent.right = newNode
	}

	// Fix any Red-Black Tree violations.
	t.fixInsert(newNode)
}

// fixInsert restores the red-black tree rules after insertion.
// Time complexity: O(log n), Space complexity: O(1)
func (t *RBTree) fixInsert(node *RBNode) {
	for node.parent.red {
		if node.parent == node.parent.parent.left {

			uncle := node.parent.parent.right

			// Case 1: Uncle is red.
			if uncle.red {
				node.parent.red = false
				uncle.red = false
				node.parent.parent.red = true

				node = node.parent.parent
			} else {
				// Case 2: Triangle.
				if node == node.parent.right {
					node = node.parent
					t.leftRotate(node)
				}

				// Case 3: Line.
				node.parent.red = false
				node.parent.parent.red = true
				t.rightRotate(node.parent.parent)
			}
		} else {
			// Mirror cases: parent is a right child.

			uncle := node.parent.parent.left

			// Case 1: Uncle is red.
			if uncle.red {
				node.parent.red = false
				uncle.red = false
				node.parent.parent.red = true

				node = node.parent.parent
			} else {
				// Case 2: Triangle.
				if node == node.parent.left {
					node = node.parent
					t.rightRotate(node)
				}

				// Case 3: Line.
				node.parent.red = false
				node.parent.parent.red = true
				t.leftRotate(node.parent.parent)
			}
		}
	}

	// The root must always be black.
	t.root.red = false
}

// leftRotate moves a node down to the left of its right child.
// Time complexity: O(1), Space complexity: O(1)
func (t *RBTree) leftRotate(x *RBNode) {
	y := x.right

	x.right = y.left

	if y.left != t.nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == t.nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

// rightRotate moves a node down to the right of its left child.
// Time complexity: O(1), Space complexity: O(1)
func (t *RBTree) rightRotate(x *RBNode) {
	y := x.left

	x.left = y.right

	if y.right != t.nil {
		y.right.parent = x
	}

	y.parent = x.parent

	if x.parent == t.nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}

	y.right = x
	x.parent = y
}

// Exists reports whether a user with the given ID is in the tree.
// Time complexity: O(log n), Space complexity: O(1)
func (t *RBTree) Exists(id int) bool {
	current := t.root

	for current != t.nil {
		if id == current.val.ID {
			return true
		}

		if id < current.val.ID {
			current = current.left
		} else {
			current = current.right
		}
	}

	return false
}

// InOrder prints users in ascending ID order with their node colors.
// Time complexity: O(n), Space complexity: O(log n)
func (t *RBTree) InOrder(node *RBNode) {
	if node == t.nil {
		return
	}

	t.InOrder(node.left)

	color := "Black"
	if node.red {
		color = "Red"
	}

	fmt.Printf("%d: %s (%s)\n",
		node.val.ID,
		node.val.Name,
		color,
	)

	t.InOrder(node.right)
}

func main() {
	tree := NewRBTree()

	users := []*User{
		{ID: 10, Name: "John"},
		{ID: 20, Name: "Jane"},
		{ID: 30, Name: "Mike"},
		{ID: 40, Name: "David"},
		{ID: 50, Name: "Sarah"},
		{ID: 25, Name: "Peter"},
		{ID: 5, Name: "Grace"},
	}

	for _, user := range users {
		tree.Insert(user)
	}

	fmt.Println("Red-Black Tree (In-order):")
	tree.InOrder(tree.root)

	fmt.Println("\nSearch results:")
	fmt.Println("ID 25 exists:", tree.Exists(25))
	fmt.Println("ID 100 exists:", tree.Exists(100))
}
