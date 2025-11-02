package infrastructure

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func Test_loadMazeFromFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "maze_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Println(err)
		}
	}(tmpFile.Name())

	maze := "###\n# #\n###\n"
	if _, err := tmpFile.WriteString(maze); err != nil {
		t.Fatal(err)
	}
	err = tmpFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "valid file",
			filename: tmpFile.Name(),
			wantErr:  false,
		},
		{
			name:     "non-existent file",
			filename: "nonexistent.txt",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadMazeFromFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadMazeFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got.Width != 1 || got.Height != 1 {
					t.Errorf("got maze size = %dx%d, want 1x1", got.Width, got.Height)
				}
				expectedCells := [][]domain.CellType{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				}
				if !reflect.DeepEqual(got.Cells, expectedCells) {
					t.Errorf("got cells = %v, want %v", got.Cells, expectedCells)
				}
			}
		})
	}
}

func Test_parsePoint(t *testing.T) {
	tests := []struct {
		name        string
		stringPoint string
		want        domain.Point
		wantErr     bool
	}{
		{
			name:        "valid point",
			stringPoint: "2,3",
			want:        domain.Point{X: 2, Y: 3},
			wantErr:     false,
		},
		{
			name:        "invalid format",
			stringPoint: " 10 , 20 ",
			want:        domain.Point{},
			wantErr:     true,
		},
		{
			name:        "invalid format-2",
			stringPoint: "5 6",
			want:        domain.Point{},
			wantErr:     true,
		},
		{
			name:        "non-integer values",
			stringPoint: "a,b",
			want:        domain.Point{},
			wantErr:     true,
		},
		{
			name:        "empty string",
			stringPoint: "",
			want:        domain.Point{},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsePoint(tt.stringPoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePoint() got = %v, want %v", got, tt.want)
			}
		})
	}
}
