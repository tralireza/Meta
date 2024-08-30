package Meta

import (
	"log"
	"testing"
)

// 797m All Paths from Source to Target
func Test797(t *testing.T) {
	// graph: is a DAG

	log.Print(" ?= ", allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	log.Print(" ?= ", allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}

// 1514m Path with Maximum Probability
func Test1514(t *testing.T) {
	log.Print("0.25 ?= ", maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2))
	log.Print("0.3 ?= ", maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2))
	log.Print("0 ?= ", maxProbability(3, [][]int{{0, 1}}, []float64{0.5}, 0, 2))
}

// 2699h Modify Graph Edge Weights
func Test2699(t *testing.T) {
	log.Print(" ?= ", modifiedGraphEdges(5, [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}}, 0, 1, 5))
	log.Print(" ?= ", modifiedGraphEdges(3, [][]int{{0, 1, -1}, {0, 2, 5}}, 0, 2, 6))
	log.Print(" ?= ", modifiedGraphEdges(4, [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, -1}}, 0, 2, 6))
	log.Print(" ?= ", modifiedGraphEdges(4, [][]int{{3, 0, -1}, {1, 2, -1}, {2, 3, -1}, {1, 3, 9}, {2, 0, 5}}, 0, 1, 7))
}
