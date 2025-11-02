package application

import (
	"math/rand"

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
