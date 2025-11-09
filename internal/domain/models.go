package domain

type CellType int

const (
	Space CellType = iota
	Wall
	Path
	Coin
	Tree
	Rock
	Start
	End
)

// When setting the starting point of the generation, keep in mind that the maze field has 1-indexing
const (
	GenerationStartPointX = 1
	GenerationStartPointY = 1
)

const (
	CycleProbability = 0.1
	CoinProbability  = 0.4
	TreeProbability  = 0.2
	RockProbability  = 0.1
)

const FilePermissions = 0644

var Directions = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type Point struct {
	X, Y int
}

type Maze struct {
	Cells         [][]CellType
	Width, Height int
}
