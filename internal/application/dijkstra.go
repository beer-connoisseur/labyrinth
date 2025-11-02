package application

import (
	"errors"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type DijkstraNode struct {
	domain.Point
	Dist   int
	Parent *DijkstraNode
}

func (n *DijkstraNode) GetCoords() (int, int) {
	return n.X, n.Y
}

func (n *DijkstraNode) GetParent() PathNode {
	if n.Parent == nil {
		return nil
	}
	return n.Parent
}

type DijkstraSolver struct{}

func NewDijkstraSolver() *DijkstraSolver {
	return &DijkstraSolver{}
}

// Solve maze with Dijkstra algorithm.
// Each vertex stores the minimal distance from the start point.
// Vertices are processed using a priority queue sorted by this distance.
// The algorithm updates distances and parents until the end point is reached or no path exists.
func (s *DijkstraSolver) Solve(start, end domain.Point, maze *domain.Maze) (*domain.Maze, error) {
	if maze.Cells[start.X][start.Y] == domain.Wall || maze.Cells[end.X][end.Y] == domain.Wall {
		return nil, errors.New("invalid start/end points")
	}
	maze.Cells[start.X][start.Y] = domain.Start
	maze.Cells[end.X][end.Y] = domain.End

	d := make([][]int, maze.Width+2)
	for i := range d {
		d[i] = make([]int, maze.Height+2)
		for j := range d[i] {
			d[i][j] = math.MaxInt
		}
	}

	d[start.X][start.Y] = 0
	startNode := DijkstraNode{
		Point: start,
		Dist:  0,
	}
	vertexes := NewPriorityQueue(func(a, b *DijkstraNode) bool {
		return a.Dist < b.Dist
	})
	vertexes.PushNode(&startNode)

	for vertexes.Len() > 0 {
		current := vertexes.PopNode()
		if current.Point == end {
			return findPath(current, maze), nil
		}

		for _, neighbor := range current.GetNeighbors(domain.Directions) {
			if !maze.IsInsideTheMaze(neighbor.X, neighbor.Y) || maze.Cells[neighbor.X][neighbor.Y] == domain.Wall {
				continue
			}

			newCost := d[current.X][current.Y] + domain.SurfacesCost[maze.Cells[neighbor.X][neighbor.Y]]
			if d[neighbor.X][neighbor.Y] > newCost {
				d[neighbor.X][neighbor.Y] = newCost

				v := DijkstraNode{
					Point:  neighbor,
					Parent: current,
					Dist:   newCost,
				}
				vertexes.PushNode(&v)
			}
		}
	}
	return nil, errors.New("there is no solution")
}
