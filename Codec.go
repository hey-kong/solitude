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

// Serializes a tree to a single string.
func (this *Codec) serialize2(root *TreeNode) string {
	res := ""
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res += "nil,"
		} else {
			res += strconv.Itoa(node.Val) + ","
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize2(data string) *TreeNode {
	l := strings.Split(data[:len(data)-1], ",")
	if l[0] == "nil" {
		return nil
	}

	val, _ := strconv.Atoi(l[0])
	l = l[1:]
	root := &TreeNode{}
	root.Val = val
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur != nil {
			// left
			left := &TreeNode{}
			if l[0] == "nil" {
				left = nil
			} else {
				val, _ := strconv.Atoi(l[0])
				left.Val = val
			}
			l = l[1:]
			cur.Left = left
			queue = append(queue, left)

			// right
			right := &TreeNode{}
			if l[0] == "nil" {
				right = nil
			} else {
				val, _ := strconv.Atoi(l[0])
				right.Val = val
			}
			l = l[1:]
			cur.Right = right
			queue = append(queue, right)
		}
	}
	return root
}
