package main

import (
	"fmt"
)

func main() {
	// test
	codec := new(Codec)
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	fmt.Printf("root: %v\n", root)
	fmt.Printf("deserialize and serialize: %v\n", codec.Deserialize(codec.Serialize(root)))
}
