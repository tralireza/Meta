package Meta

import "log"

func init() {
	log.SetFlags(0)
}

// 117m Populating Next Right Pointers in Each Node II
type Node struct {
	Val               int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
	Q := []*Node{}
	if root != nil {
		Q = append(Q, root)
	}

	var v *Node
	for len(Q) > 0 {
		for i := range Q[:len(Q)-1] {
			Q[i].Next = Q[i+1] // Wire-up
		}
		for range len(Q) { // BSF
			v, Q = Q[0], Q[1:]
			for _, u := range []*Node{v.Left, v.Right} {
				if u != nil {
					Q = append(Q, u)
				}
			}
		}
	}

	return root
}
