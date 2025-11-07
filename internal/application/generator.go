package application

import (
	"math/rand"
	"time"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

type Generator interface {
	Generate(width, height int) (*domain.Maze, error)
}

func getRandomSurface(r *rand.Rand) domain.CellType {
	probability := r.Float64()

	if probability < domain.RockProbability {
		return domain.Rock
	} else if probability < domain.TreeProbability+domain.RockProbability {
		return domain.Tree
	} else if probability < domain.CoinProbability+domain.TreeProbability+domain.RockProbability {
		return domain.Coin
	} else {
		return domain.Space
	}
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

func checkSeed(seed ...int64) rand.Source {
	var src rand.Source
	if len(seed) > 0 {
		src = rand.NewSource(seed[0])
	} else {
		src = rand.NewSource(time.Now().UnixNano())
	}
	return src
}
