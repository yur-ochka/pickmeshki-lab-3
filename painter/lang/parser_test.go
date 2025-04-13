package lang_test

import (
	"strings"
	"testing"

	"github.com/yur-ochka/pickmeshki-lab-3/painter"
	"github.com/yur-ochka/pickmeshki-lab-3/painter/lang"
)

func TestParseWhiteFill(t *testing.T) {
	input := "white\nupdate"
	parser := lang.Parser{}
	ops, err := parser.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(ops) != 2 {
		t.Fatalf("Expected 2 operations, got %d", len(ops))
	}
}

func TestParseGreenFill(t *testing.T) {
	input := "green\nupdate"
	parser := lang.Parser{}
	ops, err := parser.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(ops) != 2 {
		t.Fatalf("Expected 2 operations, got %d", len(ops))
	}
}

func TestParseBgRect(t *testing.T) {
	input := "bgrect 10 20 30 40\nupdate"
	parser := lang.Parser{}
	ops, err := parser.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	found := false
	for _, op := range ops {
		if rect, ok := op.(*painter.BgRectangle); ok {
			if rect.X1 != 10 || rect.Y1 != 20 || rect.X2 != 30 || rect.Y2 != 40 {
				t.Errorf("Unexpected bgrect values: %+v", rect)
			}
			found = true
		}
	}
	if !found {
		t.Error("Expected BgRectangle operation not found")
	}
}

func TestParseFigure(t *testing.T) {
	input := "figure 5 5\nupdate"
	parser := lang.Parser{}
	ops, err := parser.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	found := false
	for _, op := range ops {
		if fig, ok := op.(*painter.Figure); ok {
			if fig.X != 5 || fig.Y != 5 {
				t.Errorf("Unexpected figure values: %+v", fig)
			}
			found = true
		}
	}
	if !found {
		t.Error("Expected Figure operation not found")
	}
}

func TestParseMove(t *testing.T) {
	input := `
figure 1 1
figure 2 2
move 10 10
update
`
	parser := lang.Parser{}
	ops, err := parser.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	foundMove := false
	for _, op := range ops {
		if move, ok := op.(*painter.Move); ok {
			if move.X != 10 || move.Y != 10 {
				t.Errorf("Unexpected move values: %+v", move)
			}
			if len(move.Figures) != 2 {
				t.Errorf("Expected 2 figures in move, got %d", len(move.Figures))
			}
			foundMove = true
		}
	}
	if !foundMove {
		t.Error("Expected Move operation not found")
	}
}

func TestParseReset(t *testing.T) {
	input := `
green
bgrect 0 0 100 100
figure 1 2
move 5 5
reset
update
`
	parser := lang.Parser{}
	ops, err := parser.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(ops) != 2 {
		t.Fatalf("Expected 2 operations (white fill and update) after reset, got %d", len(ops))
	}
}
