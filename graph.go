package Meta

import (
	"math"
)

// 797m All Paths from Source to Target
func allPathsSourceTarget(graph [][]int) [][]int {
	R := [][]int{}

	var r []int
	var BT func(int)
	BT = func(v int) {
		if v == len(graph)-1 {
			R = append(R, append([]int{}, r...))
			return
		}
		for _, u := range graph[v] {
			r = append(r, u)
			BT(u)
			r = r[:len(r)-1]
		}
	}

	r = append(r, 0)
	BT(0)

	return R
}

// 1514m Path with Maximum Probability
func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	// Bellman-Ford

	Dist := make([]float64, n)
	Dist[start_node] = 1

	for range n - 1 {
		updated := false // terminate early
		for i, e := range edges {
			p := succProb[i]
			v, u := e[0], e[1]

			// Relaxing: 2-way as for undirected graph
			if Dist[v]*p > Dist[u] {
				Dist[u] = Dist[v] * p
				updated = true
			}
			if Dist[u]*p > Dist[v] {
				Dist[v] = Dist[u] * p
				updated = true
			}
		}
		if !updated {
			break
		}
	}

	return Dist[end_node]
}

// 2699h Modify Graph Edge Weights
func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	G := make([][][2]int, n)

	for i, e := range edges {
		v, u := e[0], e[1]
		G[v], G[u] = append(G[v], [2]int{u, i}), append(G[u], [2]int{v, i})
	}

	var Dist, Prev []int
	SPF := func(src, dst int) {
		for i := range Dist {
			Dist[i] = math.MaxInt32
		}

		Dist[src] = 0
		Q := []int{src}

		var v int
		for len(Q) > 0 {
			v, Q = Q[0], Q[1:]
			if v != dst {
				for _, u := range G[v] {
					u, w := u[0], edges[u[1]][2]
					if w == -1 {
						continue
					}
					if Dist[v]+w < Dist[u] {
						Dist[u] = Dist[v] + w // relax edge
						Prev[u] = v

						Q = append(Q, u)
					}
				}
			}
		}
	}

	Dist, Prev = make([]int, n), make([]int, n)
	SPF(source, destination)

	if Dist[destination] < target {
		return [][]int{}
	}

	mTarget := Dist[destination] == target
	for i := range edges {
		if edges[i][2] > 0 {
			continue
		}

		if mTarget {
			edges[i][2] = math.MaxInt32
		} else {
			edges[i][2] = 1
			Dist, Prev = make([]int, n), make([]int, n)
			SPF(source, destination)
			if Dist[destination] <= target {
				mTarget = true
				edges[i][2] += target - Dist[destination]
			}
		}
	}
	if mTarget {
		return edges
	}

	return [][]int{}
}
