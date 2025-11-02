package domain

import (
	"errors"
	"strings"
)

func NewMaze(width, height int) (*Maze, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("invalid maze attributes")
	}
	// we need odd lengths for the generation algorithm to work correctly
	if width%2 == 0 {
		width = width + 1
	}
	if height%2 == 0 {
		height = height + 1
	}

	cells := make([][]CellType, width+2)
	for i := range cells {
		cells[i] = make([]CellType, height+2)
		for j := range cells[i] {
			cells[i][j] = Wall
		}
	}

	return &Maze{
		Cells:  cells,
		Width:  width,
		Height: height,
	}, nil
}

func (m *Maze) IsInsideTheMaze(x, y int) bool {
	return (x > 0 && x < m.Width+1) && (y > 0 && y < m.Height+1)
}

func (m *Maze) DrawMazeUnicode() string {
	var sb strings.Builder
	for y := 0; y < m.Height+2; y++ {
		for x := 0; x < m.Width+2; x++ {
			cell := m.Cells[x][y]
			switch cell {
			case Space:
				sb.WriteString("\u2B1B")
			case Wall:
				sb.WriteString("\U0001F9F1")
			case Path:
				sb.WriteString("\u2B55")
			case Coin:
				sb.WriteString("\U0001FA99")
			case Tree:
				sb.WriteString("\U0001F335")
			case Rock:
				sb.WriteString("\U0001FAA8")
			case Start:
				sb.WriteString("\U0001F3E1")
			case End:
				sb.WriteString("\U0001F3C1")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (m *Maze) DrawMazeASCII() string {
	var sb strings.Builder
	for y := 0; y < m.Height+2; y++ {
		for x := 0; x < m.Width+2; x++ {
			cell := m.Cells[x][y]
			switch cell {
			case Space:
				sb.WriteString(" ")
			case Wall:
				sb.WriteString("#")
			case Path:
				sb.WriteString(".")
			case Coin:
				sb.WriteString("$")
			case Tree:
				sb.WriteString("+")
			case Rock:
				sb.WriteString("@")
			case Start:
				sb.WriteString("O")
			case End:
				sb.WriteString("X")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
