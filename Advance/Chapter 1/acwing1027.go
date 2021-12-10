package main

import "fmt"

const maxn int = 20

var (
	g [maxn][maxn]int
	f [2 * maxn][maxn][maxn]int
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var n, r, c, w int
	fmt.Scanf("%d", &n)
	for fmt.Scanf("%d%d%d", &r, &c, &w); r > 0 || c > 0 || w > 0; fmt.Scanf("%d%d%d", &r, &c, &w) {
		g[r][c] = w
	}
	for k := 2; k <= 2*n; k++ {
		for i1 := 1; i1 <= n && i1 < k; i1++ {
			for i2 := 1; i2 <= n && i2 < k; i2++ {
				j1, j2 := k-i1, k-i2
				t := g[i1][j1]
				if i1 != i2 {
					t += g[i2][j2]
				}
				var x *int = &f[k][i1][i2]
				*x = max(*x, f[k-1][i1-1][i2-1]+t)
				*x = max(*x, f[k-1][i1-1][i2]+t)
				*x = max(*x, f[k-1][i1][i2-1]+t)
				*x = max(*x, f[k-1][i1][i2]+t)
			}
		}
	}
	fmt.Println(f[2*n][n][n])
}
