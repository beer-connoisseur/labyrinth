package application

import (
	"errors"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type Solver interface {
	Solve(start, end domain.Point, m *domain.Maze) (*domain.Maze, error)
}

type PathNode interface {
	GetParent() PathNode
	GetCoords() (int, int)
}

func findPath[T PathNode](current T, maze *domain.Maze) *domain.Maze {
	for node := current.GetParent(); node.GetParent() != nil; node = node.GetParent() {
		x, y := node.GetCoords()
		maze.Cells[x][y] = domain.Path
	}
	return maze
}

func checkAndSetStartEnd(start, end domain.Point, maze *domain.Maze) error {
	if maze.Cells[start.X][start.Y] == domain.Wall || maze.Cells[end.X][end.Y] == domain.Wall {
		return errors.New("invalid start/end points")
	}
	maze.Cells[start.X][start.Y] = domain.Start
	maze.Cells[end.X][end.Y] = domain.End

	return nil
}
