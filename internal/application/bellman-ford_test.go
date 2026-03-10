package application

import (
	"reflect"
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func edgesEqualIgnoringOrder(a, b []Edge) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[Edge]int)
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

func TestBellmanFordNode_GetCoords(t *testing.T) {
	tests := []struct {
		name  string
		point domain.Point
		wantX int
		wantY int
	}{
		{
			name:  "simple coords",
			point: domain.Point{X: 2, Y: 3},
			wantX: 2,
			wantY: 3,
		},
		{
			name:  "zero coords",
			point: domain.Point{X: 0, Y: 0},
			wantX: 0,
			wantY: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BellmanFordNode{
				Point: tt.point,
			}
			gotX, gotY := n.GetCoords()
			if gotX != tt.wantX {
				t.Errorf("GetCoords() got = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("GetCoords() got1 = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestBellmanFordNode_GetParent(t *testing.T) {
	parent := &DijkstraNode{Point: domain.Point{X: 1, Y: 2}}
	child := &DijkstraNode{Point: domain.Point{X: 3, Y: 4}, Parent: parent}

	tests := []struct {
		name string
		node *DijkstraNode
		want PathNode
	}{
		{
			name: "has parent",
			node: child,
			want: parent,
		},
		{
			name: "no parent",
			node: &DijkstraNode{Point: domain.Point{X: 0, Y: 0}, Parent: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.GetParent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBellmanFordSolver_Solve(t *testing.T) {
	type args struct {
		start domain.Point
		end   domain.Point
		maze  *domain.Maze
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Maze
		wantErr bool
	}{
		{
			name: "correctness",
			args: args{
				maze: &domain.Maze{
					//🧱🧱🧱🧱🧱🧱🧱
					//🧱🪙🧱⬛🌵🌵🧱
					//🧱⬛🧱🪙🧱🪙🧱
					//🧱🪨🧱🪙🧱🪙🧱
					//🧱🪨🧱🪙🧱🌵🧱
					//🧱🪙🪨⬛⬛⬛🧱
					//🧱🧱🧱🧱🧱🧱🧱
					// the x and y coordinates are inverted because I'm creating a field like this
					Cells: [][]domain.CellType{
						{1, 1, 1, 1, 1, 1, 1},
						{1, 3, 0, 5, 5, 3, 1},
						{1, 1, 1, 1, 1, 5, 1},
						{1, 0, 3, 3, 3, 0, 1},
						{1, 4, 1, 1, 1, 0, 1},
						{1, 4, 3, 3, 4, 0, 1},
						{1, 1, 1, 1, 1, 1, 1},
					},
					Width:  5,
					Height: 5,
				},
				start: domain.Point{X: 3, Y: 1},
				end:   domain.Point{X: 5, Y: 3},
			},
			//🧱🧱🧱🧱🧱🧱🧱
			//🧱🪙🧱🏡🌵🌵🧱
			//🧱⬛🧱⭕🧱🪙🧱
			//🧱🪨🧱⭕🧱🏁🧱
			//🧱🪨🧱⭕🧱⭕🧱
			//🧱🪙🪨⭕⭕⭕🧱
			//🧱🧱🧱🧱🧱🧱🧱
			want: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1, 1, 1},
					{1, 3, 0, 5, 5, 3, 1},
					{1, 1, 1, 1, 1, 5, 1},
					{1, 6, 2, 2, 2, 2, 1},
					{1, 4, 1, 1, 1, 2, 1},
					{1, 4, 3, 7, 2, 2, 1},
					{1, 1, 1, 1, 1, 1, 1},
				},
				Width:  5,
				Height: 5,
			},
			wantErr: false,
		},
		{
			name: "correctness coin cycle",
			args: args{
				maze: &domain.Maze{
					//🧱🧱🧱🧱🧱🧱🧱
					//🧱🪙🧱⬛🌵🌵🧱
					//🧱⬛🧱🪙🧱🪙🧱
					//🧱🪙🪙🪙🧱🪙🧱
					//🧱🪙🧱🪙🧱🌵🧱
					//🧱🪙🪙🪙⬛⬛🧱
					//🧱🧱🧱🧱🧱🧱🧱
					Cells: [][]domain.CellType{
						{1, 1, 1, 1, 1, 1, 1},
						{1, 3, 0, 3, 3, 3, 1},
						{1, 1, 1, 3, 1, 3, 1},
						{1, 0, 3, 3, 3, 3, 1},
						{1, 4, 1, 1, 1, 0, 1},
						{1, 4, 3, 3, 4, 0, 1},
						{1, 1, 1, 1, 1, 1, 1},
					},
					Width:  5,
					Height: 5,
				},
				start: domain.Point{X: 3, Y: 1},
				end:   domain.Point{X: 5, Y: 3},
			},
			//🧱🧱🧱🧱🧱🧱🧱
			//🧱🪙🧱🏡🌵🌵🧱
			//🧱⬛🧱⭕🧱🪙🧱
			//🧱🪙🪙⭕🧱🏁🧱
			//🧱🪙🧱⭕🧱⭕🧱
			//🧱🪙🪙⭕⭕⭕🧱
			//🧱🧱🧱🧱🧱🧱🧱
			want: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1, 1, 1},
					{1, 3, 0, 3, 3, 3, 1},
					{1, 1, 1, 3, 1, 3, 1},
					{1, 6, 2, 2, 2, 2, 1},
					{1, 4, 1, 1, 1, 2, 1},
					{1, 4, 3, 7, 2, 2, 1},
					{1, 1, 1, 1, 1, 1, 1},
				},
				Width:  5,
				Height: 5,
			},
			wantErr: false,
		},
		{
			name: "invalid start",
			args: args{
				maze: &domain.Maze{
					//🧱🧱🧱🧱🧱🧱🧱
					//🧱🪙🧱⬛🌵🌵🧱
					//🧱⬛🧱🪙🧱🪙🧱
					//🧱🪨🧱🪙🧱🪙🧱
					//🧱🪨🧱🪙🧱🌵🧱
					//🧱🪙🪨⬛⬛⬛🧱
					//🧱🧱🧱🧱🧱🧱🧱
					Cells: [][]domain.CellType{
						{1, 1, 1, 1, 1, 1, 1},
						{1, 3, 0, 5, 5, 3, 1},
						{1, 1, 1, 1, 1, 5, 1},
						{1, 0, 3, 3, 3, 0, 1},
						{1, 4, 1, 1, 1, 0, 1},
						{1, 4, 3, 3, 4, 0, 1},
						{1, 1, 1, 1, 1, 1, 1},
					},
					Width:  5,
					Height: 5,
				},
				start: domain.Point{X: 2, Y: 1},
				end:   domain.Point{X: 5, Y: 3},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid end",
			args: args{
				maze: &domain.Maze{
					//🧱🧱🧱🧱🧱🧱🧱
					//🧱🪙🧱⬛🌵🌵🧱
					//🧱⬛🧱🪙🧱🪙🧱
					//🧱🪨🧱🪙🧱🪙🧱
					//🧱🪨🧱🪙🧱🌵🧱
					//🧱🪙🪨⬛⬛⬛🧱
					//🧱🧱🧱🧱🧱🧱🧱
					Cells: [][]domain.CellType{
						{1, 1, 1, 1, 1, 1, 1},
						{1, 3, 0, 5, 5, 3, 1},
						{1, 1, 1, 1, 1, 5, 1},
						{1, 0, 3, 3, 3, 0, 1},
						{1, 4, 1, 1, 1, 0, 1},
						{1, 4, 3, 3, 4, 0, 1},
						{1, 1, 1, 1, 1, 1, 1},
					},
					Width:  5,
					Height: 5,
				},
				start: domain.Point{X: 3, Y: 1},
				end:   domain.Point{X: 2, Y: 1},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no solution",
			args: args{
				maze: &domain.Maze{
					//🧱🧱🧱🧱🧱🧱🧱
					//🧱🪙🧱⬛🧱🌵🧱
					//🧱⬛🧱🪙🧱🪙🧱
					//🧱🪨🧱🪙🧱🪙🧱
					//🧱🪨🧱🪙🧱🌵🧱
					//🧱🪙🪨⬛🧱⬛🧱
					//🧱🧱🧱🧱🧱🧱🧱
					Cells: [][]domain.CellType{
						{1, 1, 1, 1, 1, 1, 1},
						{1, 3, 0, 5, 5, 3, 1},
						{1, 1, 1, 1, 1, 5, 1},
						{1, 0, 3, 3, 3, 0, 1},
						{1, 1, 1, 1, 1, 1, 1},
						{1, 4, 3, 3, 4, 0, 1},
						{1, 1, 1, 1, 1, 1, 1},
					},
					Width:  5,
					Height: 5,
				},
				start: domain.Point{X: 3, Y: 1},
				end:   domain.Point{X: 5, Y: 3},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewBellmanFordSolver()
			got, err := s.Solve(tt.args.start, tt.args.end, tt.args.maze)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() got = %v, want %v", got.DrawMazeUnicode(), tt.want.DrawMazeUnicode())
			}
		})
	}
}

func TestNewBellmanFordSolver(t *testing.T) {
	t.Run("creates new solver", func(t *testing.T) {
		got := NewBellmanFordSolver()
		if got == nil {
			t.Fatal("NewBellmanFordSolver() returned nil")
		}
	})
}

func Test_buildEdges(t *testing.T) {
	tests := []struct {
		name string
		maze *domain.Maze
		want []Edge
	}{
		{
			name: "correctness",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1},
					{1, 0, 1, 1},
					{1, 3, 4, 1},
					{1, 1, 1, 1},
				},
				Width:  2,
				Height: 2,
			},
			want: []Edge{
				{From: domain.Point{X: 1, Y: 1}, To: domain.Point{X: 2, Y: 1}, Weight: 0},
				{From: domain.Point{X: 2, Y: 1}, To: domain.Point{X: 1, Y: 1}, Weight: 1},
				{From: domain.Point{X: 2, Y: 1}, To: domain.Point{X: 2, Y: 2}, Weight: 5},
				{From: domain.Point{X: 2, Y: 2}, To: domain.Point{X: 2, Y: 1}, Weight: 0},
			},
		},
		{
			name: "correctness-2",
			maze: &domain.Maze{
				Cells: [][]domain.CellType{
					{1, 1, 1, 1, 1},
					{1, 5, 3, 4, 1},
					{1, 1, 1, 1, 1},
				},
				Width:  1,
				Height: 3,
			},
			want: []Edge{
				{From: domain.Point{X: 1, Y: 1}, To: domain.Point{X: 1, Y: 2}, Weight: 0},
				{From: domain.Point{X: 1, Y: 2}, To: domain.Point{X: 1, Y: 1}, Weight: 10},
				{From: domain.Point{X: 1, Y: 2}, To: domain.Point{X: 1, Y: 3}, Weight: 5},
				{From: domain.Point{X: 1, Y: 3}, To: domain.Point{X: 1, Y: 2}, Weight: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildEdges(tt.maze)

			if !edgesEqualIgnoringOrder(got, tt.want) {
				t.Errorf("buildEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}
