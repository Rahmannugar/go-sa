package main

import "fmt"

type User struct {
	ID   int
	Name string
}

type BSTNode struct {
	left  *BSTNode
	right *BSTNode
	val   *User
}

// insert adds a user by ID and ignores duplicate IDs.
// Time complexity: O(h), Space complexity: O(h)
func (n *BSTNode) insert(user *User) {
	if n.val == nil {
		n.val = user
		return
	}

	if user.ID == n.val.ID {
		return
	}

	if user.ID < n.val.ID {
		if n.left == nil {
			n.left = &BSTNode{val: user}
		} else {
			n.left.insert(user)
		}
		return
	}

	if n.right == nil {
		n.right = &BSTNode{val: user}
	} else {
		n.right.insert(user)
	}
}

// getMinimum prints the user with the smallest ID.
// Time complexity: O(h), Space complexity: O(h)
func (n *BSTNode) getMinimum() {
	if n == nil {
		return
	}

	if n.left == nil {
		fmt.Printf("Minimum value: %d: %s\n", n.val.ID, n.val.Name)
		return
	}

	n.left.getMinimum()
}

// getMaximum prints the user with the largest ID.
// Time complexity: O(h), Space complexity: O(h)
func (n *BSTNode) getMaximum() {
	if n == nil {
		return
	}

	if n.right == nil {
		fmt.Printf("Maximum value: %d: %s\n", n.val.ID, n.val.Name)
		return
	}

	n.right.getMaximum()
}

// delete removes a user by ID and returns the updated subtree root.
// Time complexity: O(h), Space complexity: O(h)
func (n *BSTNode) delete(user *User) *BSTNode {
	if n == nil || user == nil {
		return n
	}

	if user.ID < n.val.ID {
		n.left = n.left.delete(user)
	} else if user.ID > n.val.ID {
		n.right = n.right.delete(user)
	} else {
		if n.left == nil && n.right == nil {
			return nil
		}

		if n.left == nil {
			return n.right
		}

		if n.right == nil {
			return n.left
		}

		minRight := n.right
		for minRight.left != nil {
			minRight = minRight.left
		}

		n.val = minRight.val
		n.right = n.right.delete(minRight.val)
	}

	return n
}

// inOrder prints users in ascending ID order.
// Time complexity: O(n), Space complexity: O(h)
func (n *BSTNode) inOrder() {
	if n == nil {
		return
	}

	n.left.inOrder()
	fmt.Printf("%d: %s\n", n.val.ID, n.val.Name)
	n.right.inOrder()
}

// preOrder prints the root before its left and right subtrees.
// Time complexity: O(n), Space complexity: O(h)
func (n *BSTNode) preOrder() {
	if n == nil {
		return
	}

	fmt.Printf("%d: %s\n", n.val.ID, n.val.Name)
	n.left.preOrder()
	n.right.preOrder()
}

// postOrder prints the left and right subtrees before the root.
// Time complexity: O(n), Space complexity: O(h)
func (n *BSTNode) postOrder() {
	if n == nil {
		return
	}

	n.left.postOrder()
	n.right.postOrder()
	fmt.Printf("%d: %s\n", n.val.ID, n.val.Name)
}

// exists reports whether a user with the given ID is in the tree.
// Time complexity: O(h), Space complexity: O(h)
func (n *BSTNode) exists(id int) bool {
	if n == nil {
		return false
	}

	if id == n.val.ID {
		return true
	}

	if id < n.val.ID {
		return n.left.exists(id)
	}

	return n.right.exists(id)
}

func main() {
	root := &BSTNode{}

	users := []*User{
		{ID: 10, Name: "John"},
		{ID: 5, Name: "Jane"},
		{ID: 15, Name: "Mike"},
		{ID: 3, Name: "David"},
		{ID: 7, Name: "Sarah"},
		{ID: 5, Name: "Duplicate"},
	}

	for _, user := range users {
		root.insert(user)
	}

	fmt.Println("In-order:")
	root.inOrder()

	fmt.Println("\nPre-order:")
	root.preOrder()

	fmt.Println("\nPost-order:")
	root.postOrder()

	root.getMinimum()
	root.getMaximum()

	fmt.Println("\nExists ID 7:", root.exists(7))
	fmt.Println("Exists ID 20:", root.exists(20))

	root = root.delete(&User{ID: 5})

	fmt.Println("\nAfter deleting ID 5:")
	root.inOrder()
}
