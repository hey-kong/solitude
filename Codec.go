package main

import (
	"strconv"
	"strings"
)

// Leetcode 297. (hard)
type Codec struct {
	l []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return subSerialize(root, "")
}

func subSerialize(root *TreeNode, s string) string {
	if root == nil {
		return s + "nil,"
	}
	s += strconv.Itoa(root.Val) + ","
	s = subSerialize(root.Left, s)
	s = subSerialize(root.Right, s)
	return s
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.l = strings.Split(data[:len(data)-1], ",")
	return this.subDeserialize()
}

func (this *Codec) subDeserialize() *TreeNode {
	if this.l[0] == "nil" {
		this.l = this.l[1:]
		return nil
	}

	val, _ := strconv.Atoi(this.l[0])
	this.l = this.l[1:]
	node := &TreeNode{}
	node.Val = val
	node.Left = this.subDeserialize()
	node.Right = this.subDeserialize()
	return node
}
