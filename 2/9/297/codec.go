package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func str(nodes ...*TreeNode) string {
	r := ""
	next := []*TreeNode{}
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == nil {
			r += "null"
		} else {
			r += fmt.Sprintf("%d", nodes[i].Val)
			next = append(next, nodes[i].Left, nodes[i].Right)
		}
		if i+1 < len(nodes) {
			r += ", "
		}
	}
	if len(next) > 0 {
		r += ", "
		r += str(next...)
	}
	return r
}

func (t *TreeNode) String() string {
	return str(t)
}

type Codec struct{}

// Encodes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	return "[" + serialize(root) + "]"
}

func serialize(root ...*TreeNode) string {
	if len(root) == 0 {
		return ""
	}
	var (
		res  string
		next []*TreeNode = []*TreeNode{}
	)
	for i := 0; i < len(root); i++ {
		if root[i] == nil {
			res += ",null"
		} else {
			res += fmt.Sprintf(",%d", root[i].Val)
			next = append(next, root[i].Left, root[i].Right)
		}
	}
	res = string([]byte(res)[1:])
	if len(next) > 0 {
		res += "," + serialize(next...)
	}
	return res
}

// Decodes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	b := []byte(data)
	if len(b) < 2 {
		panic("Invalid input")
	}
	dataArray := strings.Split(string(b[1:len(b)-1]), ",")
	if len(dataArray) == 0 {
		return nil
	}
	val, err := strconv.Atoi(dataArray[0])
	if err != nil {
		panic(err.Error())
	}
	root := &TreeNode{Val: val}
	deserialize(dataArray[1:], root)
	return root
}

func deserialize(data []string, root ...*TreeNode) {
	if len(data) < 2*len(root) {
		panic("Unknown error")
	}
	if len(data) == 0 || len(root) == 0 {
		return
	}
	var (
		i    int
		next []*TreeNode = []*TreeNode{}
	)
	for i = 0; i < len(root); i++ {
		j := 2 * i
		if data[j] == "null" {
			root[i].Left = nil
		} else {
			val, err := strconv.Atoi(data[j])
			if err != nil {
				panic(err.Error())
			}
			root[i].Left = &TreeNode{Val: val}
			next = append(next, root[i].Left)
		}
		if data[j+1] == "null" {
			root[i].Right = nil
		} else {
			val, err := strconv.Atoi(data[j+1])
			if err != nil {
				panic(err.Error())
			}
			root[i].Right = &TreeNode{Val: val}
			next = append(next, root[i].Right)
		}
	}
	if len(next) > 0 {
		deserialize(data[2*i:], next...)
	}
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));

//////////////////////
// AC-ed JAVA code  //
//////////////////////

// public class Codec {
//     // Encodes a tree to a single string.
//     public String serialize(TreeNode root) {
//         //
//         List<TreeNode> nodes = new ArrayList<TreeNode>();
//         nodes.add(root);
//         return "[" + s(nodes).substring(1) + "]";
//     }

//     private String s(List<TreeNode> nodes){
//         //
//         String res = ",";
//         if (nodes.isEmpty()) {
//             return res;
//         }
//         List<TreeNode> next = new ArrayList<TreeNode>();
//         Iterator<TreeNode> iter = nodes.iterator();
//         while (iter.hasNext()) {
//             TreeNode node = iter.next();
//             if (node == null) {
//                 res += ",null";
//             }else{
//                 res += "," + String.valueOf(node.val);
//                 next.add(node.left);
//                 next.add(node.right);
//             }
//         }
//         res = res.substring(1);
//         if (!next.isEmpty()) {
//             res += s(next);
//         }
//         return res;
//     }

//     // Decodes your encoded data to tree.
//     public TreeNode deserialize(String data) {
//         //
//         if (data.length() < 2) {
//             return null;
//         }
//         String[] dataArray = data.substring(1, data.length()-1).split(",");
//         if (dataArray.length == 0) {
//             return null;
//         }
//         if (dataArray[0].equals("null")){
//             return null;
//         }
//         List<String> dataList = Arrays.asList(dataArray);
//         TreeNode root = new TreeNode(0);
//         Iterator<String> iter = dataList.iterator();
//         if (iter.hasNext()) {
//             root.val = Integer.parseInt(iter.next());
//         }
//         List<TreeNode> nodes = new ArrayList();
//         nodes.add(root);
//         d(dataList.subList(1,dataList.size()), nodes);
//         return root;
//     }

//     private void d(List<String> dataList, List<TreeNode> nodes){
//         //
//         if (dataList.isEmpty() || nodes.isEmpty()) {
//             return;
//         }
//         //
//         List<TreeNode> next = new ArrayList<TreeNode>();
//         int i = 0;
//         for (i = 0; i < nodes.size(); i++){
//             int j = 2*i;
//             if (dataList.get(j).equals("null")){
//                 nodes.get(i).left = null;
//             }else{
//                 nodes.get(i).left = new TreeNode(Integer.parseInt(dataList.get(j)));
//                 next.add(nodes.get(i).left);
//             }
//             if (dataList.get(j+1).equals("null")){
//                 nodes.get(i).right = null;
//             }else{
//                 nodes.get(i).right = new TreeNode(Integer.parseInt(dataList.get(j+1)));
//                 next.add(nodes.get(i).right);
//             }
//         }
//         if (!next.isEmpty()){
//             d(dataList.subList(2*i,dataList.size()),next);
//         }
//     }
// }
