package main

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	type Pair struct {
		r, c int
	}

	queue := make([]Pair, 0)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				dist[r][c] = 0
				queue = append(queue, Pair{r, c})
			}
		}
	}

	dirs := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	// Multi-source BFS
	head := 0
	for head < len(queue) {
		cur := queue[head]
		head++

		for _, d := range dirs {
			nr := cur.r + d[0]
			nc := cur.c + d[1]

			if nr < 0 || nr >= n || nc < 0 || nc >= n {
				continue
			}

			if dist[nr][nc] != -1 {
				continue
			}

			dist[nr][nc] = dist[cur.r][cur.c] + 1
			queue = append(queue, Pair{nr, nc})
		}
	}

	canReach := func(limit int) bool {
		if dist[0][0] < limit {
			return false
		}

		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		q := []Pair{{0, 0}}
		visited[0][0] = true

		head := 0

		for head < len(q) {
			cur := q[head]
			head++

			if cur.r == n-1 && cur.c == n-1 {
				return true
			}

			for _, d := range dirs {
				nr := cur.r + d[0]
				nc := cur.c + d[1]

				if nr < 0 || nr >= n || nc < 0 || nc >= n {
					continue
				}

				if visited[nr][nc] {
					continue
				}

				if dist[nr][nc] < limit {
					continue
				}

				visited[nr][nc] = true
				q = append(q, Pair{nr, nc})
			}
		}

		return false
	}

	hi := 2 * n
	lo := 0
	ans := 0

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if canReach(mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return ans
}