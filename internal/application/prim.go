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

// Generate maze with Prim algorithm.
// We have an array of walls that are available for cutting, and we randomly select a wall.
// Once the wall is cut, new walls are added to the array that the new cell is adjacent to.
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

		// We check that there is only one passable cell around the wall we want to break through.
		// Or if there are only two opposite passable walls around the wall,
		// then with a certain probability, CycleProbability, we will still cut through the wall to create a cycle
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
