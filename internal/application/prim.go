package application

import (
	"math/rand"
	"time"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type PrimGen struct {
	r *rand.Rand
}

func NewPrimGen() *PrimGen {
	return &PrimGen{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *PrimGen) Generate(width, height int) (*domain.Maze, error) {
	maze, err := domain.NewMaze(width, height)
	if err != nil {
		return nil, err
	}

	maze.Cells[domain.GenerationStartPointX][domain.GenerationStartPointY] = getRandomSurface(g.r)
	walls := getWalls(domain.GenerationStartPointX, domain.GenerationStartPointY, maze)

	for len(walls) > 0 {
		var wall domain.Point
		walls, wall = removeRandomElement(g.r, walls)
		surfaces := getSurfaces(wall.X, wall.Y, maze)

		if len(surfaces) == 1 {
			if !isDiagonalConnections(surfaces[0], wall, maze) {
				maze.Cells[wall.X][wall.Y] = getRandomSurface(g.r)
				newWalls := getWalls(wall.X, wall.Y, maze)
				walls = append(walls, newWalls...)
			}
		} else if len(surfaces) == 2 &&
			(surfaces[0].X == surfaces[1].X || surfaces[0].Y == surfaces[1].Y) &&
			g.r.Float64() < domain.CycleProbability {
			maze.Cells[wall.X][wall.Y] = getRandomSurface(g.r)
		}
	}

	return maze, nil
}

func getNeighborCells(x, y int, maze *domain.Maze, cellType domain.CellType) []domain.Point {
	var cells []domain.Point
	point := domain.Point{X: x, Y: y}
	for _, neighbor := range point.GetNeighbors(domain.Directions) {
		if maze.IsInsideTheMaze(neighbor.X, neighbor.Y) && maze.Cells[neighbor.X][neighbor.Y] == cellType {
			cells = append(cells, domain.Point{X: neighbor.X, Y: neighbor.Y})
		}
	}
	return cells
}

func getWalls(x, y int, maze *domain.Maze) []domain.Point {
	return getNeighborCells(x, y, maze, domain.Wall)
}

func getSurfaces(x, y int, maze *domain.Maze) []domain.Point {
	surfaces := make([]domain.Point, 0)

	surfaces = append(surfaces, getNeighborCells(x, y, maze, domain.Space)...)
	surfaces = append(surfaces, getNeighborCells(x, y, maze, domain.Coin)...)
	surfaces = append(surfaces, getNeighborCells(x, y, maze, domain.Tree)...)
	surfaces = append(surfaces, getNeighborCells(x, y, maze, domain.Rock)...)
	return surfaces
}

func removeRandomElement[T any](r *rand.Rand, slice []T) ([]T, T) {
	if len(slice) == 0 {
		var zero T
		return slice, zero
	}

	randomIndex := r.Intn(len(slice))
	element := slice[randomIndex]

	slice[randomIndex] = slice[len(slice)-1]
	newSlice := slice[:len(slice)-1]

	return newSlice, element
}

// isDiagonalConnections checks for diagonal wall connections to prevent this type of walls
// ⬛🧱⬛
// 🧱⬛🧱
func isDiagonalConnections(space, wall domain.Point, maze *domain.Maze) bool {
	direction := domain.Point{X: wall.X - space.X, Y: wall.Y - space.Y}

	p1 := domain.Point{X: wall.X - direction.Y + direction.X, Y: wall.Y + direction.Y - direction.X}
	p2 := domain.Point{X: wall.X + direction.Y + direction.X, Y: wall.Y + direction.Y + direction.X}

	return isPassable(maze.Cells[p1.X][p1.Y]) || isPassable(maze.Cells[p2.X][p2.Y])
}

func isPassable(cell domain.CellType) bool {
	return cell == domain.Space || cell == domain.Coin || cell == domain.Tree || cell == domain.Rock
}
