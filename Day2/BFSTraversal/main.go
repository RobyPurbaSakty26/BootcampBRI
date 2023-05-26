package main

type Node struct {
	data int 
	left, right *Node
}

func NewNode (data int) *Node {
	return &Node{data: data, left: nil, right: nil}
}

// func BFSTRavesial (root *Node){
// 	if root == nil {
// 		return
// 	}

// 	for len(queue) > 0 {
// 		node =: qu
// 	}
// }

func main () {

}