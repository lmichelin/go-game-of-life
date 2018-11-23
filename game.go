package main

type game struct {
	size    int
	grids   [2][][]int
	gridPtr int
}

func (g *game) init(N int) {
	g.size = N

	g.grids[0] = make([][]int, N+2)
	g.grids[1] = make([][]int, N+2)
	for i := 0; i < N+2; i++ {
		g.grids[0][i] = make([]int, N+2)
		g.grids[1][i] = make([]int, N+2)
	}
}

func (g *game) get(x, y int) int {
	return g.grids[g.gridPtr][x+1][y+1]
}

func (g *game) set(x, y, v int) {
	g.grids[g.gridPtr][x+1][y+1] = v
}

func (g *game) setPtr(p, x, y, v int) {
	g.grids[p][x+1][y+1] = v
}

func (g *game) rule(x, y int) int {
	c := -g.get(x, y)
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			c += g.get(i, j)
		}
	}

	if c == 3 {
		return 1
	} else if c == 2 {
		return g.get(x, y)
	}
	return 0
}

func (g *game) ruleColumn(ch chan int, x int) {
	for y := 0; y < g.size; y++ {
		g.setPtr(1-g.gridPtr, x, y, g.rule(x, y))
	}

	ch <- x
}

func (g *game) run() {
	ch := make(chan int)

	for x := 0; x < g.size; x++ {
		go g.ruleColumn(ch, x)
	}

	for x := 0; x < g.size; x++ {
		<-ch
	}

	g.gridPtr = 1 - g.gridPtr
}

func main() {
	N := 50

	var g game
	g.init(N)

	// Initial state
	x := 1
	y := 48

	// Ship
	g.set(x+1, y+1, 1)
	g.set(x-1, y, 1)
	g.set(x+1, y, 1)
	g.set(x, y-1, 1)
	g.set(x+1, y-1, 1)

	run(&g, 10)
}
