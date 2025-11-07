package application

import (
	"reflect"
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func TestNewPrimGen(t *testing.T) {
	t.Run("creates non-nil generator", func(t *testing.T) {
		gen := NewPrimGen()
		if gen == nil {
			t.Fatal("NewDFSGen() returned nil")
		}
		if gen.r == nil {
			t.Error("NewDFSGen() did not initialize rand")
		}
	})
}

func TestPrimGen_Generate(t *testing.T) {
	tests := []struct {
		name    string
		width   int
		height  int
		wantErr bool
	}{
		{
			name:    "valid maze 5x5",
			width:   5,
			height:  5,
			wantErr: false,
		},
		{
			name:    "valid maze 3x5 odd args",
			width:   3,
			height:  5,
			wantErr: false,
		},
		{
			name:    "valid maze 5x3 odd args",
			width:   5,
			height:  3,
			wantErr: false,
		},
		{
			name:    "valid maze 3x5 even args",
			width:   2,
			height:  4,
			wantErr: false,
		},
		{
			name:    "valid maze 5x3 even args",
			width:   4,
			height:  2,
			wantErr: false,
		},
		{
			name:    "invalid zero width",
			width:   0,
			height:  5,
			wantErr: true,
		},
		{
			name:    "invalid negative height",
			width:   5,
			height:  -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewPrimGen()
			got, err := g.Generate(tt.width, tt.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				if got == nil {
					t.Fatal("Generate() returned nil maze on success")
				}
				if tt.width%2 == 0 {
					tt.width++
				}
				if got.Width != tt.width {
					t.Errorf("wrong width: got %d, want %d", got.Width, tt.width)
				}
				if tt.height%2 == 0 {
					tt.height++
				}
				if got.Height != tt.height {
					t.Errorf("wrong height: got %d, want %d", got.Height, tt.height)
				}
				if len(got.Cells) == 0 {
					t.Errorf("maze has no cells")
				}
			}
		})
	}
}

func TestPrimGen_deterministic(t *testing.T) {
	tests := []struct {
		name        string
		width       int
		height      int
		seeds       []int64
		expectEqual bool
	}{
		{
			name:        "same seed",
			width:       10,
			height:      10,
			seeds:       []int64{42, 42},
			expectEqual: true,
		},
		{
			name:        "different seeds",
			width:       10,
			height:      10,
			seeds:       []int64{1, 2},
			expectEqual: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen1 := NewPrimGen(tt.seeds[0])
			gen2 := NewPrimGen(tt.seeds[1])

			maze1, err1 := gen1.Generate(tt.width, tt.height)
			maze2, err2 := gen2.Generate(tt.width, tt.height)

			if err1 != nil || err2 != nil {
				t.Fatalf("unexpected error: %v / %v", err1, err2)
			}

			equal := reflect.DeepEqual(maze1.Cells, maze2.Cells)
			if tt.expectEqual && !equal {
				t.Errorf("expected mazes to be equal for seeds %v", tt.seeds)
			}
			if !tt.expectEqual && equal {
				t.Errorf("expected mazes to differ for seeds %v", tt.seeds)
			}
		})
	}
}

func Test_isDiagonalConnections(t *testing.T) {
	tests := []struct {
		name  string
		maze  *domain.Maze
		space domain.Point
		wall  domain.Point
		want  bool
	}{
		{
			name: "diagonal connection",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 0, 1, 0, 1},
					{1, 1, 1, 1, 1},
					{1, 0, 0, 0, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  3,
				Height: 3,
			},
			space: domain.Point{X: 3, Y: 2},
			wall:  domain.Point{X: 2, Y: 2},
			want:  true,
		},
		{
			name: "diagonal connection-2",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 0, 1, 1, 1},
					{1, 0, 1, 3, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  3,
				Height: 3,
			},
			space: domain.Point{X: 2, Y: 1},
			wall:  domain.Point{X: 2, Y: 2},
			want:  true,
		},
		{
			name: "no diagonal connection",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 0, 1, 1, 1},
					{1, 0, 1, 1, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  3,
				Height: 3,
			},
			space: domain.Point{X: 2, Y: 1},
			wall:  domain.Point{X: 2, Y: 2},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isDiagonalConnections(tt.space, tt.wall, tt.maze)
			if got != tt.want {
				t.Errorf("isDiagonalConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPassable(t *testing.T) {
	tests := []struct {
		name string
		cell domain.CellType
		want bool
	}{
		{"wall", domain.Wall, false},
		{"coin", domain.Coin, true},
		{"tree", domain.Tree, true},
		{"rock", domain.Rock, true},
		{"space", domain.Space, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPassable(tt.cell)
			if got != tt.want {
				t.Errorf("isPassable() = %v, want %v", got, tt.want)
			}
		})
	}
}
