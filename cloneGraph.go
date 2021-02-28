package main

// Leetcode 133. (medium)
func cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	m := make(map[*GraphNode]*GraphNode)
	queue := make([]*GraphNode, 1)
	queue[0] = node
	res := newGraphNode(1)
	m[node] = res
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, n := range cur.Neighbors {
			if _, ok := m[n]; !ok {
				queue = append(queue, n)
				cpN := newGraphNode(n.Val)
				m[n] = cpN
			}
			m[cur].Neighbors = append(m[cur].Neighbors, m[n])
		}
	}
	return res
}

func newGraphNode(val int) *GraphNode {
	return &GraphNode{
		Val:       val,
		Neighbors: []*GraphNode{},
	}
}
