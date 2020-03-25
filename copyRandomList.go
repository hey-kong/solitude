package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Leetcode 138. (medium)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 生成新节点
	p := head
	for p != nil {
		newNode := new(Node)
		newNode.Val = p.Val
		newNode.Next = p.Next
		newNode.Random = nil
		p.Next = newNode
		p = p.Next.Next
	}

	// 修改 Random 指针
	p = head
	for p != nil {
		originRandom := p.Random
		if originRandom != nil {
			p.Next.Random = originRandom.Next
		} else {
			p.Next.Random = nil
		}
		p = p.Next.Next
	}

	// 修改 Next 指针
	p = head
	ret := p.Next
	for p != nil {
		originNext := p.Next.Next
		if originNext != nil {
			p.Next.Next = originNext.Next
		} else {
			p.Next.Next = nil
		}
		p.Next = originNext
		p = originNext
	}
	return ret
}
