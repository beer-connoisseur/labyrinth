package domain

import (
	"reflect"
	"testing"
)

func TestPoint_GetNeighbors(t *testing.T) {
	tests := []struct {
		name       string
		point      Point
		directions [][2]int
		want       []Point
	}{
		{
			name:       "4 directions",
			point:      Point{X: 1, Y: 1},
			directions: [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}},
			want: []Point{
				{X: 1, Y: 0},
				{X: 1, Y: 2},
				{X: 0, Y: 1},
				{X: 2, Y: 1},
			},
		},
		{
			name:       "diagonal directions",
			point:      Point{X: 2, Y: 2},
			directions: [][2]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}},
			want: []Point{
				{X: 1, Y: 1},
				{X: 3, Y: 1},
				{X: 1, Y: 3},
				{X: 3, Y: 3},
			},
		},
		{
			name:       "single direction",
			point:      Point{X: 0, Y: 0},
			directions: [][2]int{{1, 0}},
			want: []Point{
				{X: 1, Y: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.point.GetNeighbors(tt.directions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNeighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}
