package Meta

// 1514m Path with Maximum Probability
func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	// Bellman-Ford

	Dist := make([]float64, n)
	Dist[start_node] = 1

	for range n - 1 {
		for i, e := range edges {
			p := succProb[i]
			v, u := e[0], e[1]

			// Relaxing: 2-way as for undirected graph
			if Dist[v]*p > Dist[u] {
				Dist[u] = Dist[v] * p
			}
			if Dist[u]*p > Dist[v] {
				Dist[v] = Dist[u] * p
			}
		}
	}

	return Dist[end_node]
}
