package main

import (
	"fmt"
)

// TreeLinkNode. Definition for binary tree with next pointer.
type TreeLinkNode struct {
	Val               int
	Left, Right, Next *TreeLinkNode
}

func (t *TreeLinkNode) String() string {
	return fmt.Sprintf("{Val: %d, Left: %v, Right: %v, Next: %v}", t.Val, t.Left, t.Right, t.Next)
}

func connect(root *TreeLinkNode) {
	if root == nil {
		return
	}
	c(root)
}

func c(nodes ...*TreeLinkNode) {
	if len(nodes) == 0 {
		return
	}
	nodes = append(nodes, nil)
	newNodes := []*TreeLinkNode{}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
		if nodes[i].Left != nil {
			newNodes = append(newNodes, nodes[i].Left)
		}
		if nodes[i].Right != nil {
			newNodes = append(newNodes, nodes[i].Right)
		}
	}
	if len(newNodes) != 0 {
		c(newNodes...)
	}
}

func main() {
	// test
	root := &TreeLinkNode{Val: 1}
	root.Left = &TreeLinkNode{Val: 2}
	root.Right = &TreeLinkNode{Val: 3}
	root.Left.Left = &TreeLinkNode{Val: 4}
	root.Right.Left = &TreeLinkNode{Val: 5}

	fmt.Printf("tree:\n%v\n", root)
	connect(root)
	fmt.Printf("\n")
	fmt.Printf("connected:\n%v\n", root)
}

/////////////////////
// AC-ed JAVA code //
/////////////////////

/**
 * Definition for binary tree with next pointer.
 * public class TreeLinkNode {
 *     int val;
 *     TreeLinkNode left, right, next;
 *     TreeLinkNode(int x) { val = x; }
 * }
 */
// public class Solution {
//     public void connect(TreeLinkNode root) {
//         if (root == null){
//             return;
//         }
//         ArrayList<TreeLinkNode> source = new ArrayList<TreeLinkNode>();
//         source.add(root);
//         c(source);
//     }

//     private void c(ArrayList<TreeLinkNode> nodes){
//         if (nodes.isEmpty()) {
//             return;
//         }
//         ArrayList<TreeLinkNode> newNodes = new ArrayList<TreeLinkNode>();
//         nodes.add(null);
//         TreeLinkNode[] oldNodes = new TreeLinkNode[nodes.size()];
//         nodes.toArray(oldNodes);
//         for (int i = 0; i < oldNodes.length - 1; i++){
//             oldNodes[i].next = oldNodes[i+1];
//             if (oldNodes[i].left != null){
//                 newNodes.add(oldNodes[i].left);
//             }
//             if (oldNodes[i].right != null){
//                 newNodes.add(oldNodes[i].right);
//             }
//         }
//         if (newNodes.size() > 0){
//             c(newNodes);
//         }
//     }
// }
