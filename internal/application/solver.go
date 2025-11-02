package application

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"

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
