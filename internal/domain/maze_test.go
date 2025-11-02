package domain

import (
	"testing"
)

func TestMaze_DrawMazeASCII(t *testing.T) {
	tests := []struct {
		name string
		maze *Maze
		want string
	}{
		{
			name: "maze",
			maze: &Maze{
				Cells: [][]CellType{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
				Width:  1,
				Height: 1,
			},
			want: "###\n# #\n###\n",
		},
		{
			name: "maze-2",
			maze: &Maze{
				Cells: [][]CellType{
					{1, 1, 1},
					{1, 3, 1},
					{1, 1, 1},
				},
				Width:  1,
				Height: 1,
			},
			want: "###\n#$#\n###\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maze.DrawMazeASCII(); got != tt.want {
				t.Errorf("DrawMazeASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaze_DrawMazeUnicode(t *testing.T) {
	tests := []struct {
		name string
		maze *Maze
		want string
	}{
		{
			name: "maze",
			maze: &Maze{
				Cells: [][]CellType{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
				Width:  1,
				Height: 1,
			},
			want: "🧱🧱🧱\n🧱⬛🧱\n🧱🧱🧱\n",
		},
		{
			name: "maze-2",
			maze: &Maze{
				Cells: [][]CellType{
					{1, 1, 1},
					{1, 3, 1},
					{1, 1, 1},
				},
				Width:  1,
				Height: 1,
			},
			want: "🧱🧱🧱\n🧱🪙🧱\n🧱🧱🧱\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maze.DrawMazeUnicode(); got != tt.want {
				t.Errorf("DrawMazeUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaze_IsInsideTheMaze(t *testing.T) {
	tests := []struct {
		name string
		maze *Maze
		x, y int
		want bool
	}{
		{
			name: "inside",
			maze: &Maze{
				Width:  3,
				Height: 3,
			},
			x:    1,
			y:    1,
			want: true,
		},
		{
			name: "outside negative",
			maze: &Maze{
				Width:  3,
				Height: 3,
			},
			x:    -1,
			y:    0,
			want: false,
		},
		{
			name: "outside overflow",
			maze: &Maze{
				Width:  3,
				Height: 3,
			},
			x:    4,
			y:    3,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.maze.IsInsideTheMaze(tt.x, tt.y); got != tt.want {
				t.Errorf("IsInsideTheMaze(%d,%d) = %v, want %v", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestNewMaze(t *testing.T) {
	tests := []struct {
		name    string
		width   int
		height  int
		wantErr bool
	}{
		{"valid 3x3", 3, 3, false},
		{"zero width", 0, 3, true},
		{"zero height", 3, 0, true},
		{"negative size", -1, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMaze(tt.width, tt.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMaze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got.Width != tt.width || got.Height != tt.height {
					t.Errorf("NewMaze() size mismatch: got %dx%d, want %dx%d", got.Width, got.Height, tt.width, tt.height)
				}
			}
		})
	}
}
