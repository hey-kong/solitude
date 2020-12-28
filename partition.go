package main

// Leetcode 131. (medium)
func partition(s string) [][]string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for l := 1; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			if s[i] == s[i+l-1] && ((l <= 3) || dp[i+1][i+l-2] == true) {
				dp[i][i+l-1] = true
			}
		}
	}
	res := [][]string{}
	res = dfsPartition(s, 0, []string{}, dp, res)
	return res
}

func dfsPartition(s string, i int, tmp []string, dp [][]bool, res [][]string) [][]string {
	if i == len(s) {
		ans := make([]string, len(tmp))
		copy(ans, tmp)
		res = append(res, ans)
		return res
	}
	for j := i; j < len(s); j++ {
		if dp[i][j] == true {
			tmp = append(tmp, s[i:j+1])
			res = dfsPartition(s, j+1, tmp, dp, res)
			tmp = tmp[:len(tmp)-1]
		}
	}
	return res
}

// Leetcode 86. (medium)
func partitionList(head *ListNode, x int) *ListNode {
	dummySmall, dummyBig := &ListNode{}, &ListNode{}
	sHead, bHead := dummySmall, dummyBig
	for head != nil {
		if head.Val < x {
			dummySmall.Next = head
			dummySmall = dummySmall.Next
		} else {
			dummyBig.Next = head
			dummyBig = dummyBig.Next
		}
		head = head.Next
	}
	dummySmall.Next = bHead.Next
	dummyBig.Next = nil
	return sHead.Next
}
