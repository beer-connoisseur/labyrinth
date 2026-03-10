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

func Test_getSurfaces(t *testing.T) {
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

func Test_getWalls(t *testing.T) {
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
