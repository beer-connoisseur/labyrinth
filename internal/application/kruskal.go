package application

import (
	"math/rand"
	"time"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type KruskalGen struct {
	disjointSet map[domain.Point]domain.Point
	r           *rand.Rand
}

func NewKruskalGen() *KruskalGen {
	return &KruskalGen{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *KruskalGen) find(point domain.Point) domain.Point {
	if g.disjointSet[point] != point {
		g.disjointSet[point] = g.find(g.disjointSet[point])
	}
	return g.disjointSet[point]
}

func (g *KruskalGen) union(a, b domain.Point) {
	parentA, parentB := g.find(a), g.find(b)
	if parentA != parentB {
		g.disjointSet[parentB] = parentA
	}
}

// Generate maze with kruskal algorithm.
// We start with a sparse field, and using the system of disjoint sets, we cut through walls:
// if two cells are in different sets, we destroy the wall.
// With a certain probability, CycleProbability, even if two cells are in the same set,
// we'll still break the wall between them to create a cycle.
func (g *KruskalGen) Generate(width, height int) (*domain.Maze, error) {
	maze, err := domain.NewMaze(width, height)
	if err != nil {
		return nil, err
	}

	var walls []domain.Point
	g.disjointSet = make(map[domain.Point]domain.Point)
	for x := 1; x <= maze.Width; x = x + 2 {
		for y := 1; y <= maze.Height; y = y + 2 {
			maze.Cells[x][y] = getRandomSurface(g.r)
			walls = append(walls, getWalls(x, y, maze)...)
			g.disjointSet[domain.Point{X: x, Y: y}] = domain.Point{X: x, Y: y}
		}
	}

	for len(walls) > 0 {
		var wall domain.Point
		walls, wall = removeRandomElement(g.r, walls)
		surfaces := getSurfaces(wall.X, wall.Y, maze)
		if g.find(surfaces[0]) != g.find(surfaces[1]) {
			g.union(surfaces[0], surfaces[1])
			maze.Cells[wall.X][wall.Y] = getRandomSurface(g.r)
		} else if g.r.Float64() < domain.CycleProbability {
			maze.Cells[wall.X][wall.Y] = getRandomSurface(g.r)
		}
	}

	return maze, nil
}
