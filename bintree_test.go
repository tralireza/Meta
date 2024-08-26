package Meta

import (
	"log"
	"testing"
)

// 117m Populating Next Right Pointers in Each Node II
func Test117(t *testing.T) {
	type N = Node

	for _, tree := range []*N{
		{1, &N{2, &N{Val: 4}, &N{Val: 5}, nil}, &N{3, nil, &N{Val: 7}, nil}, nil},
		{},
	} {
		log.Print(" ?= ", connect(tree))
	}
}

// 590 N-ary Tree Postorder Traversal
func Test590(t *testing.T) {
	type N = Node590

	log.Print("", postorder(&N{1, []*N{{3, []*N{{Val: 5}, {Val: 6}}}, {Val: 2}, {Val: 4}}}))
}
