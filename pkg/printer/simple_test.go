package printer

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple_positionString(t *testing.T) {
	tests := []struct {
		pos      token.Position
		expected string
	}{
		{
			token.Position{Filename: "my/file", Offset: 0, Line: 10, Column: 4},
			"my/file:10:4",
		},
		{
			token.Position{Filename: "my/file", Offset: 0, Line: 1, Column: 0},
			"my/file:1:0",
		},
		{
			token.Position{Filename: "my/file", Offset: 0, Line: 0, Column: 4},
			"my/file",
		},
		{
			token.Position{Filename: "", Offset: 0, Line: 5, Column: 32},
			"5:32",
		},
	}

	for _, test := range tests {
		p := positionString(&test.pos)
		assert.Equal(t, test.expected, p)
	}
}

func TestSimple_Print(t *testing.T) {
	p := NewSimple()
	res := generateFileResult()
	got := captureOutput(func() {
		assert.NoError(t, p.Print(res))
	})
	expected := fmt.Sprintf("foo.txt:1:6: [warning] %s\n", res.Results[0].Reason())
	assert.Equal(t, expected, got)
}
