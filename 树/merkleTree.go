package main

import (
	"crypto/sha256"
	"fmt"
)

type MerkleNode struct {
	left  *MerkleNode
	right *MerkleNode
	data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := &MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.data = hash[:]
	} else {
		preHashes := append(left.data, right.data...)
		hash := sha256.Sum256(preHashes)
		node.data = hash[:]
	}
	node.left = left
	node.right = right
	return node
}

func NewMerkleTree(data [][]byte) *MerkleNode {
	var nodes []*MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, item := range data {
		node := NewMerkleNode(nil, nil, item)
		nodes = append(nodes, node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newNodes []*MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(nodes[j], nodes[j+1], nil)
			newNodes = append(newNodes, node)
		}
		// 两两结合，防止出现越界，
		if len(newNodes)%2 != 0 {
			nodes = append(newNodes, newNodes[len(newNodes)-1])
		}
	}
	return nodes[0]
}

func main() {
	fmt.Println(NewMerkleTree([][]byte{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}}))
}
