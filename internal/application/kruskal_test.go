package application

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
	"reflect"
	"testing"
)

func TestKruskalGen_Generate(t *testing.T) {
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
			g := NewKruskalGen()
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
				if tt.height%2 == 0 {
					tt.height++
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

func TestKruskalGen_find(t *testing.T) {
	a := domain.Point{X: 0, Y: 0}
	b := domain.Point{X: 1, Y: 0}
	c := domain.Point{X: 2, Y: 0}

	tests := []struct {
		name        string
		disjointSet map[domain.Point]domain.Point
		point       domain.Point
		want        domain.Point
	}{
		{
			name: "correctness",
			disjointSet: map[domain.Point]domain.Point{
				a: a,
			},
			point: a,
			want:  a,
		},
		{
			name: "correctness-2",
			disjointSet: map[domain.Point]domain.Point{
				a: a,
				b: a,
				c: b,
			},
			point: c,
			want:  a,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &KruskalGen{
				disjointSet: tt.disjointSet,
			}
			if got := g.find(tt.point); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKruskalGen_union(t *testing.T) {
	a := domain.Point{X: 0, Y: 0}
	b := domain.Point{X: 1, Y: 0}
	c := domain.Point{X: 2, Y: 0}

	tests := []struct {
		name           string
		disjointSet    map[domain.Point]domain.Point
		pointA, pointB domain.Point
		want           map[domain.Point]domain.Point
	}{
		{
			name: "correctness",
			disjointSet: map[domain.Point]domain.Point{
				a: a,
				b: b,
			},
			pointA: a,
			pointB: b,
			want: map[domain.Point]domain.Point{
				a: a,
				b: a,
			},
		},
		{
			name: "correctness-2",
			disjointSet: map[domain.Point]domain.Point{
				a: a,
				b: a,
				c: c,
			},
			pointA: b,
			pointB: c,
			want: map[domain.Point]domain.Point{
				a: a,
				b: a,
				c: a,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &KruskalGen{
				disjointSet: tt.disjointSet,
			}
			for k, v := range tt.disjointSet {
				g.disjointSet[k] = v
			}

			g.union(tt.pointA, tt.pointB)

			for k, wantParent := range tt.want {
				got := g.find(k)
				if !reflect.DeepEqual(got, wantParent) {
					t.Errorf("after union, find(%v) = %v, want %v", k, got, wantParent)
				}
			}
		})
	}
}

func TestNewKruskalGen(t *testing.T) {
	t.Run("creates non-nil generator", func(t *testing.T) {
		gen := NewKruskalGen()
		if gen == nil {
			t.Fatal("NewKruskalGen() returned nil")
		}
		if gen.r == nil {
			t.Error("NewKruskalGen() did not initialize rand")
		}
	})
}
