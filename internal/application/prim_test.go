package application

import (
	"math/rand"
	"reflect"
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func pointsEqualIgnoringOrder(a, b []domain.Point) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[domain.Point]int)
	for _, p := range a {
		counts[p]++
	}
	for _, p := range b {
		if counts[p] == 0 {
			return false
		}
		counts[p]--
	}

	return true
}

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
				if got.Width != tt.width {
					t.Errorf("wrong width: got %d, want %d", got.Width, tt.width)
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

func TestPrimGen_getSurfaces(t *testing.T) {
	tests := []struct {
		name  string
		maze  *domain.Maze
		point domain.Point
		want  []domain.Point
	}{
		{
			name: "four different surfaces",
			maze: &domain.Maze{
				// the x and y coordinates are inverted because I'm creating a field like this
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 0, 3, 4, 1},
					{1, 0, 0, 4, 1},
					{1, 0, 5, 0, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  3,
				Height: 3,
			},
			point: domain.Point{X: 2, Y: 2},
			want: []domain.Point{
				{2, 1},
				{1, 2},
				{3, 2},
				{2, 3},
			},
		},
		{
			name: "one way",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 0, 1, 4, 1},
					{1, 0, 0, 1, 1},
					{1, 0, 1, 0, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  3,
				Height: 3,
			},
			point: domain.Point{X: 2, Y: 2},
			want: []domain.Point{
				{2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSurfaces(tt.point.X, tt.point.Y, tt.maze); !pointsEqualIgnoringOrder(got, tt.want) {
				t.Errorf("getSurfaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrimGen_getWalls(t *testing.T) {
	tests := []struct {
		name  string
		maze  *domain.Maze
		point domain.Point
		want  []domain.Point
	}{
		{
			name: "four walls",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 0, 1, 4, 1},
					{1, 1, 0, 1, 1},
					{1, 0, 1, 0, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  3,
				Height: 3,
			},
			point: domain.Point{X: 2, Y: 2},
			want: []domain.Point{
				{2, 1},
				{1, 2},
				{3, 2},
				{2, 3},
			},
		},
		{
			name: "one wall",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 0, 0, 4, 1},
					{1, 1, 0, 0, 1},
					{1, 0, 0, 0, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  3,
				Height: 3,
			},
			point: domain.Point{X: 2, Y: 2},
			want: []domain.Point{
				{2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWalls(tt.point.X, tt.point.Y, tt.maze); !pointsEqualIgnoringOrder(got, tt.want) {
				t.Errorf("getWalls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNeighborCells(t *testing.T) {
	tests := []struct {
		name     string
		maze     *domain.Maze
		point    domain.Point
		cellType domain.CellType
		want     []domain.Point
	}{
		{
			name: "only spaces",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 0, 0, 4, 1},
					{1, 3, 0, 0, 1},
					{1, 0, 1, 0, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  3,
				Height: 3,
			},
			point:    domain.Point{X: 2, Y: 2},
			cellType: domain.Space,
			want: []domain.Point{
				{1, 2},
				{2, 3},
			},
		},
		{
			name: "only coins",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 0, 3, 4, 1},
					{1, 1, 0, 0, 1},
					{1, 0, 0, 0, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  3,
				Height: 3,
			},
			point:    domain.Point{X: 2, Y: 2},
			cellType: domain.Coin,
			want: []domain.Point{
				{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNeighborCells(tt.point.X, tt.point.Y, tt.maze, tt.cellType); !pointsEqualIgnoringOrder(got, tt.want) {
				t.Errorf("getWalls() = %v, want %v", got, tt.want)
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

func Test_removeRandomElement(t *testing.T) {
	type args[T any] struct {
		r     *rand.Rand
		slice []T
	}
	type testCase[T any] struct {
		name  string
		args  args[T]
		want  []T
		want1 T
	}

	intTests := []testCase[int]{
		{
			name:  "four ints",
			args:  args[int]{r: rand.New(rand.NewSource(42)), slice: []int{1, 2, 3, 4}},
			want:  []int{1, 4, 3},
			want1: 2,
		},
		{
			name:  "three ints",
			args:  args[int]{r: rand.New(rand.NewSource(1)), slice: []int{10, 20, 30}},
			want:  []int{10, 20},
			want1: 30,
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeRandomElement(tt.args.r, tt.args.slice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeRandomElement() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("removeRandomElement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}

	strTests := []testCase[string]{
		{
			name:  "three strings",
			args:  args[string]{r: rand.New(rand.NewSource(42)), slice: []string{"a", "b", "c"}},
			want:  []string{"a", "b"},
			want1: "c",
		},
		{
			name:  "three strings-2",
			args:  args[string]{r: rand.New(rand.NewSource(1)), slice: []string{"x", "y", "z"}},
			want:  []string{"x", "y"},
			want1: "z",
		},
	}

	for _, tt := range strTests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeRandomElement(tt.args.r, tt.args.slice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeRandomElement() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("removeRandomElement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
