package application

import (
	"errors"
	"math"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type Edge struct {
	From, To domain.Point
	Weight   int
}

type BellmanFordNode struct {
	domain.Point
	Dist   int
	Parent *BellmanFordNode
}

func (n *BellmanFordNode) GetCoords() (int, int) {
	return n.X, n.Y
}

func (n *BellmanFordNode) GetParent() PathNode {
	if n.Parent == nil {
		return nil
	}
	return n.Parent
}

type BellmanFordSolver struct{}

func NewBellmanFordSolver() *BellmanFordSolver {
	return &BellmanFordSolver{}
}

// Solve maze with Bellman-Ford algorithm.
// Each vertex stores the minimal distance from the start point and a reference to its parent vertex.
// The algorithm iteratively relaxes all edges in the graph V-1 times (where V is the number of vertices),
// updating distances and parents if a shorter path is found.
// The process continues until distances stabilize, guaranteeing the shortest path
// to all reachable vertices, including the end point if it exists.
func (s *BellmanFordSolver) Solve(start, end domain.Point, maze *domain.Maze) (*domain.Maze, error) {
	if maze.Cells[start.X][start.Y] == domain.Wall || maze.Cells[end.X][end.Y] == domain.Wall {
		return nil, errors.New("invalid start/end points")
	}
	maze.Cells[start.X][start.Y] = domain.Start
	maze.Cells[end.X][end.Y] = domain.End

	nodes := make(map[domain.Point]*BellmanFordNode)

	edges := buildEdges(maze)

	for _, e := range edges {
		if _, ok := nodes[e.From]; !ok {
			nodes[e.From] = &BellmanFordNode{Point: e.From, Dist: math.MaxInt}
		}
		if _, ok := nodes[e.To]; !ok {
			nodes[e.To] = &BellmanFordNode{Point: e.To, Dist: math.MaxInt}
		}
	}
	nodes[start].Dist = 0

	for i := 0; i < len(nodes)-1; i++ {
		changed := false
		for _, e := range edges {
			if nodes[e.From].Dist == math.MaxInt {
				continue
			}

			newCost := nodes[e.From].Dist + e.Weight
			if newCost < nodes[e.To].Dist {
				nodes[e.To].Dist = newCost
				nodes[e.To].Parent = nodes[e.From]
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	if nodes[end].Dist == math.MaxInt {
		return nil, errors.New("there is no solution")
	}
	return findPath(nodes[end], maze), nil
}

func buildEdges(maze *domain.Maze) []Edge {
	var edges []Edge

	for x := 1; x <= maze.Width; x++ {
		for y := 1; y <= maze.Height; y++ {
			if maze.Cells[x][y] == domain.Wall {
				continue
			}
			from := domain.Point{X: x, Y: y}
			for _, neighbor := range from.GetNeighbors(domain.Directions) {
				if !maze.IsInsideTheMaze(neighbor.X, neighbor.Y) || maze.Cells[neighbor.X][neighbor.Y] == domain.Wall {
					continue
				}
				to := domain.Point{X: neighbor.X, Y: neighbor.Y}
				weight := domain.SurfacesCost[maze.Cells[neighbor.X][neighbor.Y]]
				edges = append(edges, Edge{
					From:   from,
					To:     to,
					Weight: weight,
				})
			}
		}
	}
	return edges
}
