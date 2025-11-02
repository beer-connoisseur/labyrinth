package application

import (
	"math/rand"
	"time"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type DFSGen struct {
	visited [][]bool
	r       *rand.Rand
}

func NewDFSGen() *DFSGen {
	return &DFSGen{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *DFSGen) Generate(width, height int) (*domain.Maze, error) {
	maze, err := domain.NewMaze(width, height)
	if err != nil {
		return nil, err
	}
	g.visited = make([][]bool, width+2)
	for i := range g.visited {
		g.visited[i] = make([]bool, height+2)
	}

	g.dfs(domain.GenerationStartPointX, domain.GenerationStartPointY, maze)

	return maze, nil
}

func (g *DFSGen) dfs(x, y int, maze *domain.Maze) {
	maze.Cells[x][y] = getRandomSurface(g.r)
	g.visited[x][y] = true

	directions := make([][2]int, len(domain.Directions))
	copy(directions, domain.Directions)
	g.r.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})
	point := domain.Point{X: x, Y: y}
	for _, neighbor := range point.GetNeighbors(directions) {
		overNeighborX := 2*neighbor.X - x
		overNeighborY := 2*neighbor.Y - y
		if maze.IsInsideTheMaze(overNeighborX, overNeighborY) {
			if !g.visited[overNeighborX][overNeighborY] {
				maze.Cells[neighbor.X][neighbor.Y] = getRandomSurface(g.r)
				g.dfs(overNeighborX, overNeighborY, maze)
			} else if g.r.Float64() < domain.CycleProbability {
				maze.Cells[neighbor.X][neighbor.Y] = getRandomSurface(g.r)
			}
		}
	}
}
