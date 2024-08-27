package Meta

import (
	"log"
	"testing"
)

// 1514m Path with Maximum Probability
func Test1514(t *testing.T) {
	log.Print("0.25 ?= ", maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2))
	log.Print("0.3 ?= ", maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2))
	log.Print("0 ?= ", maxProbability(3, [][]int{{0, 1}}, []float64{0.5}, 0, 2))
}
