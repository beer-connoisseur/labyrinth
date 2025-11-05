package application

import (
	"testing"
)

func TestDFSGen_Generate(t *testing.T) {
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
			g := NewDFSGen()
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

func TestNewDFSGen(t *testing.T) {
	t.Run("creates non-nil generator", func(t *testing.T) {
		gen := NewDFSGen()
		if gen == nil {
			t.Fatal("NewDFSGen() returned nil")
		}
		if gen.r == nil {
			t.Error("NewDFSGen() did not initialize rand")
		}
	})
}
