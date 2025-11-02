package application

import (
	"errors"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type AStarNode struct {
	domain.Point
	Parent *AStarNode
	F, G   int
}

func (n *AStarNode) GetCoords() (int, int) {
	return n.X, n.Y
}

func (n *AStarNode) GetParent() PathNode {
	if n.Parent == nil {
		return nil
	}
	return n.Parent
}

type AStarSolver struct{}

func NewAStarSolver() *AStarSolver {
	return &AStarSolver{}
}

func (s *AStarSolver) Solve(start, end domain.Point, maze *domain.Maze) (*domain.Maze, error) {
	if maze.Cells[start.X][start.Y] == domain.Wall || maze.Cells[end.X][end.Y] == domain.Wall {
		return nil, errors.New("invalid start/end points")
	}
	maze.Cells[start.X][start.Y] = domain.Start
	maze.Cells[end.X][end.Y] = domain.End

	startNode := AStarNode{
		Point: start,
		G:     0,
	}
	startNode.F = startNode.G + manhattanDistance(start, end)

	vertexes := NewPriorityQueue(func(a, b *AStarNode) bool {
		return a.F < b.F
	})
	vertexes.PushNode(&startNode)

	visited := make(map[domain.Point]*AStarNode)

	for vertexes.Len() > 0 {
		current := vertexes.PopNode()
		if current.Point == end {
			return findPath(current, maze), nil
		}
		visited[current.Point] = current

		for _, neighbor := range current.GetNeighbors(domain.Directions) {
			if !maze.IsInsideTheMaze(neighbor.X, neighbor.Y) || maze.Cells[neighbor.X][neighbor.Y] == domain.Wall {
				continue
			}

			tentativeScore := current.G + domain.SurfacesCost[maze.Cells[neighbor.X][neighbor.Y]]
			if v, ok := visited[neighbor]; ok && tentativeScore >= v.G {
				continue
			}

			v := AStarNode{
				Point:  neighbor,
				Parent: current,
				G:      tentativeScore,
			}
			v.F = v.G + manhattanDistance(neighbor, end)
			vertexes.PushNode(&v)
		}
	}

	return nil, errors.New("there is no solution")
}

func manhattanDistance(current, end domain.Point) int {
	return int(math.Abs(float64(current.X-end.X)) + math.Abs(float64(current.Y-end.Y)))
}
